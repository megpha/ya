package mytube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	API_KEY  = "AIzaSyDTU4jimlpM7ddJvf7vBTY6v_N_vJNgI0w"
	ENDPOINT = "https://www.googleapis.com/youtube/v3/"
	HANDLE   = "iDreamTeluguNews"
)

func URL(resource string, params map[string]string) string {
	queryParams := url.Values{}
	queryParams.Set("key", API_KEY)

	for k, v := range params {
		queryParams.Set(k, v)
	}

	queryUrl := fmt.Sprint(ENDPOINT, resource, "?", queryParams.Encode())
	return queryUrl
}

func Videos(channel string) []Video {
	channelID := channelID(channel)

	params := map[string]string{
		"part":       "snippet",
		"maxResults": "50",
		"channelId":  channelID,
	}
	queryUrl := URL("search", params)
	type videos struct {
		Items []struct {
			ID struct {
				Value string `json:"videoId"`
			}
			Snippet struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"snippet"`
		} `json:"items"`
	}
	resp, err := http.Get(queryUrl)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	var answer videos
	jsonErr := json.Unmarshal(data, &answer)

	if jsonErr != nil {
		log.Fatal(err)
	}

	results := make([]Video, len(answer.Items))

	for i, item := range answer.Items {
		results[i] = Video{item.Snippet.Title, item.Snippet.Description, item.ID.Value, channel}
	}

	return results
}

func channelID(channelHandle string) string {
	params := map[string]string{
		"forUsername": channelHandle,
		"part":        "id",
		"order":       "date",
	}

	queryUrl := URL("channels", params)
	resp, err := http.Get(queryUrl)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)

	type channel struct {
		Etag  string `json:"etag"`
		Items []struct {
			Id string
		} `json:"items"`
	}
	var answer interface{}
	json.Unmarshal(data, &answer)

	var empty channel
	jsonErr := json.Unmarshal(data, &empty)

	if jsonErr != nil {
		log.Fatal(err)
	}
	return empty.Items[0].Id
}
