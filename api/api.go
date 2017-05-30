package api

import "fmt"

type XKCD struct {
	DownloadDir string
	SaveImage   bool
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

func (x *XKCD) GetLatest() (Comic, error) {
	var err error
	url := "https://www.xkcd.com/info.0.json"
	comic := Comic{}
	err = getJson(url, &comic)
	if err == nil && x.SaveImage {
		err = x.getImage(comic)
	}
	return comic, err
}

func (x *XKCD) Get(comicNum int) (Comic, error) {
	var err error
	url := fmt.Sprintf(baseUrl, comicNum)
	comic := Comic{}
	err = getJson(url, &comic)
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

func (x *XKCD) getImage(comic Comic) error {
	filePath := fmt.Sprintf("%s/%d - %s.png",
		x.DownloadDir, comic.Num, comic.Title)
	return saveImage(comic.Img, filePath)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
