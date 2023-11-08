package note

import "fmt"

// 2.1 EscapedCharacters
func EscapedCharacters() {
	fmt.Println("\"hello world\"")
	fmt.Println("hello\\world")
}

const (
	Version int = 1000
)

// 2.2 Variable and Constants

func VariablesAndConstants() {
	fmt.Println("\n1 Variables")
	var v1 int
	var v2 int = 2
	var v3 = 3
	v1 = 1
	v4 := 4
	var (
		v5     = 5
		v6 int = 6
		v7 int
	)
	fmt.Printf("v1=%v,v2=%v,v3=%v,v4=%v,v5=%v,v6=%v,v7=%v\n", v1, v2, v3, v4, v5, v6, v7)

	fmt.Println("\n2 Constants")
	const (
		c1 = 8
		c2 = iota //the row number, which count from 0
		c3 = iota
		c4 //default value is the row before this one
		c5 = 12
		c6
	)

	fmt.Printf("c1=%v,c2=%v,c3=%v,c4=%v,c5=%v,c6=%v\n", c1, c2, c3, c4, c5, c6)
}
