package main

import "net/http"

// an authentication for admin users
func adminCheck(u string, p string) bool {
	pgt, found := adminCreds[u]
	if !found {
		return false
	}
	if pgt != p {
		return false
	}
	return true
}

// an authentication for simple users
func userCheck(u string, p string) bool {
	pgt, found := DBgetUserPassword(u)

	if !found {
		return false
	}
	if pgt != p {
		return false
	}
	return true
}

// adminauth auth handler for admin users
func adminauth(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !adminCheck(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		f(w, r)
	}
}

// userauth auth handler for simple users
func userauth(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !userCheck(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		f(w, r)
	}
}
