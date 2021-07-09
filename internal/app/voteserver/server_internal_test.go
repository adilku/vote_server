package voteserver

import (
	"bytes"
	"encoding/json"
	"github.com/adilku/vote_server/internal/app/store/simplestore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handlePollCreate(t *testing.T) {
	s := newServer(simplestore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string][]string{
				"fields": {"black", "white", "yellow"},
			},
			expectedCode: http.StatusCreated,
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
			req, _ := http.NewRequest(http.MethodPost, "/createPoll", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
	/*
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/polls", nil)
		s := newServer(simplestore.New())
		s.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)
	*/
}
