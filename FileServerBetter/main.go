package main

import(
	"fmt"
	"net/http"
	"io"
)

func dogpic(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w,`<img src="hello.png">`)

}

func main(){
	fmt.Print("Hiiii")
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.HandleFunc("/doggie",dogpic)
	http.ListenAndServe(":8080",nil)
	
}