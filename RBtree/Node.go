package RBtree


type Node struct{
	Key int
	Value interface{}
	Left *Node
	Right *Node
	Parent *Node
}



type (node *Node)