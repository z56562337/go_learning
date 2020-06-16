package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!!????!?!!?!?!?!?!?!?!?!")
	loop()
}

func loop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
