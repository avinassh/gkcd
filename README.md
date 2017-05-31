# gkcd

An xkcd downloader written in Go.

## Installation

    $ go get github.com/avinassh/gkcd
    $ go install github/avinassh/gkcd

## Usage

```
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
```

## License

The mighty MIT license. Please check `LICENSE` for more details.