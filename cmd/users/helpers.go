package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func writeJSON(w http.ResponseWriter, val interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(val)

}
func writeErr(w http.ResponseWriter, err error, status int) {
	logrus.Error(err)
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}
