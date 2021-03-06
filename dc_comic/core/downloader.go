package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/petershen0307/comic_getter/dc_comic/utility"
)

//var mangaURL = url.URL{Scheme: "http", Host: "readcomicbooksonline.net"}

const comicRootDir = "dc_comic"

// DownloadEntry to download with chapter and catlog
func DownloadEntry(ch int, comicCatlog, rootURL string) {
	utility.CreateDir(comicRootDir)
	os.Chdir(comicRootDir)
	utility.CreateDir(comicCatlog)
	os.Chdir(comicCatlog)
	rootPageDetail, err := getPage(rootURL)
	if err != nil {
		log.Fatalln(err)
		return
	}
	chapters := decomposeAllChapter(rootPageDetail)
	if ch > len(chapters) || ch < 1 {
		log.Fatalf("Got chapter length(%v) and Chapter(%v) not exist", len(chapters), ch)
		return
	}
	chDir := fmt.Sprintf("%02d", ch)
	currentDir, _ := os.Getwd()
	downloadPath := path.Join(currentDir, chDir)
	log.Printf("Download path: %v\n", downloadPath)
	log.Printf("All chapters: %v\n", chapters)
	downloadChapter(downloadPath, chapters[ch-1])
}

func downloadChapter(dir, mangaURL string) error {
	maxPageDetail, err := getPage(mangaURL)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	maxPage, err := decomposeChapterMaxPage(maxPageDetail)
	if err != err {
		log.Fatalln(err)
		return err
	}
	for i := 1; i <= maxPage; i++ {
		pictureName := fmt.Sprintf("%02d.jpg", i)
		if utility.IsFileExist(dir, pictureName) {
			continue
		}
		mangaPageURL := fmt.Sprintf("%v/%d", mangaURL, i)
		log.Printf("Download page: %v\n", mangaPageURL)
		pageDetail, err := getPage(mangaPageURL)
		pictureURL, err := decomposePictureURL(pageDetail)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		err = downloadPicture(dir, pictureName, pictureURL)
		if err != nil {
			log.Fatalln(err)
			continue
		}
	}
	return nil
}

func downloadPicture(dir, pageFileName, pictureURL string) error {
	timeoutClient := http.Client{Timeout: time.Minute * 5}
	pictureURL = strings.Replace(pictureURL, " ", "%20", -1)
	req, err := http.NewRequest("GET", pictureURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")
	response, err := timeoutClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	if http.StatusOK != response.StatusCode {
		log.Fatalf("http not ok, http status code=%d, url: %v", response.StatusCode, pictureURL)
	}
	err = utility.WriteFile(response.Body, dir, pageFileName)
	if err != nil {
		return err
	}
	return nil
}
