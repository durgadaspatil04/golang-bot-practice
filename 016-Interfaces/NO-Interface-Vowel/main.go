package main

import "fmt"

// type vowelFinder interface {
// 	findVowel() []rune
// }

/*To define a method on a type, the definition of the receiver type and the definition
of the method should be present in the same package. */
type newString string

func (s newString) findVowel() []rune {
	var vowels []rune
	for _, rune := range s {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' ||
			rune == 'A' || rune == 'E' || rune == 'I' || rune == 'O' || rune == 'U' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	str := newString("Hello Eve of Crismus")
	//var a vowelFinder
	//a = str // possible since newString implements VowelsFinder
	//fmt.Printf("Vowels are %c\n", a.findVowel())
	fmt.Printf("Vowels are %c\n", str.findVowel())

}
