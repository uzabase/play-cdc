package main

import (
	"fmt"
	"specify/domain"
	"specify/repository"
)

func main() {
	contract := repository.FindContract("example.json")
	spec := domain.GenerateSpec(contract)
	fmt.Println(spec)
}
