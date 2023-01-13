package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
)

func Test_main(t *testing.T) {

}
func Test_PlayerServer(t *testing.T) {
	got := "Server starting..."
	want := "Server starting..."

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	request, _ := http.NewRequest(http.MethodGet, "/3000", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)
}
