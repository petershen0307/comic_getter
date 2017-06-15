package utility

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
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

func decomposePictureURL(page string) string {
	// 1. find base url
	reBaseURL := regexp.MustCompile("base href=\"([a-zA-Z0-9.:/])+\"")
	fullBaseURL := reBaseURL.FindAllString(page, 1)
	if len(fullBaseURL) == 0 {
		log.Fatalln("Can't find base url!")
		return ""
	}
	noQuote := strings.Replace(fullBaseURL[0], "\"", "", -1)
	baseURL := strings.Split(noQuote, "=")
	if len(baseURL) != 2 {
		log.Fatalln("Split url with some error")
		return ""
	}
	// 2. find picture url
	rePic := regexp.MustCompile("mangas/([a-zA-Z0-9 _/-])+.jpg")
	picture := rePic.FindAllString(page, 1)
	if len(picture) == 0 {
		log.Fatalln("Can't find the picture!")
		return ""
	}
	// 3. compose base url and picture url
	return baseURL[1] + picture[0]
}

//func decomposeChapterMaxPage()
