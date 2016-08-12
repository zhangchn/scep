package scepclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"golang.org/x/net/context"
)

func TestGetCACaps(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(dump))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	client := NewClient(server.URL + "/scep")
	ctx := context.Background()

	resp, err := client.GetCACaps(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(resp))
}

func TestGetCACert(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(dump))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	client := NewClient(server.URL + "/scep")
	ctx := context.Background()

	resp, _, err := client.GetCACert(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(resp))
}

func TestPKIOperationGET(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(dump))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	client := NewClient(server.URL + "/scep")
	ctx := context.Background()

	resp, err := client.PKIOperation(ctx, []byte("foobarbaz"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(resp))
}
