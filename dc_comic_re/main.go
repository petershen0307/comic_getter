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
	testExtractor()
}

func testDownload() {
	buf, bRet := download("http://readcomicbooksonline.org/the-hellblazer-2016")
	if bRet && buf != nil {
		fo, err := os.OpenFile("test/testPage.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
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
