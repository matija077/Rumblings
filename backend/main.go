package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type TRequest struct {
	id uint32
}

type TRequestProcessor struct {
	counter  int16
	requests []*TRequest
}

var requestProcessor *TRequestProcessor

func createRequestProcessor() {
	if requestProcessor == nil {
		requestProcessor = &TRequestProcessor{0, make([]*TRequest, 0, 2)}
	}
}

func main() {

	//var wg sync.WaitGroup

	/*http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}*/

	//slicePlay(&wg)
	//wg.Wait()

	go createRequestProcessor()

	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	/*file2, err := os.OpenFile("/logging/logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(file2)
	}*/

	log.SetOutput(file)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		/*if reqHeadersBytes, err := json.Marshal(r.Header); err != nil {
			log.Println("Could not Marshal Req Headers")
		} else {
			log.Println(string(reqHeadersBytes))
		}*/

		w.Header().Set("access-control-allow-origin", "*")

		go handle(w, r)
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

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {

	log.Printf("counter: " + fmt.Sprintf("%v", requestProcessor.counter))

	if requestProcessor.counter < 1 {
		requestProcessor.counter++

		log.Printf("logged a succes request")
		//var duration = time.Duration(1) * time.Millisecond
		//time.Sleep(duration)

		w.Write([]byte(`{"key": "success"}`))

		requestProcessor.counter--
		return
	}

	log.Printf("logged a get request, not executing it")
	w.Write([]byte(`{"key": "failed request"}`))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recived: ", r.Method)
	fmt.Println("ajmo")
	var a string = "aloooo"
	w.Write([]byte(a))

	switch r.Method {
	case http.MethodGet:
		log.Println("get: ")
		break
	default:
		log.Println("Wnt wrong: ")
		break
	}
}

func slicePlay(wg *sync.WaitGroup) {
	var slice = make([]int, 6, 10)
	//var nilSlice = []int{}
	slice = []int{1, 2, 3, 4, 5, 6}
	slice2 := clone(slice)
	slice = append(slice, 7)
	//fmt.Println(nilSlice == nil)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println("----")

	fmt.Println(slice[4+2:])

	//(*wg).Done()

	slice = remove(slice, 4, 2, true)
	fmt.Println(slice)

}

func clone(originalSlice []int) (clonedSlice []int) {
	if originalSlice != nil {
		clonedSlice = make([]int, len(originalSlice))

		copy(clonedSlice, originalSlice)
	}

	return clonedSlice
}

func remove(slice []int, start int, count int, preserveOrder bool) (returnSlice []int) {
	returnSlice = make([]int, 0, len(slice)-count)

	returnSlice = append(returnSlice, slice[:start]...)
	returnSlice = append(returnSlice, slice[start+count:]...)

	fmt.Println(returnSlice)

	return returnSlice
}
