package config

import (
	"fmt"
)

type ToDo struct {
	ID    int    `json:"id"`
	TASK  string `json:"task"`
	STATE string `json:"state"`
}

func Result(str string) {
	fmt.Println(str + " Success!")
}
func Failed(str string) {
	fmt.Println(str + " Failed!")
}
