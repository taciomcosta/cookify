package main

import (
	"fmt"
	"net/http"

	"github.com/taciomcosta/cookify/pkg/giphy"
)

func main() {
	fmt.Println("Hello, world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := []byte("Hello, world")

		//recipes, err := recipepuppy.FindRecipes("onions, tomato", "", 3)
		//if err != nil {
		//fmt.Println(err)
		//}
		//fmt.Println(recipes)

		gifs, err := giphy.Search("cheeseburger")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(gifs)

		w.Write(res)
	})

	http.ListenAndServe(":3000", nil)
}
