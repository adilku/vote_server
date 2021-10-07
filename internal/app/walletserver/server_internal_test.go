package walletserver

import (
	"bytes"
	"encoding/json"
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store/simplestore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handleCheckBalance(t *testing.T) {
	s := newServer(simplestore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]int{
				"id": 102,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "valid",
			payload: map[string]int{
				"id": 102,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/checkBalance", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}


func TestServer_handleChangeBalance(t *testing.T) {
	s := newServer(simplestore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]int{
				"id": 102,
				"delta": 2,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			payload: map[string]int{
				"id": 102,
				"delta": -2,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			payload: map[string]int{
				"id": 103,
				"delta": -2,
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/changeBalance", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_handleTransaction(t *testing.T) {
	s := newServer(simplestore.New())
	w1 := model.TestWallet(t)
	w2 := model.TestWallet2(t)
	s.store.GetWallet().Create(w1)
	s.store.GetWallet().Create(w2)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]int{
				"IdSender": 101,
				"IdReceiver": 100,
				"delta": 2,
			},
			expectedCode: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/transfer", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
