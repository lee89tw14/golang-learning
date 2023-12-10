package note

import "fmt"

// 2.1 EscapedCharacters
func EscapedCharacters() {
	fmt.Println("\"hello world\"")
	fmt.Println("hello\\world")
}

const (
	Version int = 11110
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

// 2.3 Basic Data Types
func BasicDataTypes() {
	fmt.Println("\n2.3.1 整數型")
	var (
		n1        = 0b0110
		n2 int8   = 0o70
		n3 uint16 = 0xFF
	)
	fmt.Printf("n1=%v,type is %T\n", n1, n1)
	fmt.Printf("n2=%v,type is %T\n", n2, n2)
	fmt.Printf("n3=%v,type is %T\n", n3, n3)

	fmt.Println("\n2.3.2 浮點型")
	var (
		f1         = 1.0
		f3 float32 = 1
	)
	fmt.Printf("f1=%v,type is %T\n", f1, f1)
	fmt.Printf("f3=%v,type is %T\n", f3, f3)

	fmt.Println("\n2.3.4 數值類型轉換")
	n2 = int8(n3)
	fmt.Printf("n2=%v\n", n2) //精度需注意

	fmt.Println("\n2.3.5 字符型ASCII/UTF8")
	var (
		c1 byte
		c2 = '0' //雙引號爲字符串（即文字列）單引號爲字符
		c3 = '中'
	)
	fmt.Printf("c1的碼值=%v,碼值對應的字符是%c,type is %T\n", c1, c1, c1)
	fmt.Printf("c2的碼值=%v,碼值對應的字符是%c,type is %T\n", c2, c2, c2)
	fmt.Printf("c3的碼值=%v,碼值對應的字符是%c,type is %T\n", c3, c3, c3)

}

//2.4

//2.5
