package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/gangleri/protohello/stubs"
)

func main() {
	p := &stubs.Person{
		Name: "John Doe",
	}

	o, _ := proto.Marshal(p)

	var p2 stubs.Person
	proto.Unmarshal(o, &p2)

	fmt.Println(p2.Name)

}
