MODULE_NAME=pam_oauth2

build:
	go build -buildmode=c-shared -o ${MODULE}.so

clean:
	rm -f ${MODULE}.so ${MODULE}.h

.PHONY: build clean
