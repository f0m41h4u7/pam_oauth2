package main

import (
  "fmt"
  "log/syslog"
  "runtime"
  "strings"
)

type AuthResult int
const (
  AuthError AuthResult = iota
  AuthSuccess
)

const targetUsername = "macdit"

func pamLog(format string, args ...interface{}) {
  l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-oauth2")
  if err != nil {
    return
  }
  l.Warning(fmt.Sprintf(format, args...))
}

func pamAuthenticate(username string, argv []string) AuthResult {
  runtime.GOMAXPROCS(1)

  if username != targetUsername {
    pamLog("user %s is not interesting\n", username)
    return AuthSuccess
  }

  if len(argv) == 0 {
    pamLog("empty args list")
    return AuthError
  }

  opt := strings.Split(argv[0], "=")
  switch opt[0] {
  case "access_token":
    token := opt[1]
    if sendAuthRequest(token) {
      pamLog("successfully authorized with token: %s\n", token)
      return AuthSuccess
    }
    pamLog("failed to authorize with token: %s\n", token)
  default:
    pamLog("unkown option: %s\n", opt[0])
  }

  return AuthError
}

func main() {}
