package main

import "time"

func main() {
	time.Sleep(5 * time.Minute)

	test := [][]string{}
	for i := 0; i < 100000; i++ {
		test = append(test, make([]string, 10000000))
	}
}
