package main

import (
	"fmt"
	"net/http"
)

func main() {

	version := "3.0"

	color := "#e033ff" //Blue 44B3C2 and Yellow F1A94E

	http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<div class='pod' style='background:%s'> ver: %s\n </div>", color, version)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}
