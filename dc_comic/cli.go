package main

import (
	"fmt"
	"log"
	"os"

	"github.com/petershen0307/comic_getter/dc_comic/core"
	"github.com/petershen0307/comic_getter/dc_comic/utility"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "DC comic"
	app.Usage = "Download comic"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "catalog, ca",
			Usage: "Set comic catalog",
			Value: "wonder_woman_2016",
		},
		cli.IntFlag{
			// this is command, using --chapter or --ch
			Name: "chapter, ch",
			// this is default value
			Value: 0,
			// this is the description about this command
			Usage: "Set the chapter",
		},
		cli.StringFlag{
			// --path or -p
			Name:  "path, p",
			Value: "./bin/settings.json",
			Usage: "Setting path",
		},
	}
	app.Action = func(c *cli.Context) error {
		catalog := c.String("catalog")
		log.Println("input catalog ", catalog)
		if "" == catalog {
			fmt.Println("Please input comic catalog!")
		}
		ch := c.Int("chapter")
		log.Println("get input chapter ", ch)
		if 0 == ch {
			fmt.Println("Please input chapter number!")
		}
		path := c.String("path")
		log.Println("input setting path", path)
		// Entry point
		cMap := utility.GetURLTemplate(path)
		core.DownloadEntry(ch, catalog, cMap[catalog])
		return nil
	}

	app.Run(os.Args)
}
