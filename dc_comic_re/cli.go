package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/petershen0307/goLogger/logClient"

	"github.com/urfave/cli"
)

// GetURLTemplate to get url map, key:comicCatlog, value:url template
func GetURLTemplate(settingPath string) map[string]string {
	urlTemplate := map[string]string{
		"justice_league":      "http://readcomicbooksonline.net/justice-league",
		"wonder_woman_2016":   "http://readcomicbooksonline.net/wonder-woman-2016",
		"the-hellblazer-2016": "http://readcomicbooksonline.org/the-hellblazer-2016",
		"batman-hush":         "http://readcomicbooksonline.net/batman-the-complete-hush",
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
		logClient.Log(logClient.LevelErr, "open file failed, err:%v", err)
	}
	var settings Settings
	err = json.Unmarshal(rawData, &settings)
	if err != nil {
		logClient.Log(logClient.LevelErr, "unmarshal json failed, err:%v", err)
	}
	return settings.Catalogs
}

func cliMain() {
	app := cli.NewApp()
	app.Name = "DC comic"
	app.Usage = "Download comic"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dir, d",
			Usage: "Set output dir",
			Value: `c:\`,
		},
		cli.StringFlag{
			Name:  "catalog, n",
			Usage: "Set comic catalog",
			Value: "the-hellblazer-2016",
		},
		cli.UintFlag{
			// this is command, using --chapter or -c
			Name: "chapter, c",
			// this is default value
			Value: 1,
			// this is the description about this command
			Usage: "Set the chapter",
		},
		cli.StringFlag{
			// --path or -p
			Name:  "path, p",
			Value: "settings.json",
			Usage: "Setting path",
		},
		cli.StringFlag{
			// --log or -l
			Name:  "log, l",
			Value: "console",
			Usage: "Set log mode(pipe, console)",
		},
	}
	app.Action = func(c *cli.Context) error {
		log := c.String("log")
		if log == "console" {
			logClient.ModeSetting = logClient.ModePrint
			logClient.Log(logClient.LevelInfo, "log mode is console")
		} else {
			logClient.ModeSetting = logClient.ModePipe
			logClient.Log(logClient.LevelInfo, "log mode is pipe")
		}
		catalog := c.String("catalog")
		logClient.Log(logClient.LevelInfo, "input catalog (%v)", catalog)
		chapter := c.Uint("chapter")
		logClient.Log(logClient.LevelInfo, "input chapter (%v)", chapter)
		path := c.String("path")
		logClient.Log(logClient.LevelInfo, "input setting path (%v)", path)
		dir := c.String("dir")
		logClient.Log(logClient.LevelInfo, "input download dir (%v)", dir)
		// Entry point
		cMap := GetURLTemplate(path)
		config := downloadConfig{
			mangaName:       catalog,
			mangaChapter:    chapter,
			mangaMainURL:    cMap[catalog],
			downloadMainDir: dir,
		}
		downloadMain(config)
		return nil
	}

	app.Run(os.Args)
}
