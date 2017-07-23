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
			Name:  "catalog, ca",
			Usage: "Set comic catalog",
		},
	}
	app.Action = func(c *cli.Context) error {
		catalog := c.String("catalog")
		log.Println("input catalog ", catalog)
		if "" == catalog {
			fmt.Println("Please input comic catalog!")
		}
		ch := c.String("chapter")
		log.Println("get input chapter ", ch)
		if "" == ch {
			fmt.Println("Please input chapter number!")
		}
		if chInt, err := strconv.Atoi(ch); err == nil {
			cMap := utility.GetURLTemplate()
			core.DownloadEntry(chInt, catalog, cMap[catalog])
		}
		return nil
	}

	app.Run(os.Args)
}
