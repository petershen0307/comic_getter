package core

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/petershen0307/comic_getter/dc_comic/utility"
)

func downloadPage(mangaURL string) string {
	timeoutRequest := http.Client{Timeout: time.Second * 30}
	response, err := timeoutRequest.Get(mangaURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	if http.StatusOK != response.StatusCode {
		log.Fatalf("http not ok, http status code=%d", response.StatusCode)
	}
	byteData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(byteData)
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
	rePic := regexp.MustCompile("mangas/([a-zA-Z0-9 _/-])+.jpg")
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
