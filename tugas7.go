package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var tipestring = "Halo"
	var tipeint = 10
	go bacatipe2(tipestring)
	bacatipe1(tipeint)
	var input string
	fmt.Scanln(&input)

}

func bacatipe1(tipeint int) {
	var reflectnumber = reflect.ValueOf(tipeint)
	if reflectnumber.Kind() == reflect.Int {
		fmt.Println("tipe data ini adalah int")
	}
}

func bacatipe2(tipestring string) {
	var reflectnumber = reflect.ValueOf(tipestring)
	if reflectnumber.Kind() == reflect.String {
		fmt.Println("tipe data ini adalah string")
	}
}

// func tampil_pesan(x int, pesan string) {
// 	for i := 0; i < x; i++ {
// 		fmt.Println((i + 1), pesan)
// 	}
// }
