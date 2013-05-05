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
	url := "https://api.4chan.org/boards.json"
	type boardResult struct{ Boards []Board }
	var result boardResult
	err := fetchAndUnmarshal(url, &result)
	return result.Boards, err
}

func Catalog(board string) ([]Page, error) {
	url := fmt.Sprintf("https://api.4chan.org/%s/catalog.json", sanitizeBoard(board))
	var pages []Page
	err := fetchAndUnmarshal(url, &pages)
	return pages, err
}

func Thread(board string, postNumber int64) (ThreadResult, error) {
	url := fmt.Sprintf("https://api.4chan.org/%s/res/%d.json", sanitizeBoard(board), postNumber)
	var result ThreadResult
	err := fetchAndUnmarshal(url, &result)
	return result, err
}

func BoardPage(board string, page int) (PageResult, error) {
	url := fmt.Sprintf("https://api.4chan.org/%s/%d.json", sanitizeBoard(board), page)
	var result PageResult
	err := fetchAndUnmarshal(url, &result)
	return result, err
}

func Image(board string, post Post) ([]byte, error) {
	url := fmt.Sprintf("https://images.4chan.org/%s/src/%d%s", sanitizeBoard(board), post.RenamedFilename, post.Ext)
	resp, err := fetch(url)
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

func ParseThreadUrl(url string) (string, int64, error) {
	rx, err := regexp.Compile(`/(?P<board>[^/])/res/(?P<thread>\d+)$`)
	if err != nil {
		return "", 0, err
	}

	matches := rx.FindStringSubmatch(url)
	if len(matches) != 3 {
		return "", 0, errors.New("invalid 4chan url")
	}

	thread, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return "", 0, errors.New("invalid 4chan thread id")
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
		return errors.New("specified resource is no longer available!")
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
