package main

import (
	"net/http"
	"net/url"
	"strings"
)

const authServerURL = "http://0.0.0.0:9096/api/authorize"

func sendAuthRequest(token string) bool {
	client := http.Client{}
	req := &http.Request{}

	form := url.Values{}
	form.Set("access_token", token)

	reqReader := strings.NewReader(form.Encode())
	req, err := http.NewRequest(http.MethodGet, authServerURL, reqReader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
