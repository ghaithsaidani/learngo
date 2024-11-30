// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	argsLength := len(os.Args) - 1
	name1, name2, name3, name4, name5 := os.Args[1],os.Args[2],os.Args[3],os.Args[4],os.Args[5]

	fmt.Printf("There are %d!\n", argsLength)
	fmt.Println("Hello great", name1,"!")
	fmt.Println("Hello great", name2,"!")
	fmt.Println("Hello great", name3,"!")
	fmt.Println("Hello great", name4,"!")
	fmt.Println("Hello great", name5,"!")
	fmt.Println("Nice to meet you all.")
}
