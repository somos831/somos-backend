package handlers

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("GET /events", HttpHandler(s.ListAllEvents))
	s.Router.HandleFunc("GET /events/{id}", HttpHandler(s.GetEvent))
	s.Router.HandleFunc("POST /events", HttpHandler(s.CreateEvent))
	s.Router.HandleFunc("PATCH /events/{id}", HttpHandler(s.UpdateEvent))
	s.Router.HandleFunc("DELETE /events/{id}", HttpHandler(s.DeleteEvent))

	s.Router.HandleFunc("GET /categories", HttpHandler(s.ListAllCategories))
	s.Router.HandleFunc("GET /categories/{id}", HttpHandler(s.GetCategory))
	s.Router.HandleFunc("POST /categories", HttpHandler(s.CreateCategory))
	s.Router.HandleFunc("PATCH /categories/{id}", HttpHandler(s.UpdateCategory))
	s.Router.HandleFunc("DELETE /categories/{id}", HttpHandler(s.DeleteCategory))

	s.Router.HandleFunc("POST /users", s.CreateUser)
	s.Router.HandleFunc("GET /users/{id}", s.GetUserByID)
	s.Router.HandleFunc("DELETE /users/{id}", s.DeleteUser)
	s.Router.HandleFunc("PUT /users/{id}", s.UpdateUser)
}
