package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/khalidsaifuddin/pmo/api/models"
	"github.com/khalidsaifuddin/pmo/api/responses"
	"github.com/khalidsaifuddin/pmo/api/utils/formaterror"
)

func (server *Server) CreateRegistration(w http.ResponseWriter, r *http.Request) {
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
