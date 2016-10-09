package serer

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Server struct {
  Red        []string
	Blue       []string
  Champs     []string
}



func (s *Server) Listen() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.Home)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	r.HandleFunc("/link/add", s.AddChamp)
	http.Handle("/", r)
	fmt.Println("The server is now live @ localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// CSS loads style into the page
func CSS(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/css/", http.FileServer(http.Dir("css/")))
}

// Add Champ makes possible to add champs into teams
func (s *Server) AddChamp() {

}

// Home loads the assets
func (s *Server) Home()  {

}
