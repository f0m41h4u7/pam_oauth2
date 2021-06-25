package main

import (
  "bufio"
  "bytes"
  "crypto/rand"
  "fmt"
  "io"
  "io/ioutil"
  "log/syslog"
  "net"
  "os"
  "path"
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

func authenticate(w io.Writer, uid int, username, ca string, principals map[string]struct{}) AuthResult {
    for _, p := range cert.ValidPrincipals {
      if _, ok := principals[p]; ok {
        pamLog("Authentication succeded for %s. Matched principal %s, cert %d",
          cert.ValidPrincipals[0], p, cert.Serial)
        return AuthSuccess
      }
    }
  }
  pamLog("no valid certs found")
  return AuthError
}

func loadValidPrincipals(principals string) (map[string]struct{}, error) {
  f, err := os.Open(principals)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  p := make(map[string]struct{})
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    p[scanner.Text()] = struct{}{}
  }
  return p, nil
}

func pamAuthenticate(w io.Writer, username string, argv []string) AuthResult {
  runtime.GOMAXPROCS(1)

  if username != targetUsername {
    pamLog("user %s is not interesting\n", username)
    return AuthSuccess
  }

  if len(argv) == 0 {
    pamLog("empty args list")
    return AuthError
  }

  opt := strings.Split(arg, "=")
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
