package main

import (
	"bytes"
	"os"

	"github.com/petershen0307/goLogger/logClient"
)

func main() {
	logClient.ModeSetting = logClient.ModePrint
	logClient.Log(logClient.LevelDebug, "hello")
	//testDownload()
	//testExtractor()
	//testExtractAllImages()
	testDownlad()
}

func testDownload() {
	buf, bRet := download("http://readcomicbooksonline.org/reader/The_Hellblazer_2016/The_Hellblazer_2016_Issue_01/?q=fullchapter")
	if bRet && buf != nil {
		fo, err := os.OpenFile("test/all_pic.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			logClient.Log(logClient.LevelDebug, "open file failed")
		}
		defer fo.Close()
		buf.WriteTo(fo)
	} else {
		logClient.Log(logClient.LevelDebug, "download file failed")
	}
}

func testExtractor() {
	fo, err := os.OpenFile("C:/Users/PC/Desktop/code/go/test/testPage.html", os.O_RDONLY, 0755)
	if err != nil {
		logClient.Log(logClient.LevelDebug, "open file failed")
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(fo)
	chapterList := extractChapter(buf.String())
	if chapterList == nil {
		logClient.Log(logClient.LevelDebug, "chapterList is nil")
	}
	logClient.Log(logClient.LevelDebug, "chapterList:%+v, len:%v", chapterList, len(chapterList))
}

func testExtractAllImages() {
	fo, err := os.OpenFile("C:/Users/PC/Desktop/code/go/test/all_pic.html", os.O_RDONLY, 0755)
	if err != nil {
		logClient.Log(logClient.LevelDebug, "open file failed")
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(fo)
	chapterList := extractChapterImages(buf.String())
	if chapterList == nil {
		logClient.Log(logClient.LevelDebug, "chapterList is nil")
	}
	logClient.Log(logClient.LevelDebug, "chapterList:%+v, len:%v", chapterList, len(chapterList))
}

func testDownlad() {
	config := downloadConfig{
		mangaChapter:    1,
		mangaMainURL:    "http://readcomicbooksonline.org/the-hellblazer-2016",
		mangaName:       "the-hellblazer-2016",
		downloadMainDir: `C:\Users\PC\Desktop\code\go\test`,
	}
	downloadMain(config)
}
