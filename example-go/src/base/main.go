package main

import "example-go/_io"

func main()  {
	properties, err := _io.ReaderProperties("./resource", "application")
	if err != nil {
		print(err)
	}
	println(properties)
}




