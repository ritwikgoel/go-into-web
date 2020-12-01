package main

import(
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

func init(){
	tpl=template.Must(template.ParseGlob("temps/*"))
}

func main(){
	//tpl,_:=template.ParseGlob("temps/*")
	// er:=tpl.Execute(os.Stdout,nil)
	// if er!=nil{
	// 	log.Fatalln(er)
	// }
	err:=tpl.ExecuteTemplate(os.Stdout,"test1.html",nil)
	if err!=nil{
		log.Fatalln(err)
	}


}