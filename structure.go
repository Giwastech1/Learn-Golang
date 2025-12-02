package main
import "fmt"

type student struct{
	name string
	age int
	school string
}

func updateInfo(pointer *student, changedName string, updateAge int){
	pointer.name = changedName
	pointer.age = updateAge
}


func studentInfo(){
	toyyib := student{}
	toyyib.name = "Giwa Toyyib"
	toyyib.age = 16
	toyyib.school = "Unilorin"
	fmt.Println(toyyib)

	updateInfo(&toyyib, "Giwa Omotosho",25)
	fmt.Println(toyyib)
}

func main(){
	studentInfo()
}