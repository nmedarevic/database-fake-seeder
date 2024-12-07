package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {

	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.LastName())
	fmt.Println(gofakeit.FirstName())
	fmt.Println(gofakeit.PastDate().UTC())
	fmt.Println(gofakeit.FutureDate())
	fmt.Println(gofakeit.BuzzWord())
	fmt.Println(gofakeit.Company())
	fmt.Println(gofakeit.HackerPhrase())
	fmt.Println(gofakeit.HackerNoun())
	fmt.Println(gofakeit.BS())
	fmt.Println(gofakeit.BeerStyle())
}