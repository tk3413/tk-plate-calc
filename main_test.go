package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	impl "github.com/tk3413/tk-weight-calc/server_impl"
)

func TestFullServer_GetWeights(t *testing.T) {
	srv := impl.NewServer()
	mux := http.NewServeMux()
	h := impl.HandlerFromMux(srv, mux)

	ts := httptest.NewServer(h)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/weights?weight=60")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	var got map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if got["5"] != 1 || got["2pt5"] != 1 {
		t.Fatalf("unexpected plate counts: %+v", got)
	}
}

func TestInitLogger(t *testing.T) {
	logger := setupLogger()
	if logger == nil {
		t.Fatal("expected logger to be initialized, got nil")
	}
}
