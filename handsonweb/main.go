package main

import(
	"fmt"
	"log"
	"net/http"
	"html/template"
)


func index(w http.ResponseWriter, r *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err =tpl.ExecuteTemplate(w,"index.html","Ritwik")
	if(err!=nil){
		log.Fatal(err)
	}
}

func main(){
	http.HandleFunc("/", index)
	fmt.Println("This is a hands on example")
	log.Fatal(http.ListenAndServe(":8080", nil))

}