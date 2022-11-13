package main

import "fmt"

var outsideInt int = 10
var outsideString = "outer str"

//outsideNoVar := "outside no var"  this is error as := no var style is not allowed outside a function

const PulicConst float32 = 1234567.890

func main() {
	fmt.Println("Variables")

	//string
	var username string = "Ashamol CM"
	fmt.Println("UserName:", username)
	fmt.Printf("Type: %T \n", username)

	//bool
	var isLoggedIn bool = true
	fmt.Println("isLoggedIn:", isLoggedIn)
	fmt.Printf("Type: %T \n", isLoggedIn)

	//uint8
	var unsignedInt8 uint8 = 255
	fmt.Println("unsignedInt8:", unsignedInt8)
	fmt.Printf("Type: %T \n", unsignedInt8)

	//byte - alias uint8
	var unsignedInt8AsByte byte = 255
	fmt.Println("unsignedInt8AsByte:", unsignedInt8AsByte)
	fmt.Printf("Type: %T \n", unsignedInt8AsByte)

	//int8
	var signedInt8 int8 = -128
	fmt.Println("signedInt8:", signedInt8)
	fmt.Printf("Type: %T \n", signedInt8)

	//rune alias int32
	var signedInt32AsRune rune = -128
	fmt.Println("signedInt32AsRune:", signedInt32AsRune)
	fmt.Printf("Type: %T \n", signedInt32AsRune)

	//int (int is never alias to int32)
	var normalInt int = 255
	fmt.Println("normalInt:", normalInt)
	fmt.Printf("Type: %T \n", normalInt)

	//float32
	var floating32 float32 = 255.9998888997778890000
	fmt.Println("floating8:", floating32)
	fmt.Printf("Type: %T \n", floating32)

	//float64
	var floating64 float64 = 255.9998888997778890000
	fmt.Println("floating64:", floating64)
	fmt.Printf("Type: %T \n", floating64)

	//implicit type
	var varWithImplicitlyIdentifiedType = "some string"
	fmt.Println("varWithImplicitlyIdentifiedType", varWithImplicitlyIdentifiedType)
	fmt.Printf("Type: %T", varWithImplicitlyIdentifiedType)
	//varWithImplicitlyIdentifiedType = 1; error comes since this si already considered string and cannot hold int now.

	//no var style (assumes to int, float64, string etc. which are more generic to hold the data assigned.)
	varWithNoVar := -255
	fmt.Println("varWithNoVar:", varWithNoVar)
	fmt.Printf("Type: %T \n", varWithNoVar)
	//varWithNoVar = "Sample Str" error since type is already int

	fmt.Printf("%T %v", PulicConst, PulicConst)

}
