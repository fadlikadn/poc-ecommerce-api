package controllers

import "github.com/fadlikadn/poc-ecommerce-api/api/middlewares"

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	s.Router.HandleFunc("/transaction", middlewares.SetMiddlewareJSON(s.CreateTransaction)).Methods("POST")
	s.Router.HandleFunc("/transaction-details/{id}", middlewares.SetMiddlewareJSON(s.AddTransactionDetail)).Methods("POST")
	s.Router.HandleFunc("/transaction-by-customer/{id}", middlewares.SetMiddlewareJSON(s.GetTransactionByCustomer)).Methods("GET")

}
