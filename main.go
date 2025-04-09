package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/mikerybka/util"
)

func searchYouTube(query, apiKey string) error {
	params := url.Values{}
	params.Set("part", "snippet")
	params.Set("q", query)
	params.Set("maxResults", "2")
	params.Set("type", "video")
	params.Set("key", apiKey)

	url := "https://www.googleapis.com/youtube/v3/search?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("YouTube search failed: %s", resp.Status)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<query>")
		os.Exit(2)
	}
	err := searchYouTube(os.Args[1], util.RequireEnvVar("YOUTUBE_API_KEY"))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
