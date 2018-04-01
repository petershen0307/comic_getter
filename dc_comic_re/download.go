package main

import (
	"bytes"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/petershen0307/goLogger/logClient"
)

type downloadConfig struct {
	mangaName       string
	mangaMainURL    string
	mangaChapter    uint
	downloadMainDir string
}

func download(mangaURL string) (*bytes.Buffer, bool) {
	timeoutRequest := http.Client{Timeout: time.Minute * 5}
	response, err := timeoutRequest.Get(mangaURL)
	if err != nil {
		logClient.Log(logClient.LevelErr, "url:%v, err:%v", mangaURL, err)
		return nil, false
	}
	defer response.Body.Close()
	if http.StatusOK != response.StatusCode {
		logClient.Log(logClient.LevelErr, "url:%v, status:%v", mangaURL, response.StatusCode)
		return nil, false
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf, true
}

func downloadMain(config downloadConfig) {
	// check download main dir exist or not
	CreateDir(config.downloadMainDir)
	// generate download dir to manga name
	downloadMangaDir := filepath.Join(config.downloadMainDir, config.mangaName)
	CreateDir(downloadMangaDir)
	// generate download dir to chapter
	downloadChapterDir := filepath.Join(downloadMangaDir, strconv.FormatUint(uint64(config.mangaChapter), 10))
	CreateDir(downloadChapterDir)

	// get manga chapter list
	var mangaChapterList []string
	if mainPage, bRet := download(config.mangaMainURL); bRet {
		mangaChapterList = extractChapter(mainPage.String())
	} else {
		logClient.Log(logClient.LevelWarning, "Can't get manga main page(%v).", config.mangaMainURL)
	}

	// get manga images by chapter
	var mangaChapterImageList []string
	if int(config.mangaChapter) <= len(mangaChapterList) {
		fullChapterURL := mangaChapterList[config.mangaChapter-1] + "/?q=fullchapter"
		if chapterPage, bRet := download(fullChapterURL); bRet {
			mangaChapterImageList = extractChapterImages(chapterPage.String())
		} else {
			logClient.Log(logClient.LevelWarning, "Can't get manga chapter page(%v).", fullChapterURL)
		}
	}

	// download all images
	if len(mangaChapterImageList) == 0 {
		logClient.Log(logClient.LevelWarning, "Can't extract manga chapter image url.(ch:%v),(mainURL:%v)", config.mangaChapter, config.mangaMainURL)
	}
	for _, imageURL := range mangaChapterImageList {
		imageFileName := path.Base(imageURL)
		if !IsFileExist(downloadChapterDir, imageFileName) {
			if image, bRet := download(imageURL); bRet {
				if err := WriteFile(image, downloadChapterDir, imageFileName); err != nil {
					logClient.Log(logClient.LevelWarning, "Write manga image to disk failed(%v).", imageURL)
				}
			} else {
				logClient.Log(logClient.LevelWarning, "Can't download manga image(%v).", imageURL)
			}
		}
	}
}
