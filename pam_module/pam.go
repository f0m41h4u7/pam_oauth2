package main

/*
#cgo LDFLAGS: -lpam -fPIC

#include <security/pam_appl.h>
#include <stdlib.h>

char* string_from_argv(int, char**);
char* get_user(pam_handle_t*);
char* get_password(pam_handle_t*);
*/
import "C"
import "unsafe"

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	cUsername := C.get_user(pamh)

	if cUsername == nil {
		return C.PAM_USER_UNKNOWN
	}
	defer C.free(unsafe.Pointer(cUsername))

	cPass := C.get_password(pamh)
	if cPass == nil {
		return C.PAM_AUTH_ERR
	}
	defer C.free(unsafe.Pointer(cPass))

	r := pamAuthenticate(C.GoString(cUsername), C.GoString(cPass))
	if r == AuthError {
		return C.PAM_AUTH_ERR
	}

	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	return C.PAM_IGNORE
}
