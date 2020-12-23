package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		err := http.ListenAndServe(":8080", nil)

		if err != nil {
			panic(err.Error())
		}

	})
}