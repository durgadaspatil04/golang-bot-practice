package main

import "fmt"

type vowelFinder interface {
	findVowel() []rune
	findLength() int
}

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
func (s newString) findLength() int {
	var r []rune
	r = []rune(s)
	i := 1
	for _, char := range r {
		fmt.Printf("%c ", char)
		i++
	}
	return i

}

func main() {
	str := newString("Hello Eve of Crismus")
	var a vowelFinder
	a = str // possible since newString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", a.findVowel())
	fmt.Printf("Length is %d\n", a.findLength())
	fmt.Printf("Length is %d\n", str.findLength())
	fmt.Printf("Vowels are %c\n", str.findVowel())

}
