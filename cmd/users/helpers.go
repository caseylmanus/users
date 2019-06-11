package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/caseylmanus/users/dbmodels"
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

func validateUser(user models.PlanetUser) error {
	if user.FirstName == "" {
		return errors.New("first name is required")
	}
	if user.LastName == "" {
		return errors.New("last name is required")
	}
	if user.UserID == "" {
		return errors.New("user id is required")
	}
	return nil

}
