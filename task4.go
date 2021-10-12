package main

import (
	"fmt"
)

func main(){
	var i, j, max int
	arr:= []int{17,18,5,4,6,1}
	fmt.Printf("Input array - %d\n", arr)

	for i < len(arr){
		j = i + 1

		if len(arr)-1 != i{
			for j < len(arr){
				if max < arr[j]{
					max = arr[j]
				}
				j++
			}
		} else{
			max = -1
		}
		
		arr[i] = max
		i++
		max = 0
	}

	fmt.Printf("Output array - %d", arr)
}