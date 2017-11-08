package main


import (
	"github.com/mickeyyawn/dscribe"
	"fmt"
)

func main(){
	dscribe.Init()
	for i:=0;i<1000;i++ {
		s := dscribe.Generate()

		fmt.Print(s)
		fmt.Println()
	}
	
}