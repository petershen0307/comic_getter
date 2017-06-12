package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
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
	createDir(comicRootDir)
	os.Chdir(comicRootDir)
	createDir(comicCatlog)
	os.Chdir(comicCatlog)
	for i := 1; downloadOnePage(ch, i, urlTemplate) != ePageNotFound; i++ {
	}
}

func createDir(pathName string) {
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		fmt.Println("Create root dir: ", pathName)
		if err := os.Mkdir(pathName, os.ModeDir); err != nil {
			panic(err)
		}
	}
}

func downloadOnePage(ch, page int, urlTemplate string) errorCode {
	dir := fmt.Sprintf("%02d", ch)
	fileName := fmt.Sprintf("%02d.jpg", page)
	if isFileExist(dir, fileName) {
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
	writeFile(response.Body, dir, fileName)
	return eSuccess
}

func isFileExist(dir, fileName string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Create dir: ", dir)
		if err := os.Mkdir(dir, os.ModeDir); err != nil {
			panic(err)
		}
	}
	// check path exist or not
	// if path exist os.Stat() return err will be nil, can't use os.IsExist()
	/*
		os.IsExist(), os.IsNotExist() reference: https://gist.github.com/mastef/05f46d3ab2f5ed6a6787
		os.IsExist(err) is good for cases when you expect the file to not exist yet,
		but the file actually exists :
	*/
	path := fmt.Sprintf("%s/%s", dir, fileName)
	if _, err := os.Stat(path); err == nil {
		fmt.Println("File exist: ", path)
		return true
	}
	return false
}

func writeFile(reader io.Reader, dir, fileName string) {
	if isFileExist(dir, fileName) {
		return
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("2 Data Length: ", len(data))
	path := fmt.Sprintf("%s/%s", dir, fileName)
	fo, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	fo.Write(data)
}
