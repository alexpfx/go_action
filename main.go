package main

import "fmt"

type test struct {
	str string
	
	slic [] string
	
}

func (t test) testRun()  {
	
	fmt.Println(t.str)
}


func main()  {
	myslice := make([]string, 0)
	myslice = append(myslice, "oi")
	var t test
	
	t.testRun()
	myslice = append(myslice, t.slic...)
	
	fmt.Println(myslice)
	
}