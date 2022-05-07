package faker

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// Create structs with random injected data
type Foo struct {
	Str           string         `json:"str"`
	Int           int            `json:"int"`
	Pointer       *int           `json:"pointer"`
	Sentence      string         `json:"sentence" fake:"{sentence:3}"` // Can call with parameters
	RandStr       string         `json:"randStr" fake:"{randomstring:[hello,world]}"`
	Number        string         `json:"number" fake:"{number:1,10}"`      // Comma separated for multiple values
	Regex         string         `json:"regex" fake:"{regex:[abcdef]{5}}"` // Generate string from regex
	Map           map[string]int `fakesize:"2"`
	Array         []string       `fakesize:"2"`
	Bar           Bar            `json:"bar"`
	Skip          *string        `fake:"skip"` // Set to "skip" to not generate data for
	Created       time.Time      // Can take in a fake tag as well as a format tag
	CreatedFormat time.Time      `fake:"{year}-{month}-{day}" format:"2006-01-02"`
	Details       Details        `json:"details"`
}

type Details struct {
	Name     string
	LastName string
	SSN      string
	Street   []string `fakesize:"10"`
}

type Bar struct {
	Name   string
	Number int
	Float  float32
}

type Person struct {
	Name     string    `json:"name"`
	LastName string    `json:"lastName" `
	SSN      int       `json:"SSN" fake:"skip"`
	Date     time.Time `json:"time" `
	Street   []string  `json:"street"`
}

type FakerMap map[int]interface{}

type Faker interface {
	faker() interface{}
}

func fooFaker() Foo {
	// Pass your struct as a pointer
	var f Foo
	gofakeit.Struct(f)
	return f
}
func personFaker() Person {
	// Pass your struct as a pointer
	var p Person
	gofakeit.Struct(&p)
	return p
}

func CreateUniqueFactory(end int) func() int {
	var start = 1
	return func() int {
		var current = start
		start = start + 1
		return current
	}
}

func RunFaker(loops int) map[int]Person {
	var t = make(map[int]Person)
	var getNextUnique = CreateUniqueFactory(loops)
	for i := 1; i < loops; i++ {
		person := personFaker()
		person.SSN = getNextUnique()
		person.Date = gofakeit.Date()
		t[i] = person
	}
	return t
}

/*func personUniqueSSN(loops int, data map[int]Person) map[int]Person {
	var getNextUnique = CreateUniqueFactory(loops)
	for i := 1; i < loops; i++ {
		var person = data[i]
		person.SSN = getNextUnique()
		data[i] = person
	}
	return data
}*/
