package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// logWebServer a wrapper to log to server
func logWebServer(msg string) {
	log.New(os.Stdout, "http: ", log.LstdFlags).Println(msg)
}

// aux for LogHTTP
type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

// aux for LogHTTP
func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// aux for LogHTTP
func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

// LogHTTP  a nice wrapper for the handler so that the response can be accessed
func LogHTTP(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		handler.ServeHTTP(&sw, r)
		logWebServer(fmt.Sprintf("Request:%q, Duration:%v, Status:%d", html.EscapeString(r.URL.Path), time.Now().Sub(start), sw.status))
	}
}

// validationChecks auxiliary helper function for similar requests
// returns an asset that is requested to perform actions on it
// in case sth is wrong the asset is nil with the respective http status
func validationChecks(r *http.Request) (int, *Asset) {
	// Get params
	params := mux.Vars(r)

	// check param format
	userID, err := strconv.Atoi(params["userid"])
	if err != nil {
		return http.StatusBadRequest, nil
	}

	// check param format
	assetID, err := strconv.Atoi(params["assetid"])
	if err != nil {
		return http.StatusBadRequest, nil
	}

	// check asset existence
	asset, found := DBgetAssetByID(assetID)
	if !found {
		return http.StatusNotFound, nil
	}

	// check asset ownwership
	if getUserId(asset) != userID {
		return http.StatusNotFound, nil
	}

	// check if the request is for the same user
	// a user with correct user and password should not be able to perform actions for another user/asset
	un, _, _ := r.BasicAuth()

	ugt, found := DBgetUserNameByID(userID)
	if !found {
		// theres no such user
		return http.StatusNotFound, nil
	}
	if ugt != un {
		return http.StatusUnauthorized, nil
	}

	return http.StatusOK, asset
}
