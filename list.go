package picsum

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ImageDetails struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	URL         string `json:"url"`
	DownloadURL string `json:"download_url"`
}

type ImagesList struct {
	Value    []*ImageDetails
	LastPage bool
}

func GetImagesList(page, limit int) (*ImagesList, error) {
	url := fmt.Sprintf("https://picsum.photos/v2/list?page=%d&limit=%d",
		page, limit)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var images []*ImageDetails
	err = json.NewDecoder(res.Body).Decode(&images)
	if err != nil {
		return nil, err
	}

	imagesList := new(ImagesList)
	imagesList.Value = images

	linkHeader := res.Header.Get("Link")
	if !strings.Contains(linkHeader, "rel=\"next\"") {
		imagesList.LastPage = true
	}

	return imagesList, nil
}

func GetImageDetails(identifier string, idType IdentifierType) (*ImageDetails, error) {
	var url string
	switch idType {
	case ID:
		url = fmt.Sprintf("https://picsum.photos/id/%s/info", identifier)
	case Seed:
		url = fmt.Sprintf("https://picsum.photos/seed/%s/info", identifier)
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, Err404Image()
	}

	var imageDetails *ImageDetails
	err = json.NewDecoder(res.Body).Decode(&imageDetails)

	return imageDetails, err
}
