package utility

// GetURLTemplate to get url map, key:comicCatlog, value:url template
func GetURLTemplate() map[string]string {
	urlTemplate := map[string]string{
		"justice_league": "reader/mangas/Justice League/Justice League %03[1]d/jlu-ch%03[1]d-%02[2]d.jpg",
	}
	return urlTemplate
}

// const urlTemplate = "http://www.readcomics.tv/images/manga/justice-league/%02d/%d.jpg"
// const urlTemplate = "http://readcomicbooksonline.net/reader/mangas/Justice League/Justice League %03[1]d/jlu-ch%03[1]d-%02[2]d.jpg"
