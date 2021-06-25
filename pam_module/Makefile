MODULE_NAME := pam_oauth2

build:
	go build -buildmode=c-shared -o ${MODULE_NAME}.so

install:
	cp ${MODULE_NAME}.so /lib64/security/

clean:
	rm -f ${MODULE_NAME}.so ${MODULE_NAME}.h

.PHONY: build clean
