package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/petershen0307/comic_getter/dc_comic/justiceLeague"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "New52 Justice League"
	app.Usage = "Download comic"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			// this is command, using --chapter or -c
			Name: "chapter,c",
			// this is default value
			// Value: "0",
			// this is the description about this command
			Usage: "Get the chapter",
		},
	}
	app.Action = func(c *cli.Context) error {
		ch := c.String("chapter")
		log.Println("get input chapter ", ch)
		if "" == ch {
			fmt.Println("Please input chapter number!")
		}
		if ch, err := strconv.Atoi(ch); err == nil {
			downloadOneChapter(ch, justiceLeague.ComicCatlog, justiceLeague.URLTemplate)
		}
		return nil
	}

	app.Run(os.Args)
}
