package main

import "fmt"


func TreatError(err error){
	if err != nil{
		fmt.Println("Error:", err)
	}
}

func Map(slice []string, f func(string) []string) [][]string {
	var result [][]string
	// result := make([][]string, len(slice))
	for _, v := range slice {
		result = append(result,f(v))
	}
	return result
}