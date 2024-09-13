package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	//mapping and endpoint "/" to the http server with a print stmt as return val
	//anything other than /goodbye will go here -  kind of default one
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
		log.Println("Hello World!") // ----> curl -v localhost:9090
		d, err := io.ReadAll(r.Body)

		if err != nil{
			http.Error(rw, "Oops" , http.StatusBadRequest)
		}
		//log.Printf("Data %s\n", d) ---> To print it as part of server log output

		fmt.Fprintf(rw, "Hello %s", d) // ----> To print it as a part of the user terminal i.e. from where we're executing curl commands ----- curl -v -d "TP" localhost:9090 
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("BBye!") // ----> curl -v localhost:9090/goodbye
	})
	http.ListenAndServe(":9090", nil) 
}