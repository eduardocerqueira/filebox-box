package main

/*
	filebox-box

	Microservice running Gorilla Mux on port 8000 and serving the folling endpoints:

	/v1/test
	/v1/version
	/v1/upload
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const apiTest string = "/test"
const apiVersion string = "/version"
const apiUpload string = "/upload"

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func test(w http.ResponseWriter, r *http.Request) {
	log.Println("test")
	fmt.Fprintf(w, "test")
}

func version(w http.ResponseWriter, r *http.Request) {
	log.Println("version")
	fmt.Fprintf(w, "version")
}

func upload(w http.ResponseWriter, r *http.Request) {
	// curl -vv -X POST -F file=@./test-file.txt http://localhost:8000/v1/upload

	log.Println("upload")

	file, header, err := r.FormFile("file")
	defer file.Close()

	if err != nil {

		return
	}

	out, err := os.Create("/tmp/" + header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Failed to open the file for writing")
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File %s uploaded successfully.", header.Filename)
}

func addRouters(r *mux.Router) {
	r.HandleFunc(apiTest, test).Methods("GET")
	r.HandleFunc(apiVersion, version).Methods("GET")
	r.HandleFunc(apiUpload, upload).Methods("POST")
}

func main() {
	log.Println("filebox-box running MUX Routers")
	log.Println("server is running at: http://localhost:8000/")

	router := mux.NewRouter().StrictSlash(true)
	addRouters(router.PathPrefix("/v1").Subrouter())
	log.Fatal(http.ListenAndServe(":8000", router))
}
