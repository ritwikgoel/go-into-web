package main

import(
	"fmt"
	"net/http"
	"io"
)

func dogpic(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w,`<img src="assets/hello.png">`)

}

func main(){
	fmt.Print("Hiiii")
	http.Handle("/assets/",http.StripPrefix("/assets",http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/doggie",dogpic)
	http.ListenAndServe(":8081",nil)
	
}