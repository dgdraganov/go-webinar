package main

import "fmt"

type Object struct {
	Id   string // public
	name string // private
}

func (obj *Object) ChangeId() {
	obj.Id = "01234"
}

func (obj Object) ChangeName(name string) {
	obj.name = name
}

func (obj Object) Describe() {
	// ...
}

type Composite struct {
	Object
}

func main() {
	object := &Object{Id: "123456", name: "murloc"}
	object.ChangeName("gnoll") // name will not be changed
	fmt.Println(object)

	comp := Composite{Object: *object}
	comp.Describe() // instead of comp.Object.Describe()

	fmt.Println(comp.Id)        // instead of comp.Object.Id
	fmt.Println(comp.Object.Id) // instead of comp.Object.Id

	fmt.Println(comp.Object) // access nested object
}
