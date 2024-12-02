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
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {

    if len(os.Args) != 3 {
        fmt.Println("Usage: [username] [password]")
        os.Exit(0)
    }
    const (
        username = "jack"
        password = "1888"
    )
    u, p := os.Args[1], os.Args[2]

    if u != username {
        fmt.Printf("Access denied for %q\n",u)
    } else if p != password {
        fmt.Printf("Invalid password for %q\n",u)
    } else {
        fmt.Printf("Acces granted to %q\n",u)
    }

    
}
