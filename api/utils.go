package api

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// from https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return errors.New("Non 200 response received")
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func saveImage(filePath, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func dumpJson(filePath string, jsonData interface{}) error {
	jsonD, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, jsonD, os.ModePerm)
}
