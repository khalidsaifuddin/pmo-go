package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khalidsaifuddin/pmo/api/models"
	"github.com/khalidsaifuddin/pmo/api/responses"
)

func (server *Server) GetWilayahs(w http.ResponseWriter, r *http.Request) {

	wilayah := models.Wilayah{}

	wilayahs, err := wilayah.FindAllWilayah(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, wilayahs)
}

func (server *Server) GetWilayah(w http.ResponseWriter, r *http.Request) {
	responses.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	wilayah := models.Wilayah{}
	vars := mux.Vars(r)
	kode := vars["kode"]

	wilayahs, err := wilayah.FindWilayahByID(server.DB, kode)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, wilayahs)
}

func (server *Server) GetWilayahsByInduk(w http.ResponseWriter, r *http.Request) {
	responses.EnableCors(&w)

	vars := mux.Vars(r)
	id_level_wilayah, err := strconv.ParseUint(vars["id_level_wilayah"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	induk_kode, ok := vars["induk_kode"]

	wilayah := models.Wilayah{}
	wilayahs := &[]models.Wilayah{}

	if ok {
		wilayahs, err = wilayah.FindWilayahByInduk(server.DB, id_level_wilayah, induk_kode)
	} else {
		wilayahs, err = wilayah.FindWilayahByInduk(server.DB, id_level_wilayah, "")
	}

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, wilayahs)

}
