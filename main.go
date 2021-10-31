package main

import "fmt"

type Node struct {
	id int
	kids []*Node
	parent *Node
}

func createNode(id int)*Node{
	t := Node{id, nil, nil}
	return &t
}

func (this *Node)addChild(kid *Node){
	kid.parent = this
	this.kids = append(this.kids, kid)
}

func (this *Node)iterateNodes(){
	for k,v:=range(this.kids){
		fmt.Println(k,v.kids)
	}
}


func main(){
	head := createNode(0)
	t1 := createNode(1)
	t2 := createNode(2)
	t3 := createNode(3)
	t4 := createNode(4)
	t5 := createNode(5 )
	head.addChild(t1)
	head.addChild(t2)
	t1.addChild(t3)
	t2.addChild(t3)
	t3.addChild(t4)
	t3.addChild(t5)

	head.iterateNodes()





}
