package utils

import (
	"io"
	"net/http"
)

type TflData struct {
	URL string
}

func (t TflData) FetchData() ([]byte, error) {
	req, err := http.NewRequest("GET", t.URL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
