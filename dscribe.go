package dscribe

import (
	"io/ioutil"
	s "strings"
	"math/rand"
	"time"
)


var connectives []byte
var lastNames []byte
var words []byte

var aConnectives []string
var aLastNames []string
var aWords []string

var err error


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Generate() string {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	
	i_lastNames := random.Intn(len(aLastNames))
	i_words := random.Intn(len(aWords))
	i_connectives := random.Intn(len(aConnectives))

	theWord := string(aConnectives[i_connectives]) + "-" + string(aLastNames[i_lastNames]) + "-" + string(aWords[i_words])

	return s.ToLower(theWord)

}

func Init() {

	connectives, err = ioutil.ReadFile("connectives.txt")
	check(err)
	aConnectives = s.Split(string(connectives), "\n")

	words, err = ioutil.ReadFile("words.txt")
	check(err)
	aWords = s.Split(string(words), "\n")

	lastNames, err = ioutil.ReadFile("last-names.txt")
	check(err)
	aLastNames = s.Split(string(lastNames), "\n")

}