package main

/*
#include <stdio.h>

int hello(void)
{
	printf("hello");
	return 123;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.hello())
}
