package main

import (
	"fmt"
	"log/syslog"
	"net/http"
	"runtime"
)

type AuthResult int

const (
	AuthError AuthResult = iota
	AuthSuccess
)

const (
	targetUsername = "oauthuser"
	authServerURL  = "http://localhost:9096/api/authorize"
)

func sendAuthRequest(token string) bool {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, authServerURL, nil)

	query := req.URL.Query()
	query.Add("access_token", token)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		pamLog("authentication request error: %s\n", err.Error())
		return false
	}

	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func pamLog(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-oauth2")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}

func pamAuthenticate(username string, token string) AuthResult {
	runtime.GOMAXPROCS(1)

	if username != targetUsername {
		pamLog("user %s is not interesting\n", username)
		return AuthSuccess
	}

	if sendAuthRequest(token) {
		pamLog("successfully authorized with token: %s\n", token)
		return AuthSuccess
	}
	pamLog("failed to authorize with token: %s\n", token)

	return AuthError
}

func main() {}
