package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "port")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered ", r)
		}
	}()

	http.HandleFunc("/options", func(w http.ResponseWriter, r *http.Request) {
		/*read, err := r.GetBody()
		if err != nil {
			log.Printf("error reading body from options request")
			return
		}

		var data []byte = make([]byte, 100)
		_, err = read.Read(data)
		if err != nil {
			log.Printf("error reading body from read in options reques")
			return
		}

		fmt.Sprintln(string(data))*/

		fmt.Sprintln("HERE")

	})

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
