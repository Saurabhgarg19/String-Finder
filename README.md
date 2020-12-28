StringFinder
==================

This is a KMP(Knuth–Morris–Pratt algorithm) implement and related string function `Strstr` and `Strchr` which using `KMP` feature inside.

In computer science, the Knuth–Morris–Pratt string searching algorithm (or KMP algorithm) searches for occurrences of a "word" W within a main "text string" S by employing the observation that when a mismatch occurs, the word itself embodies sufficient information to determine where the next match could begin, thus bypassing re-examination of previously matched characters.
 
Install
---------------
`go get github.com/kkdai/kmp`


Usage
---------------

```go
package main

import (
	"fmt"
)

func main() {
    
	  fmt.Println("Enter The Main String: ") 
    var first string 
    fmt.Scanln(&first)
    
    fmt.Println("Enter SubString that need to find: ") 
    var second string 
    fmt.Scanln(&second) 
  
    fmt.Println("The Index location of the Substring in Main String: ") 
	
	  //Using KMP
	  fmt.Println(KMP(second, first))
}	

Project
---------------

It is one of my GoLang Project
