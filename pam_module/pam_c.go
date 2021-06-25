package main

/*
#include <errno.h>
#include <pwd.h>
#include <security/pam_appl.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

char* string_from_argv(int i, char** argv)
{
  return strdup(argv[i]);
}

char* get_user(pam_handle_t* pamh)
{
  if (!pamh)
    return NULL;

  int pam_err = 0;
  const char* user;

  if((pam_err = pam_get_item(pamh, PAM_USER, (const void**)&user)) != PAM_SUCCESS)
    return NULL;

  return strdup(user);
}

char* get_password(pam_handle_t* pamh)
{
  if (!pamh)
    return NULL;

  int pam_err = 0;
  const char* prompt;

  if((pam_err = pam_get_item(pamh, PAM_AUTHTOK, (const void**)&prompt)) != PAM_SUCCESS)
    return NULL;

  return strdup(prompt);
}
*/
import "C"
