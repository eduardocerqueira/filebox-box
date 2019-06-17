package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestApiTest(t *testing.T) {
	expected := []byte("test")
	req, err := http.NewRequest(http.MethodGet, apiTest, nil)

	checkError(err, t)

	res := httptest.NewRecorder()
	test(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v want 200", res.Code)
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf("Response body was '%v' want '%v'", expected, res.Body)
	}
}

func TestApiVersion(t *testing.T) {
	expected := []byte("version")
	req, err := http.NewRequest(http.MethodGet, apiVersion, nil)

	checkError(err, t)

	res := httptest.NewRecorder()
	version(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v want 200", res.Code)
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf("Response body was '%v' want '%v'", expected, res.Body)
	}
}

func TestApiUpload(t *testing.T) {
	// TODO: upload test-file

	file, err := os.Open("./test-file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	expected := []byte("version")
	req, err := http.NewRequest(http.MethodPost, apiUpload, nil)

	checkError(err, t)

	res := httptest.NewRecorder()
	upload(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v want 200", res.Code)
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf("Response body was '%v' want '%v'", expected, res.Body)
	}
}
