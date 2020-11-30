package main

import(
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","text/html")
    fmt.Fprint(w,"<h1> this is the welcome screen</h1>")

}
func contact(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","text/html")
    fmt.Fprint(w,"this is the contact page")
}


func main(){
    rr:=mux.NewRouter()
    rr.HandleFunc("/",home)
    rr.HandleFunc("/contact",contact)
    http.ListenAndServe(":8080",rr)
}