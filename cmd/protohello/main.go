package main

import (
	"fmt"

	"github.com/gangleri/protohello/stubs"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &stubs.Person{
		Name:  "John Doe",
		Email: "j.doe@gmail.com",
	}

	o, _ := proto.Marshal(p)

	var p2 stubs.Person
	proto.Unmarshal(o, &p2)

	fmt.Println(p2.Name)

}
