package controllers

import (
	"net/http"

	"github.com/khalidsaifuddin/pmo/api/models"
	"github.com/khalidsaifuddin/pmo/api/responses"
)

func (server *Server) GetJenisPendaftarans(w http.ResponseWriter, r *http.Request) {

	responses.EnableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	jenis_pendaftaran := models.JenisPendaftaran{}

	jenis_pendaftarans, err := jenis_pendaftaran.FindAllJenisPendaftaran(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, jenis_pendaftarans)
}

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }
