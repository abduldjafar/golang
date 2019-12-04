package main

import (
	"backend-code/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "/Users/abdulharisdjafar/go/src/backend-code/controller/TestPath/main.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		data, _ := ioutil.ReadAll(r.Body)
		asString := string(data)

		fmt.Println(asString)
		util.RespondError(w, http.StatusBadGateway, asString)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}