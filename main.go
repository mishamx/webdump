package main

import (
	"net/http"
	"log"
	"os"
	"io"
	"encoding/json"
)

var listenHost string = ":3001"

const fileHeader = "request_header.txt"
const fileBody = "request_full.txt"

type resp struct {
	HeaderFile string `json:"header"`
	BodyFile   string `json:"full"`
	Message    string `json:"message"`
}

func init() {
	if (os.Getenv("LISTEN") != "") {
		listenHost = os.Getenv("LISTEN")
	}
}

func main() {
	http.HandleFunc("/request/test/", postHandler)
	http.HandleFunc("/"+fileHeader, getHandlerFileHeader)
	http.HandleFunc("/"+fileBody, getHandlerFileBody)
	log.Println("Listening " + listenHost + " ...")
	http.ListenAndServe(listenHost, nil)
}

func getHandlerFileHeader(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(fileHeader)
	if err != nil {
		log.Println(r)
	}
	defer file.Close()
	io.Copy(w, file)
}

func getHandlerFileBody(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(fileBody)
	if err != nil {
		log.Println(r)
	}
	defer file.Close()
	io.Copy(w, file)
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(formatRequest(r))
	defer r.Body.Close()

	// write header file
	head, err := os.Create(fileHeader)
	if err != nil {
		log.Println(r)
	}
	defer head.Close()
	head.WriteString(formatRequest(r))

	// write header file
	full, err := os.Create(fileBody)
	if err != nil {
		log.Println(r)
	}
	defer full.Close()
	full.WriteString(formatRequest(r))
	io.Copy(full, r.Body)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	rjson := resp{
		"/" + fileHeader,
		"/" + fileBody,
		"curl http://localhost:3001/request_header.txt or curl http://localhost:3001/request_full.txt -o request_full.txt",
	}
	if err := json.NewEncoder(w).Encode(rjson); err != nil {
		log.Panic(err)
		return
	}
	
}
