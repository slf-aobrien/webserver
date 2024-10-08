package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("Server starting Starting")
	mux := http.NewServeMux()

	//endpoints
	mux.HandleFunc("/", loadFromRoot)
	mux.HandleFunc("/assets/", loadFromAsset)
	mux.HandleFunc("/services/", services)

	err := http.ListenAndServe(":8880", mux)
	check(err)
	log.Println("Server Ending")
}

func loadFromAsset(w http.ResponseWriter, r *http.Request) {
	anAsset := r.URL.Path[len("/assets/"):]
	httpFileLoaderSupportingGET(w, r, "./ui/dist/assets/"+anAsset)
}

func loadFromRoot(w http.ResponseWriter, r *http.Request) {
	aFile := r.URL.Path[len("/"):]
	if len(aFile) == 0 {
		aFile = "index.html"
	}
	httpFileLoaderSupportingGET(w, r, "./ui/dist/"+aFile)
}

func services(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPatch:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	case http.MethodOptions:
		w.Header().Set("Allow", "GET, PATCH, POST, PUT, DELTE, OPTIONS")
		w.WriteHeader(http.StatusNoContent)
	default:
		w.Header().Set("Allow", "GET, PATCH, POST, PUT, DELTE, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func httpFileLoaderSupportingGET(w http.ResponseWriter, r *http.Request, fileName string) {
	switch r.Method {
	case http.MethodGet:
		file, err := os.Open(fileName)
		check(err)
		defer file.Close()
		fileInfo, err := file.Stat()
		check(err)

		http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
	case http.MethodOptions:
		w.Header().Set("Allow", "GET, OPTIONS")
		w.WriteHeader(http.StatusNoContent)
	default:
		w.Header().Set("Allow", "GET, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func readFile() {
	dat, err := os.ReadFile("/Users/SM58/Development/Digix/digix/README.md")
	check(err)
	content := string(dat)
	lines := strings.Split(content, "\n")
	log.Println(lines[4])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
