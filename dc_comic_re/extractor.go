package main

import (
	"strings"

	"golang.org/x/net/html"
)

func extractChapter(mangaChapterPage string) []string {
	z := html.NewTokenizer(strings.NewReader(mangaChapterPage))
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
						chapterList = append(chapterList, a.Val)
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
