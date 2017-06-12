package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
			// this is command, using --chapter or -c
			Name: "chapter, ch",
			// this is default value
			// Value: "0",
			// this is the description about this command
			Usage: "Set the chapter",
		},
		cli.StringFlag{
			Name:  "catlog, ca",
			Usage: "Set comic catlog",
		},
	}
	app.Action = func(c *cli.Context) error {
		catlog := c.String("catlog")
		log.Println("input catlog ", catlog)
		if "" == catlog {
			fmt.Println("Please input comic catlog!")
		}
		ch := c.String("chapter")
		log.Println("get input chapter ", ch)
		if "" == ch {
			fmt.Println("Please input chapter number!")
		}
		if ch, err := strconv.Atoi(ch); err == nil {
			cMap := utility.GetURLTemplate()
			core.DownloadOneChapter(ch, catlog, cMap[catlog])
		}
		return nil
	}

	app.Run(os.Args)
}
