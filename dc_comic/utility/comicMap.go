package utility

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// GetURLTemplate to get url map, key:comicCatlog, value:url template
func GetURLTemplate(settingPath string) map[string]string {
	urlTemplate := map[string]string{
		"justice_league":    "http://readcomicbooksonline.net/justice-league",
		"wonder_woman_2016": "http://readcomicbooksonline.net/wonder-woman-2016",
	}
	if "" != settingPath {
		urlTemplate = make(map[string]string)
		mangaList := readSettings(settingPath)
		for _, catalog := range mangaList {
			urlTemplate[catalog.Name] = catalog.Path
		}
	}
	return urlTemplate
}

// Catalog collect manga name and url
type Catalog struct {
	Name string
	Path string
}

// Settings is the manga configure file structure
type Settings struct {
	Catalogs []Catalog
}

func readSettings(settingPath string) []Catalog {
	rawData, err := ioutil.ReadFile(settingPath)
	if err != nil {
		log.Fatalln("open file failed, err:", err)
	}
	var settings Settings
	err = json.Unmarshal(rawData, &settings)
	if err != nil {
		log.Fatalln("unmarshal json failed, err:", err)
	}
	return settings.Catalogs
}

// const urlTemplate = "http://www.readcomics.tv/images/manga/justice-league/%02d/%d.jpg"
// const urlTemplate = "http://readcomicbooksonline.net/reader/mangas/Justice League/Justice League %03[1]d/jlu-ch%03[1]d-%02[2]d.jpg"
