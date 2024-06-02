package handlers

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/events", s.ListEvents).Methods("GET")
	s.Router.HandleFunc("/events/{id}", s.GetEvent).Methods("GET")
	s.Router.HandleFunc("/events", s.CreateEvent).Methods("POST")
	s.Router.HandleFunc("/events/{id}", s.UpdateEvent).Methods("PATCH")
	s.Router.HandleFunc("/events/{id}", s.DeleteEvent).Methods("DELETE")

	s.Router.HandleFunc("/categories", s.ListAllCategories).Methods("GET")

	s.Router.HandleFunc("/locations", s.CreateLocation).Methods("POST")

	s.Router.HandleFunc("/users", s.CreateUser).Methods("POST")
	s.Router.HandleFunc("/users/{id}", s.GetUserByID).Methods("GET")
	s.Router.HandleFunc("/users/{id}", s.DeleteUser).Methods("DELETE")
	s.Router.HandleFunc("/users/{id}", s.UpdateUser).Methods("PUT")
}
