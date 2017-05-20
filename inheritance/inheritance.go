/* Example of how to reuse existing structs with other structs
   Not exactly the same as object oriented inheritace, possibly better as it is
   simpler */

package main

import (
	"fmt"
)

type structA struct {
	x int
	s string
}

type structB struct {
	y int
	s string
}

type structCombo struct {
	structA  // reuse struct members from existing struct..
	structB  // ...
	z int
}

func main() {
	var combo structCombo
	combo.x = 10
	combo.y = 20
	combo.z = 30
	fmt.Println("Values: ", combo.x, ",", combo.y, ",", combo.z,",")

    // combo.s = "data" // note: cannot do this, as it conflicts/is duplicate
	// fmt.Println(combo.s)
}