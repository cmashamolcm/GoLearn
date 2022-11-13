package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the rating for service")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Given rating =", input)

	//fmt.Println("Tampered rating = ", input + 1) cannot add input + 1 so need to convert.
	trimmedInput := strings.TrimSpace(input)
	//ratingFloat, err := strconv.ParseFloat(input, 64)// gets error since input will have \n as well. So, need to trim that.
	ratingFloat, err := strconv.ParseFloat(trimmedInput, 64)

	if err != nil {
		fmt.Println("Error while converting input = ", err)
	} else {
		fmt.Println("Tampered rating = ", ratingFloat+1)
	}

}
