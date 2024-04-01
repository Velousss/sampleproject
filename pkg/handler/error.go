package handler

import "fmt"

func HandleError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
