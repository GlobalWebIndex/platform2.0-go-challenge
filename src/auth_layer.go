package main

import "net/http"

// an authentication for admin users
func (db *dbMock) adminCheck(u string, p string) bool {
	pgt, found := db.adminCreds[u]
	if !found {
		return false
	}
	if pgt != p {
		return false
	}
	return true
}

// an authentication for simple users
func (db *dbMock) userCheck(u string, p string) bool {
	pgt, found := db.DBgetUserPassword(u)

	if !found {
		return false
	}
	if pgt != p {
		return false
	}
	return true
}

// adminauth auth handler for admin users
func (db *dbMock) adminauth(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !db.adminCheck(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		f(w, r)
	}
}

// userauth auth handler for simple users
func (db *dbMock) userauth(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !db.userCheck(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		f(w, r)
	}
}
