package main

import (
	"net/http"
	"log"
	"os"
	"io"
	"encoding/json"
)

var listenHost string = ":3001"
var testingPath string = "/request/test/"
var fileHeader string = "request_header.txt"
var fileFull string = "request_full.txt"

type resp struct {
	HeaderFile string `json:"header"`
	BodyFile   string `json:"full"`
	Message    string `json:"message"`
}

func init() {
	if (os.Getenv("WEBDUMP_LISTEN") != "") {
		listenHost = os.Getenv("WEBDUMP_LISTEN")
	}
	if (os.Getenv("WEBDUMP_TESTING_PATH") != "") {
		testingPath = os.Getenv("WEBDUMP_TESTING_PATH")
	}
	if (os.Getenv("WEBDUMP_FILE_HEADER") != "") {
		fileHeader = os.Getenv("WEBDUMP_FILE_HEADER")
	}
	if (os.Getenv("WEBDUMP_FILE_FULL") != "") {
		fileFull = os.Getenv("WEBDUMP_FILE_FULL")
	}
}

func main() {
	http.HandleFunc(testingPath, postHandler)
	http.HandleFunc("/"+fileHeader, getHandlerFileHeader)
	http.HandleFunc("/"+fileFull, getHandlerfileFull)
	log.Println("Listening " + listenHost + " ...")
	http.ListenAndServe(listenHost, nil)
}

func getHandlerFileHeader(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.Method, r.URL, r.Proto)
	file, err := os.Open(fileHeader)
	if err != nil {
		log.Println(r)
	}
	defer file.Close()
	io.Copy(w, file)
}

func getHandlerfileFull(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v", r.Method, r.URL, r.Proto)
	file, err := os.Open(fileFull)
	if err != nil {
		log.Println(r)
	}
	defer file.Close()
	io.Copy(w, file)
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("%v %v %v", r.Method, r.URL, r.Proto)

	defer r.Body.Close()

	// write header file
	head, err := os.Create(fileHeader)
	if err != nil {
		log.Println(r)
	}
	defer head.Close()
	head.WriteString(formatRequest(r))

	// write header file
	full, err := os.Create(fileFull)
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
		"/" + fileFull,
		"curl http://localhost:3001/"+fileHeader+" or for ngrok curl http://"+r.Host+"/"+fileHeader,
	}
	if err := json.NewEncoder(w).Encode(rjson); err != nil {
		log.Panic(err)
		return
	}

}
