package main

import (
	"fmt"
	"strings"
)

func main() {
	
	words := "elephant"

 	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	token := []string{}

	// lookup for entries made by the user.
	placeholder := make(map[string]bool)
	
	length := len(words)
	chances := length

	//Append "_" in token array
	//In placeholder of all keys initialize with 'false' value
	for i := 0; i < length; i++ {
		token = append(token, "_")
		placeholder[string(words[i])] = false
	}

	//when user enter the choice but its choice is wrong then add guess array
	guess := []string{}

	for {

		Tokenstrs := strings.Join(token, "")

		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		if chances == 0 && words != Tokenstrs {
			fmt.Println("You are loss the game ! try agin")
			break
		}

		//win the game
		if words == Tokenstrs {
			fmt.Println("Given string is = ", Tokenstrs)
			fmt.Println("Win the game !")
			break
		}

		
		fmt.Println("\n", token)
		fmt.Println("Chances left :", chances)
		fmt.Println("Guesses : ", guess)
		fmt.Println("Guess the letter or word :")
		
		//Input from the strings
		Inputstrings := ""
		fmt.Scanf("%s", &Inputstrings)

		//Inputstring compare with words
		if len(Inputstrings) == length && Inputstrings == words {
			fmt.Println("Win the game...!")
			return
		}
		

		_, ok := placeholder[string(Inputstrings)]
		
		//check key is present and that key value is false
		if ok && placeholder[Inputstrings] == false { 
			//replace the value of that key is ture
			placeholder[Inputstrings] = true
			
						
			for i:= 0; i < length; i++ {
				// compare the words of each character in placeholder 
				if placeholder[string(words[i])] == true && string(token[i]) == "_" {
					//Inputstring is match
					//replace the "_" character
					token[i] = string(words[i])
				}
			}
		}else {
			//Inputstring is not match 
			fmt.Println("Miss chances.....!\n")
			//Inputstring append in guess array
			guess = append(guess, Inputstrings)
			chances--
		}
	}
}

