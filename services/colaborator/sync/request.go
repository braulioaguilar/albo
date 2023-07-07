package sync

import (
	"albo/pkg/albohttp"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CharacterResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Count   int `json:"count"`
		Results []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"results"`
	} `json:"data"`
}

const (
	TS         = "1688497908"
	MARVEL_API = "http://gateway.marvel.com"
	API_KEY    = "9b2f073ee571a47da65a782a644cfffa"
	HASH       = "53c019e897a0467afba1614fdd342df9"
)

// Gets ID for specific character Iron Man and Captain America
func getCharacterByName(name string, client *albohttp.Request) ([]byte, error) {
	options := &albohttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/v1/public/characters?apikey=%s&ts=%s&hash=%s&name=%s", MARVEL_API, API_KEY, TS, HASH, name),
	}

	res, err := client.MakeRequest(options)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		str := fmt.Sprintf("Getting characters error with status code %d", res.StatusCode)
		return nil, errors.New(str)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Request character to %s\n", options.Endpoint)

	return body, nil
}

type StoryResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Count   int `json:"count"`
		Results []struct {
			ID     int `json:"id"`
			Comics struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
			} `json:"comics"`
		} `json:"results"`
	} `json:"data"`
}

// Gets stories where ironman and capamerica han been participated
func getStoriesByCharacterId(characterId int, client *albohttp.Request) ([]byte, error) {
	options := &albohttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/v1/public/stories?apikey=%s&ts=%s&hash=%s&characters=%d", MARVEL_API, API_KEY, TS, HASH, characterId),
	}

	res, err := client.MakeRequest(options)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		str := fmt.Sprintf("Getting stories error with status code %d", res.StatusCode)
		return nil, errors.New(str)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Request stories to %s\n", options.Endpoint)
	return body, nil
}

type ComicResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Count   int `json:"count"`
		Results []struct {
			Creators struct {
				Items []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
					Role        string `json:"role"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"creators"`
		} `json:"results"`
	} `json:"data"`
}

func getComicsByStori(CollectionURI string, client *albohttp.Request) ([]byte, error) {
	options := &albohttp.Options{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s?apikey=%s&ts=%s&hash=%s", CollectionURI, API_KEY, TS, HASH),
	}

	res, err := client.MakeRequest(options)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		str := fmt.Sprintf("Getting comics error with status code %d", res.StatusCode)
		return nil, errors.New(str)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Request comics to %s\n", options.Endpoint)
	return body, nil
}
