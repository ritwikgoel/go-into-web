package main

import(
	"net/http"
	"io"
	"fmt"
)


func pict(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html")
	io.WriteString(w,`<img src=/assets/hello.png>`)
	fmt.Print("hi")
}
 

func main(){
	http.HandleFunc("/",pict)
	http.Handle("/assets/",http.StripPrefix("/assets",http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080",nil)
}