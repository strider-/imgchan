package imgchan

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ThreadResult struct{ Posts []Post }
type PageResult struct{ Threads []ThreadResult }

var lastApiCall time.Time

func Boards() ([]Board, error) {
	url := apiBoardUrl
	type boardResult struct{ Boards []Board }
	var result boardResult
	err := fetchAndUnmarshal(url, &result)
	return result.Boards, err
}

func Catalog(board string) ([]Page, error) {
	b := sanitizeBoard(board)
	url := fmt.Sprintf(apiCatalogUrl, b)
	var pages []Page
	err := fetchAndUnmarshal(url, &pages)

	if err == nil {
		for i := range pages {
			for j := range pages[i].Threads {
				pages[i].Threads[j].Board = b
			}
		}
	}

	return pages, err
}

func Thread(board string, postNumber int64) (ThreadResult, error) {
	url := fmt.Sprintf(apiThreadUrl, sanitizeBoard(board), postNumber)
	var result ThreadResult
	err := fetchAndUnmarshal(url, &result)

	if err == nil {
		for i := range result.Posts {
			result.Posts[i].Board = board
		}
	}

	return result, err
}

func BoardPage(board string, page int) (PageResult, error) {
	b := sanitizeBoard(board)
	url := fmt.Sprintf(apiBoardPageUrl, b, page)
	var result PageResult
	err := fetchAndUnmarshal(url, &result)

	if err == nil {
		for i := range result.Threads {
			for j := range result.Threads[i].Posts {
				result.Threads[i].Posts[j].Board = b
			}
		}
	}

	return result, err
}

func Image(post Post) ([]byte, error) {
	if !post.HasAttachment() {
		return nil, errors.New(errNoImage)
	}

	return fetchImage(post.FullImageUrl())
}

func Thumbnail(post Post) ([]byte, error) {
	if !post.HasAttachment() {
		return nil, errors.New(errNoImage)
	}

	return fetchImage(post.ThumbnailUrl())
}

func ParseThreadUrl(url string) (string, int64, error) {
	rx, err := regexp.Compile(rxThreadUrl)
	if err != nil {
		return "", 0, err
	}

	matches := rx.FindStringSubmatch(url)
	if len(matches) != 3 {
		return "", 0, errors.New(errInvalidUrl)
	}

	thread, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return "", 0, errors.New(errInvalidThreadId)
	}

	return matches[1], thread, nil
}

func fetchAndUnmarshal(url string, t interface{}) error {
	resp, err := fetch(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return errors.New(err404)
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(contents, t)
	if err != nil {
		return err
	}

	return nil
}

func fetchImage(imageUrl string) ([]byte, error) {
	resp, err := fetch(imageUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func fetch(url string) (*http.Response, error) {
	// enforcing only 1 API call per second, as stated by the 4chan API ToS
	secondDiff := time.Now().Sub(lastApiCall)
	if secondDiff.Seconds() < 1 {
		sleepTime, _ := time.ParseDuration(fmt.Sprintf("%fs", (secondDiff.Seconds()*-1)+1))
		time.Sleep(sleepTime)
	}
	lastApiCall = time.Now()
	return http.Get(url)
}

func sanitizeBoard(board string) string {
	return strings.Replace(strings.Replace(board, "/", "", -1), "\\", "", -1)
}
