package dscribe

import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Init() {
	connective_words, err := ioutil.ReadFile("/usr/share/dict/connectives")
    check(err)
    fmt.Print(string(connective_words))
}