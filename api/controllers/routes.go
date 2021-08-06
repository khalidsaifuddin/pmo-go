package controllers

import "github.com/khalidsaifuddin/pmo/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	// s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// //Users routes
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// //Posts routes
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//Registration routes
	s.Router.HandleFunc("/registration", middlewares.SetMiddlewareJSON(s.CreateRegistration)).Methods("POST", "OPTIONS")

	//Wilayah routes
	s.Router.HandleFunc("/wilayah", middlewares.SetMiddlewareJSON(s.GetWilayahs)).Methods("GET")
	s.Router.HandleFunc("/wilayah_individu/{kode}", middlewares.SetMiddlewareJSON(s.GetWilayah)).Methods("GET")
	s.Router.HandleFunc("/wilayah/{id_level_wilayah}", middlewares.SetMiddlewareJSON(s.GetWilayahsByInduk)).Methods("GET")
	s.Router.HandleFunc("/wilayah/{id_level_wilayah}/{induk_kode}", middlewares.SetMiddlewareJSON(s.GetWilayahsByInduk)).Methods("GET")

	//JenisPendaftaran routes
	s.Router.HandleFunc("/jenis_pendaftaran", middlewares.SetMiddlewareJSON(s.GetJenisPendaftarans)).Methods("GET")

}
