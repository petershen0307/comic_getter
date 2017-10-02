package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/petershen0307/comic_getter/dc_comic/utility"
)

func getPage(mangaURL string) (string, error) {
	timeoutRequest := http.Client{Timeout: time.Minute * 5}
	response, err := timeoutRequest.Get(mangaURL)
	if err != nil {
		log.Panicln("url:", mangaURL, "err:", err)
	}
	defer response.Body.Close()
	if http.StatusOK != response.StatusCode {
		return "", utility.MyError{What: fmt.Sprintf("Got http status code: %v", response.StatusCode)}
	}
	byteData, _ := ioutil.ReadAll(response.Body)
	return string(byteData), nil
}

func decomposeAllChapter(rootPageDetail string) []string {
	regexChapter := regexp.MustCompile(`http://readcomicbooksonline.net/reader/[a-zA-Z0-9 /_\-\(\)]+`)
	allChapter := regexChapter.FindAllString(rootPageDetail, -1)
	// reverse slice
	for i, j := 0, len(allChapter)-1; i < j; i, j = i+1, j-1 {
		allChapter[i], allChapter[j] = allChapter[j], allChapter[i]
	}
	return allChapter
}

func decomposePictureURL(page string) (string, error) {
	// 1. find base url
	reBaseURL := regexp.MustCompile("base href=\"([a-zA-Z0-9.:/])+\"")
	fullBaseURL := reBaseURL.FindAllString(page, 1)
	if len(fullBaseURL) == 0 {
		return "", utility.MyError{What: "Can't find base url!"}
	}
	noQuote := strings.Replace(fullBaseURL[0], "\"", "", -1)
	baseURL := strings.Split(noQuote, "=")
	if len(baseURL) != 2 {
		return "", utility.MyError{What: "Split url with some error"}
	}
	// 2. find picture url
	rePic := regexp.MustCompile(`mangas/([a-zA-Z0-9 _/\-\(\)])+.jpg`)
	picture := rePic.FindAllString(page, 1)
	if len(picture) == 0 {
		return "", utility.MyError{What: "Can't find the picture!"}
	}
	// 3. compose base url and picture url
	return baseURL[1] + picture[0], nil
}

func decomposeChapterMaxPage(page string) (int, error) {
	regexMaxPage := regexp.MustCompile("of [0-9]+")
	maxPage := regexMaxPage.FindAllString(page, 1)
	if len(maxPage) == 0 {
		return 0, utility.MyError{What: "Can't find the max page number!"}
	}
	pageNumber := strings.Split(maxPage[0], " ")
	num, err := strconv.Atoi(pageNumber[1])
	if err != nil {
		return 0, err
	}
	return num, nil
}
