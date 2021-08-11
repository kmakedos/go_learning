package gui

import "fmt"


func init(){
	fmt.Println("Server starting")
}

func CallDesktop(){
	c := CreateDesktop()
	fmt.Println(c.x)
}