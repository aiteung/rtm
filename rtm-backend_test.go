package main

import (
	"fmt"
	"os"
	// "rtm/controller"
	"testing"
)

var MariaStringAkademik string = os.Getenv("MARIASTRINGAKADEMIK")

func TestPrintenv(t *testing.T) {
	fmt.Println(MariaStringAkademik)
}

// func TestInsertdata(t *testing.T) {
// 	id := "cc2"
// 	nama := "rofiganteng"
// 	email := "cth@gmail.com"
// 	hp := "000000"
// 	controller.InsertData()
// }
