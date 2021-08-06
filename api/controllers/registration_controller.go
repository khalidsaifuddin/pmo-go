package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khalidsaifuddin/pmo/api/models"
	"github.com/khalidsaifuddin/pmo/api/responses"
	"github.com/khalidsaifuddin/pmo/api/utils/formaterror"
)

func (server *Server) CreateRegistration(w http.ResponseWriter, r *http.Request) {
	responses.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	registration := models.Registration{}
	err = json.Unmarshal(body, &registration)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	registration.Prepare()
	err = registration.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	registrationCreated, err := registration.SaveRegistration(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, registrationCreated.ID))
	responses.JSON(w, http.StatusCreated, registrationCreated)
}

func (server *Server) GetRegistration(w http.ResponseWriter, r *http.Request) {
	responses.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	registration := models.Registration{}
	vars := mux.Vars(r)
	id := vars["id"]

	registrations, err := registration.FindRegistrationByID(server.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, registrations)
}
