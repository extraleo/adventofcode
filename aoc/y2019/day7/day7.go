package main

import (
	"fmt"
)


var (
	f64 float64     = 0.000000000006
)

func main(){
  fmt.Printf("%g\n", f64)
	fmt.Printf("%v\n", f64)
	fmt.Printf("%f\n", f64)
	fmt.Printf("%e\n", f64)

}

func day() {
	panic("unimplemented")
}