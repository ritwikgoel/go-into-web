package main

import(
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

//tpl is a container of tempaltes

type person struct{
	Name string
	age int 
}


func init(){
	tpl=template.Must(template.ParseGlob("temps/*"))
}

func main(){
	//tpl,_:=template.ParseGlob("temps/*")
	// er:=tpl.Execute(os.Stdout,nil)
	// if er!=nil{
	// 	log.Fatalln(er)
	// }

	p1:=person{
		Name:"Ritwik",
		age: 89,
	}
	q:=[]person{
		p1,
	}

	err:=tpl.ExecuteTemplate(os.Stdout,"test1.html",q)
	if err!=nil{
		log.Fatalln(err)
	}


}