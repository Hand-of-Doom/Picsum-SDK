package picsum

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"unicode/utf8"
)

type Image struct {
	ID      string
	Content []byte
}

type ImageProvider struct {
	url string
}

func Err404Image() error {
	return errors.New("404 not found")
}

func (p *ImageProvider) Load() (*Image, error) {
	res, err := http.Get(p.url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, Err404Image()
	}

	id := res.Header.Get("Picsum-Id")

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Image{
		ID:      id,
		Content: content,
	}, nil
}

type ImageFormat uint8

const (
	JPG ImageFormat = iota
	WebP
)

type IdentifierType uint8

const (
	ID IdentifierType = iota
	Seed
)

type identifier struct {
	value  string
	idType IdentifierType
}

type ImageOptions struct {
	width      int
	height     int
	format     ImageFormat
	identifier *identifier
	grayscale  bool
	blurDepth  int
}

func (opt *ImageOptions) Filter(grayscale bool, blurDepth int) *ImageOptions {
	opt.grayscale = grayscale
	opt.blurDepth = blurDepth

	return opt
}

func (opt *ImageOptions) Format(format ImageFormat) *ImageOptions {
	opt.format = format

	return opt
}

func (opt *ImageOptions) Identifier(value string, idType IdentifierType) *ImageOptions {
	if utf8.RuneCountInString(value) == 0 {
		value = "1"
		idType = ID
	}

	opt.identifier = &identifier{value, idType}

	return opt
}

func NewImageOptions(width, height int) *ImageOptions {
	options := new(ImageOptions)
	options.width = width
	options.height = height
	options.format = JPG

	return options
}

func NewImageProvider(options *ImageOptions) *ImageProvider {
	imageURL := "https://picsum.photos/"

	identifier := options.identifier
	if identifier != nil {
		switch identifier.idType {
		case ID:
			imageURL += fmt.Sprintf("id/%s/", identifier.value)
		case Seed:
			imageURL += fmt.Sprintf("seed/%s/", identifier.value)
		}
	}

	imageURL += fmt.Sprintf("%d/%d",
		options.width, options.height)

	switch options.format {
	case JPG:
		imageURL += ".jpg"
	case WebP:
		imageURL += ".webp"
	}

	query := url.Values{}
	if options.grayscale {
		query.Set("grayscale", "")
	}
	if options.blurDepth != 0 {
		query.Set("blur", fmt.Sprintf("%d", options.blurDepth))
	}
	queryString := query.Encode()
	if queryString != "" {
		imageURL += "?" + queryString
	}

	return &ImageProvider{imageURL}
}
