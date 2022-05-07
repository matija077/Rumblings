package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	for bytes := range readFileLines("./slips.json") {
		line := string(bytes)
		fmt.Println(&line)
		fmt.Println("vani")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6060", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func readFileLines(fileName string) chan []byte {
	lines := make(chan []byte, 64)
	func() {
		defer closeL(lines)
		f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal("no such file")
		}
		defer closeC(f)
		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadBytes('\n')
			if err != nil {
				return
			}
			fmt.Println("ovdje sam")
			lines <- line[:len(line)-1] // trim newline, last char
		}
	}()
	return lines
}

func closeL(lines chan []byte) {
	fmt.Println("close lines")
	close(lines)
}

func closeC(f *os.File) {
	fmt.Println("close channel")
	f.Close()
}

/*func ctxToSig(ctx) (chan struct{}, chan struct{}) {
	sig := make(chan struct{})
	exit := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			close(sig)
		case <-exit:
		}
	}()
	return sig, exit
}*/
