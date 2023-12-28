package main

import "fmt"

type Movable interface {
	Move()
}

type Object struct {
	Id   string // public
	name string // private
}

func (obj *Object) Move() {
	fmt.Println("Object moving...")
}

func MoverFunc(obj Movable) {
	fmt.Printf("Execute Move ...")
	obj.Move()
}

func main() {
	object := &Object{Id: "123456", name: "NPC"}
	MoverFunc(object)
	(*object).Move()

	comp := Composite{Object: *object}
	comp.Move()
	fmt.Println(comp.Id) // instead of comp.Object.Id
	fmt.Println(comp.Object)
}

type Composite struct {
	Object
}
