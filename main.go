package main

import "fmt"

var m = make(map[string]rune)
var morseCodes = [26]string{".-", "-...", "-.-.", "-..",
	".", "..-.", "--.", "....",
	"..", ".---", "-.-", ".-..",
	"--", "-.", "---", ".--.",
	"--.-", ".-.", "...", "-",
	"..-", "...-", ".--", "-..-",
	"-.--", "--.."}

func main() {
	initKey()

	smalpha(".--...-.-.-.....-.--........----.-.-..---.---.--.--.-.-....-..-...-.---..--.----..")
	// smalpha(".----...---.-....--.-........-----....--.-..-.-..--.--...--..-.---.--..-.-...--..-")
	// smalpha("..-...-..-....--.---.---.---..-..--....-.....-..-.--.-.-.--.-..--.--..--.----..-..")
}

func initKey() {
	var letter rune = 'a'

	for _, code := range morseCodes {
		m[code] := letter
		letter++
	}

	fmt.Println("map:", m)
}
func smalpha(code string) {

	fmt.Println(code)
}
