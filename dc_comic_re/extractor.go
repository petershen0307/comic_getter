package main

import (
	"strings"
	"unicode/utf8"

	"github.com/petershen0307/goLogger/logClient"

	"golang.org/x/net/html"
)

func extractChapter(mangaMainPage string) []string {
	z := html.NewTokenizer(strings.NewReader(mangaMainPage))
	var chapterList []string
	for foundChapterLink := false; ; {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return chapterList
		case tt == html.EndTagToken:
			t := z.Token()
			if t.Data == "div" {
				foundChapterLink = false
			}
		case tt == html.StartTagToken:
			t := z.Token()
			// 連結存在 <div id="chapterlist"> </div> 之間
			// <div id="chapterlist">
			if t.Data == "div" {
				for _, div := range t.Attr {
					if div.Key == "id" && div.Val == "chapterlist" {
						foundChapterLink = true
						continue
					}
				}
			}
			// <a href="">
			if foundChapterLink && t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						// insert url to head
						temp := append([]string{}, a.Val)
						chapterList = append(temp, chapterList...)
					}
				}
			}
			// <li class="chapter">
			// if t.Data == "li" {
			// 	for _, li := range t.Attr {
			// 		if li.Key == "class" && li.Val == "chapter" {
			// 			foundChapterLink = true
			// 		}
			// 	}
			// }
		}
	}
}

func extractChapterImages(fullChapterPage string) []string {
	z := html.NewTokenizer(strings.NewReader(fullChapterPage))
	var chapterList []string
	var baseURL string // should end with '/'
	for foundImageLink := false; ; {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return chapterList
		case tt == html.SelfClosingTagToken:
			t := z.Token()

			// <base href="">
			if t.Data == "base" {
				for _, base := range t.Attr {
					if base.Key == "href" {
						baseURL = base.Val
						lastRune, _ := utf8.DecodeLastRuneInString(baseURL)
						if len(baseURL) == 0 {
							logClient.Log(logClient.LevelWarning, "base url is empty")
						} else if lastRune != '/' {
							baseURL += "/"
						}
						continue
					}
				}
			}
			// <img src=""/>
			if foundImageLink && t.Data == "img" {
				for _, img := range t.Attr {
					if img.Key == "src" {
						// insert url to head
						imgFullURL := baseURL + img.Val
						temp := append([]string{}, imgFullURL)
						chapterList = append(temp, chapterList...)
					}
				}
			}
		case tt == html.StartTagToken:
			t := z.Token()
			// 連結存在 <div class="fullchapter"> </div> 這個tag之後
			if t.Data == "div" {
				for _, div := range t.Attr {
					// <div class="fullchapter">
					if div.Key == "class" && div.Val == "fullchapter" {
						foundImageLink = true
						continue
					}
					// <div id="block300">
					if div.Key == "id" && div.Val == "block300" {
						foundImageLink = false
						return chapterList
					}
				}
			}
		}
	}
}
