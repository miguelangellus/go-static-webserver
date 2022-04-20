package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("\n1.server is starting...\n2.follow: http://localhost:8080\n3.hint: it only can be accessed by binding a port, so run..\ndocker run --rm -d -p 8089:8089 go-static-site")

	log.Fatal(http.ListenAndServe(":8080", nil))

	//optional
	//err := http.ListenAndServe(":8089", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
