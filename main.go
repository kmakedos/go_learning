package main

import "fmt"

const (
	n =  iota
	m
	g
)


func reverse(s *[3]int) {
	for i,j:=0, len(s)-1; i<j; i,j=i+1,j-1 {
		s[i],s[j] = s[j],s[i]
	}
}

func reverse_b(b []rune) string {
	for i,j:=0,len(b)-1;i<j;i,j=i+1,j-1 {
		fmt.Println(b[i],b[j])
		b[i],b[j] = b[j],b[i]
	}
	return string(b)
}
func main(){
	s := [3]int {1,2,3}
	reverse(&s)
	fmt.Println(s)
	b := "κωστας 你 好! "
	b= reverse_b([]rune(b))
	fmt.Println(b)
}