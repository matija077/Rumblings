package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mongodb/source/mongoDriver"
	"mongodb/source/repo"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	url    string
	port   int
	dbRepo repo.Repo
)

func main() {
	flag.IntVar(&port, "port", 27888, "port")
	flag.StringVar(&url, "url", "127.0.0.1", "url")

	//fakerData := faker.RunFaker(2000000)
	//fmt.Println(fakerData)
	//saveToFile("test", fakerData)

	const url2 = "mongodb://mongoadmin:secret@localhost:27888/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"

	client, err := mongo.NewClient(options.Client().ApplyURI(url2))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	defer cancelContext()

	/*
	   List databases
	*/
	_, err = client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(databases)

	db := client.Database("populate")
	collection := db.Collection("populate3")
	fmt.Println(collection.CountDocuments(ctx, struct{}{}))

	dbRepo = &mongoDriver.MongoDatabaseDriver{
		Collection: collection,
	}

	//databaseDriver.ensureIndex(&ctx, collection)

	//fmt.Println(databaseDriver.FindOne(&ctx, "GUID", "8e6e7190-9da5-446a-aae7-902fb7c21b52"))
	//fmt.Println(databaseDriver.FindOne(&ctx, "IBAN", "tm44NUYn1135853586106483"))

	tempValue := repo.FindOne(dbRepo, &ctx, "name", "erDDmjanM")
	value, _ := tempValue.(string)
	fmt.Println(value)

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/api/form", handleForm)
	http.HandleFunc("/api/search", corsHandler(handleSearch))
	http.ListenAndServe(":3000", nil)

}

type Form struct {
	name     string
	lastName string
	email    string
	ssn      int
	street   string
	date     string
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func corsHandler(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		} else {
			handler(w, r)
		}
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	var form Form
	fmt.Printf("ruta pogodnea")
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Form struct...
	fmt.Fprintf(w, "Form: %+v", form)
}

type Response struct {
	value interface{}
	error string
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ruta pogodnea")
	filter := r.URL.Query()
	key, ok := filter["key"]
	if !ok {
		response := Response{error: "Wrong key"}
		var encodedResponse, _ = json.Marshal(&response)
		w.Write(encodedResponse)
		return
	}

	value, ok := filter["value"]
	fmt.Println(key)
	fmt.Println(value)

	var ctx, _ = context.WithCancel(context.Background())
	res := dbRepo.FindOne(&ctx, key[0], value[0])

	/*person := res.(faker.Person)
	encodedPerson, _ := json.Marshal(&person)
	w.Write(encodedPerson)*/
	fmt.Println(res)

}

func saveToFile(fileName string, data interface{}) {
	/*file, err := os.Create(fileName)
	if err != nil {
		log.Panic("Error opening the file")
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()*/

	byteData, err := json.Marshal(data)
	if err != nil {
		log.Panic("Error marshaling the data")
		return
	}

	err = ioutil.WriteFile(fileName, byteData, 0644)
	if err != nil {
		log.Panic("Error writing the data")
		return
	}
}
