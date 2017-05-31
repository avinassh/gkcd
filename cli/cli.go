package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	gkcd "github.com/avinassh/gkcd/api"
)

var helpText = `
gkcd is a simple command line tool to download XCKD comics, written in Go.
For complete documentation visit http://avi.im/gkcd

gkcd run without any flags/options, downloads latest comic.

Usage:

  -all
    	Downloads all the comics till date
  -comic
    	Downloads specified comic. eg. $ gkcd --comic 42
  -start
    	Start range for the comics to download
  -end
    	End range for the comics to download
  -meta
    	Do you also want to download metadata with the comic?
  -path
    	Path where you want to store the downloaded comics. If no path
    	specified, then it downloads in current directory.
`

func Start() {
	var downloadDir string
	var comicNum uint
	var start uint
	var end uint
	var all bool
	var meta bool
	flag.StringVar(&downloadDir, "path", getCurrentDir(), "")
	flag.UintVar(&comicNum, "comic", 0, "")
	flag.UintVar(&start, "start", 0, "")
	flag.UintVar(&end, "end", 0, "")
	flag.BoolVar(&all, "all", false, "")
	flag.BoolVar(&meta, "meta", false, "")

	flag.Usage = func() {
		fmt.Printf(helpText)
	}
	flag.Parse()

	validateDownloadDir(downloadDir)
	validateRange(comicNum, start, end)

	var comic gkcd.Comic
	var comics []gkcd.Comic
	var err error

	if all == true {
		comics, err = gkcd.GetAll()
	} else if start > 0 && end > 0 {
		comics, err = gkcd.GetRange(int(start), int(end))
	} else if comicNum > 0 {
		comic, err = gkcd.Get(int(comicNum))
	} else {
		comic, err = gkcd.GetLatest()
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(comics) == 0 {
		comics = append(comics, comic)
	}
	for _, c := range comics {
		if err = save(c, downloadDir, meta); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func getCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd
}

func save(c gkcd.Comic, downloadDir string, meta bool) error {
	if meta {
		return gkcd.SaveComicWithMeta(c, downloadDir)
	} else {
		return gkcd.SaveComic(c, downloadDir)
	}
}
