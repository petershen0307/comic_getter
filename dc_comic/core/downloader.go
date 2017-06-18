package core

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/petershen0307/comic_getter/dc_comic/utility"
)

var webCookieJar *cookiejar.Jar
var mangaURL = url.URL{Scheme: "http", Host: "readcomicbooksonline.net"}

const comicRootDir = "dc_comic"

type errorCode int

const (
	eSuccess = errorCode(iota)
	ePageNotFound
	eFileExisted
)

// DownloadOneChapter to download with chapter and catlog
func DownloadOneChapter(ch int, comicCatlog, urlTemplate string) {
	utility.CreateDir(comicRootDir)
	os.Chdir(comicRootDir)
	utility.CreateDir(comicCatlog)
	os.Chdir(comicCatlog)
	for i := 1; downloadOnePage(ch, i, urlTemplate) != ePageNotFound; i++ {
	}
}

func downloadOnePage(ch, page int, urlTemplate string) errorCode {
	dir := fmt.Sprintf("%02d", ch)
	fileName := fmt.Sprintf("%02d.jpg", page)
	if utility.IsFileExist(dir, fileName) {
		return eFileExisted
	}
	timeoutRequest := http.Client{Timeout: time.Second * 30}
	if nil != webCookieJar {
		timeoutRequest.Jar = webCookieJar
	}
	mangaURL.Path = fmt.Sprintf(urlTemplate, ch, page)
	targetURL := mangaURL.String()
	response, err := timeoutRequest.Get(targetURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if 404 == response.StatusCode {
		log.Println("Page not found! url: ", targetURL)
		return ePageNotFound
	}
	if -1 == response.ContentLength {
		log.Println("url: ", targetURL, ";status code: ", response.StatusCode)
		return ePageNotFound
	}
	if nil == webCookieJar {
		webCookieJar, _ = cookiejar.New(nil)
		cookieStr := strings.Split(response.Header["Set-Cookie"][0], ";")
		cookieArray := strings.Split(cookieStr[0], "=")
		cookie := &http.Cookie{Name: cookieArray[0], Value: cookieArray[1]}
		webCookieJar.SetCookies(&mangaURL, []*http.Cookie{cookie})
		log.Println("set cookie", cookie)
	}
	fmt.Println("1 Content Length: ", response.ContentLength)
	utility.WriteFile(response.Body, dir, fileName)
	return eSuccess
}
