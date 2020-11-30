package main

import(
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","text/html")
    if r.URL.Path == "/"{
        fmt.Fprint(w,"<h1> this is the welcome screen</h1>")
    } else if r.URL.Path == "/contact"{
        fmt.Fprint(w,"this is the contact page")
    } else {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w,"404 page with the request")
    }
}

func main(){
    http.HandleFunc("/",handler)
    http.ListenAndServe(":8080",nil)
}