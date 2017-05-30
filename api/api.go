package api

import "fmt"

type XKCD struct {
	DownloadDir string
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

var baseUrl = "https://www.xkcd.com/%d/info.0.json"

func (x *XKCD) GetLatest() Comic {
	url := "https://www.xkcd.com/info.0.json"
	comic := Comic{}
	getJson(url, &comic)
	return comic
}

func (x *XKCD) Get(comicNum int) Comic {
	url := fmt.Sprintf(baseUrl, comicNum)
	comic := Comic{}
	getJson(url, &comic)
	return comic
}

func (x *XKCD) GetAll() []Comic {
	latest := x.GetLatest().Num
	latest = 5
	return x.GetRange(1, latest)
}

func (x *XKCD) GetRange(start, end int) []Comic {
	comics := []Comic{}
	for _, comicNum := range makeRange(start, end) {
		comics = append(comics, Get(comicNum))
	}
	return comics
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
