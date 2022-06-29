package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	const PI float64 = 3.14
	fmt.Println("Enter a radius to calculate the circumference")
	reader := bufio.NewReader(bufio.NewReader(os.Stdin))
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	input = strings.TrimSpace(input)
	radius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalln(err)
	}
	circumference := PI * radius
	fmt.Println("The circumference is: ", circumference)

}