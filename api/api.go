package api

import (
	"fmt"
)

var (
	IMAGE_EXT = "png"
	META_EXT  = "json"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func GetLatest() (Comic, error) {
	return Get(0)
}

func Get(comicNum int) (Comic, error) {
	var err error
	url := getMetaURL(comicNum)
	comic := Comic{}
	err = getJson(url, &comic)
	return comic, err
}

func GetAll() ([]Comic, error) {
	comic, err := GetLatest()
	if err != nil {
		return nil, err
	}
	return GetRange(1, comic.Num)
}

func GetRange(start, end int) ([]Comic, error) {
	comics := []Comic{}
	for _, comicNum := range makeRange(start, end) {
		comic, err := Get(comicNum)
		if err != nil {
			return comics, err
		}
		comics = append(comics, comic)
	}
	return comics, nil
}

func getFilePath(comic Comic, downloadDir, extension string) string {
	return fmt.Sprintf("%s/%d - %s.%s",
		downloadDir, comic.Num, comic.Title, extension)
}

func SaveComic(comic Comic, downloadDir string) error {
	filePath := getFilePath(comic, downloadDir, IMAGE_EXT)
	return saveImage(filePath, comic.Img)
}

func SaveComicWithMeta(comic Comic, downloadDir string) error {
	filePath := getFilePath(comic, downloadDir, META_EXT)
	if err := dumpJson(filePath, comic); err != nil {
		return err
	}
	return SaveComic(comic, downloadDir)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func getMetaURL(comicNum int) string {
	if comicNum == 0 {
		return "https://www.xkcd.com/info.0.json"
	}
	baseUrl := "https://www.xkcd.com/%d/info.0.json"
	return fmt.Sprintf(baseUrl, comicNum)
}
