package main

import (
	"fmt"
	"net/http"
)

func main() {

	//First approach to create a http handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World! Hey Mahin!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes written: %d\n", n)
	})

	_ = http.ListenAndServe(":8000", nil)

}
