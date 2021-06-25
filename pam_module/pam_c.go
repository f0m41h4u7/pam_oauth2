package main

/*
#cgo LDFLAGS: -lpam -fPIC

#include <errno.h>
#include <pwd.h>
#include <security/pam_appl.h>
#include <security/pam_misc.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <sys/prctl.h>

char* string_from_argv(int i, char** argv)
{
  return strdup(argv[i]);
}

char* get_user(pam_handle_t *pamh)
{
  if (!pamh)
    return NULL;

  int pam_err = 0;
  const char* user;

  if((pam_err = pam_get_item(pamh, PAM_USER, (const void**)&user)) != PAM_SUCCESS)
    return NULL;

  return strdup(user);
}

int disable_ptrace()
{
  return prctl(PR_SET_DUMPABLE, 0);
}
*/
import "C"
import (
	"os"
	"unsafe"
)

func init() {
	if !disablePtrace() {
		pamLog("unable to disable ptrace")
	}
}

func sliceFromArgv(argc C.int, argv **C.char) []string {
	r := make([]string, 0, argc)
	for i := 0; i < int(argc); i++ {
		s := C.string_from_argv(C.int(i), argv)
		defer C.free(unsafe.Pointer(s))
		r = append(r, C.GoString(s))
	}
	return r
}

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	cUsername := C.get_user(pamh)
	if cUsername == nil {
		return C.PAM_USER_UNKNOWN
	}
	defer C.free(unsafe.Pointer(cUsername))

	r := pamAuthenticate(os.Stderr, C.GoString(cUsername), sliceFromArgv(argc, argv))
	if r == AuthError {
		return C.PAM_AUTH_ERR
	}

	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	return C.PAM_IGNORE
}
