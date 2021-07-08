# PAM module for authentication using OAuth2.0

The module allows to authenticate using OAuth access token, received from auth server.

### Build module
```shell
cd pam_module
make build
```

### Build auth server
```shell
cd auth_server
go build .
```

### Get token from server
```shell
curl -g 'http://localhost:9096/token?grant_type=client_credentials&client_id=093452&client_secret=824102'
```
