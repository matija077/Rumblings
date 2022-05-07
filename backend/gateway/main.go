package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf(" ia ma a gaetway modulke")

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	var servicesModel = createServicesModel()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		frontend, ok := servicesModel.getFrontend(r.URL.Path)
		if !ok {
			return
		}
		stringFrontend := fmt.Sprintf("%#v", frontend)
		fmt.Fprintf(w, "frontend is: %q", html.EscapeString(stringFrontend))
		// TODO not working
		servicesModel.redirect(&w, r, frontend)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		if file == nil {
			log.Fatal("log file not found")
			w.Write([]byte("File not found"))
		}

		write := make([]byte, 100)
		var offset int64
		var fileContents = make([]byte, 100)
		for {
			_, err := (*file).ReadAt(write, offset)
			if err != nil {
				if err == io.EOF {
					break
				}

				log.Fatal("error reading a file: " + err.Error())
			}

			offset = offset + 100
		}

		defer fmt.Println()

		w.Write(fileContents)

	})

	err = http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8090", nil))
}
