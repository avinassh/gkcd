package cli

import (
	"log"
	"os"
)

func validateDownloadDir(downloadDir string) {
	err := os.MkdirAll(downloadDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func validateRange(comicNum, start, end uint) {
	if start > 0 && comicNum > 0 {
		log.Fatal("You cannot specify both the --comic and --start flags")
	}
	if start > 0 && start > end {
		log.Fatal("Invalid range specified")
	}
	if start > 0 && end == 0 {
		log.Fatal("End flag value is missing")
	}
}
