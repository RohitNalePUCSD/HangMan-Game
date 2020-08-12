package main

import (
	"fmt"
	"strings"
)

func main() {
	
	words := "elephants"
	var char string

	lens := len(words)

	arr := []string{}
	for i := 0; i < lens; i++ {
		arr = append(arr, "_")
	}

	chances := lens
	placeholder := make(map[string]bool)

	for i := 0; i < lens; i++ {
		placeholder[string(words[i])] = false
	} 

	for {

		arrs := strings.Join(arr, "")
		if chances == 0 && words != arrs {
			fmt.Println("You are loss the game ! try agin")
			break
		}

		if words == arrs {
			fmt.Println("Given string is = ", arrs)
			fmt.Println("Win the game !")
			break
		}

		fmt.Println("\nGiven string are = ", arr)
		fmt.Println("Your left chances is = ", chances)
		fmt.Println("Enter the your selected char")
		fmt.Scanf("%s", &char)

		_, k := placeholder[string(char)]
		
		if k && placeholder[char] == false { 
			placeholder[char] = true
			for i:= 0; i < lens; i++ {
				if placeholder[string(words[i])] == true && string(arr[i]) == "_" {
					arr[i] = string(words[i])
					//chances--
				}
			}
		}else {
			fmt.Println("Miss chances.....!\n")
			chances--
		}
	}
}


