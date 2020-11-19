package main 

import(
	"fmt"
	"io/ioutil"
)
//Page is a struct which contains the meta info of the page 
type Page struct{
	Title string
	Body []byte
}

//this is a byte as the ioutil package takes that as input 

func(p *Page) save() error{
	filename:=p.Title+".txt"
	return ioutil.WriteFile(filename,p.Body,0600)	 
}

func loadpage(title string) (*Page,error){
	filename:=title+".txt"
	body,err:=ioutil.ReadFile(filename)
	if err!=nil{
		return nil,err
	}
	return &Page{
		Title: title,
		Body: body,
	},nil

}

func main(){
	page1:=Page{
		Title:"Testingpage",
		Body: []byte("This is a test"),
	}
	page1.save();
	page2,_:=loadpage(page1.Title)
	fmt.Println(string(page2.Body))
}