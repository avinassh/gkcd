package api

import (
	"fmt"
)

type XKCD struct {
	DownloadDir string
	SaveImage   bool
	SaveMeta    bool
}

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

func (x *XKCD) GetLatest() (Comic, error) {
	return x.Get(0)
}

func (x *XKCD) Get(comicNum int) (Comic, error) {
	var err error
	url := getMetaURL(comicNum)
	comic := Comic{}
	err = getJson(url, &comic)
	if err == nil && x.SaveMeta {
		err = x.getMeta(comic)
	}
	if err == nil && x.SaveImage {
		err = x.getImage(comic)
	}
	return comic, err
}

func (x *XKCD) GetAll() ([]Comic, error) {
	comic, err := x.GetLatest()
	if err != nil {
		return nil, err
	}
	return x.GetRange(1, comic.Num)
}

func (x *XKCD) GetRange(start, end int) ([]Comic, error) {
	comics := []Comic{}
	for _, comicNum := range makeRange(start, end) {
		comic, err := x.Get(comicNum)
		if err != nil {
			return comics, err
		}
		comics = append(comics, comic)
	}
	return comics, nil
}

func (x *XKCD) getMeta(comic Comic) error {
	return dumpJson(x.getFilePath(comic, "json"), comic)
}

func (x *XKCD) getImage(comic Comic) error {
	return saveImage(x.getFilePath(comic, "png"), comic.Img)
}

func (x *XKCD) getFilePath(comic Comic, extension string) string {
	return fmt.Sprintf("%s/%d - %s.%s",
		x.DownloadDir, comic.Num, comic.Title, extension)
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
