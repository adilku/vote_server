package voteserver

import (
	"encoding/json"
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/checkBalance", s.handleCheckBalance()).Methods("POST")
	s.router.HandleFunc("/changeBalance", s.handleChangeBalance()).Methods("POST")
}


func (s *server) handleCheckBalance() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if _, err := s.store.GetWallet().FindById(req.Id); err != nil {
			wallet := model.Wallet{
				ID: req.Id,
				Balance: 0,
				}
			if err := s.store.GetWallet().Create(&wallet); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			s.respond(w, r, http.StatusCreated, wallet)
		} else {
			s.respond(w, r, http.StatusOK, nil)
		}
	}
}

func (s *server) handleChangeBalance() http.HandlerFunc {
	type request struct {
		Id    int `json:"id"`
		Delta int `json:"delta"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.GetWallet().ChangeBalance(req.Id, req.Delta); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

/*
func (s *server) handlePoll() http.HandlerFunc {
	type request struct {
		PollId int    `json:"poll_id"`
		Choice string `json:"choice"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		_, err := s.store.GetWallet().FindById(req.PollId)
		if err != nil {
			s.error(w, r, http.StatusNoContent, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}
*/

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
