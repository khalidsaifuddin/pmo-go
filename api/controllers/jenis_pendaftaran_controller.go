package controllers

import (
	"net/http"

	"github.com/khalidsaifuddin/pmo/api/models"
	"github.com/khalidsaifuddin/pmo/api/responses"
)

func (server *Server) GetJenisPendaftarans(w http.ResponseWriter, r *http.Request) {

	jenis_pendaftaran := models.JenisPendaftaran{}

	jenis_pendaftarans, err := jenis_pendaftaran.FindAllJenisPendaftaran(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, jenis_pendaftarans)
}
