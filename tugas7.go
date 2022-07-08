package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	go pesan1()
	pesan2()
}

func pesan1() {
	var number = 494
	var rnumber = reflect.ValueOf(number)
	fmt.Println(rnumber.Type())
}

func pesan2() {
	var nama = "Fatah"
	var rnama = reflect.ValueOf(nama)
	fmt.Println(rnama.Type())
}
