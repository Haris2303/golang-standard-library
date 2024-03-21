package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ucup Surucup", "cup"))
	fmt.Println(strings.Split("Ucup Surucup Tamvan", " "))
	fmt.Println(strings.ToLower("Ucok Surocok"))
	fmt.Println(strings.ToUpper("Ucok Surocok"))
	fmt.Println(strings.Trim("    Ucok Surocok      ", " "))
	fmt.Println(strings.ReplaceAll("Otong Surotong, Otong Tamvan", "Otong", "Dani"))
}
