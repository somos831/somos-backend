package handlers

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/events", HttpHandler(s.ListAllEvents)).Methods("GET")
	s.Router.HandleFunc("/events/{id}", HttpHandler(s.GetEvent)).Methods("GET")
	s.Router.HandleFunc("/events", HttpHandler(s.CreateEvent)).Methods("POST")
	s.Router.HandleFunc("/events/{id}", HttpHandler(s.UpdateEvent)).Methods("PATCH")
	s.Router.HandleFunc("/events/{id}", HttpHandler(s.DeleteEvent)).Methods("DELETE")

	s.Router.HandleFunc("/categories", HttpHandler(s.ListAllCategories)).Methods("GET")
	s.Router.HandleFunc("/categories/{id}", HttpHandler(s.GetCategory)).Methods("GET")
	s.Router.HandleFunc("/categories", HttpHandler(s.CreateCategory)).Methods("POST")
	s.Router.HandleFunc("/categories/{id}", HttpHandler(s.UpdateCategory)).Methods("PATCH")
	s.Router.HandleFunc("/categories/{id}", HttpHandler(s.DeleteCategory)).Methods("DELETE")

	s.Router.HandleFunc("/users", s.CreateUser).Methods("POST")
	s.Router.HandleFunc("/users/{id}", s.GetUserByID).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.DeleteUser).Methods("DELETE")
	s.Router.HandleFunc("/users/{id}", s.UpdateUser).Methods("PUT")
}
