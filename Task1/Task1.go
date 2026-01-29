package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func intro(p *Person) string{
    msg := ""
    fmt.Printf("Hello!, I am %s, my age is %d",p.Name,p.Age)
	if p.Age < 18 {
       msg +=  fmt.Sprintf(" and I am not eligible to Vote")
	}else{
	   msg += fmt.Sprintf(" and I am eligible to Vote")
	}
	return msg
}

func main(){
	var p1 Person
	fmt.Print("Enter Name: ")
	if _,err := fmt.Scan(&p1.Name); err != nil {
		fmt.Println("Error reading name",err)
		return
	}

	fmt.Print("Enter Age: ")
	if _,err := fmt.Scan(&p1.Age); err != nil {
		fmt.Println("Error reading age",err)
		return
	}

	fmt.Println(intro(&p1))
}