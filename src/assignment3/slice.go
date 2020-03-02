package main

import (
	"fmt"
	"strconv"
	"sort"
	
)

func main() {
	
	sli := make([]int, 0, 3)

	for {
		var value string
		_, err := fmt.Scan(&value)
		if(err == nil) {
			if(value == "X") {
				break
			}
				
			val, err := strconv.Atoi(value)
			if (err== nil) {
				sli = append(sli, val)
			}
		}		
	}
	
	sort.Ints(sli)
	
	fmt.Println(sli)
}