package model

import "fmt"

func model() {
	fmt.Println("Models in go")
}

type data struct {
	Id      int64 `json:"id" bson:"id"`
	Name    string
	Watched string `json:"watched"`
}
