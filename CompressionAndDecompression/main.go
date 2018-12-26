//
// Solution to Google Tech Dev Guide Advanced Programming question
// Compression and Decompression
// https://techdevguide.withgoogle.com/paths/advanced/compress-decompression/
//
// Author: Dylan Wright
// Date: December 25, 2018
//

package main

import "fmt"
import "strconv"
import "unicode"

func main() {
	fmt.Println(decompress("3[abc]4[ab]c"))
	fmt.Println(decompress("2[3[a]b]"))
	fmt.Println(decompress("5[2[c]2[a]2[ke]4[ger]re]2[er]f"))
	fmt.Println(decompress("catsa2[re]a2[l]ycuteandilovethem"))
	fmt.Println(decompress("l2[o]katmeimwrit2[ing]olang"))
	fmt.Println(decompress("l100[o]katmeimwrit2[ing]olang"))
}

func decompress(str string) string {
	fmt.Println("Input: " + str)

	decompressedStr := "" // Returned decompressed string
	var runeArr []rune    // Character array

	// Convert String into Rune Array
	// and decompress as you go
	for _, char := range str {

		// When "]" is found begin popping characters until "[" is found
		// Otherwise just append character to runeArr
		if string(char) == "]" {

			tempStr := ""
			// Pop characters and add to a temporary string
			for i := len(runeArr) - 1; string(runeArr[i]) != "["; i-- {
				tempStr = string(runeArr[i]) + tempStr
				runeArr = runeArr[:i]
			}

			// Get number multiplier and clean up "[" and number from array
			runeArr = runeArr[:len(runeArr)-1]
			multiplierStr := ""
			// Get all digits until next rune isn't a number
			for i := len(runeArr) - 1; (len(runeArr) > 0) && unicode.IsNumber(runeArr[i]); i-- {
				multiplierStr = string(runeArr[i]) + multiplierStr
				runeArr = runeArr[:i]
			}
			multiplier, _ := strconv.Atoi(multiplierStr)

			newTempStr := ""
			// Multiply temp string according to multiplier
			for i := 0; i <= multiplier-1; i++ {
				newTempStr = newTempStr + tempStr
			}

			// Append characters back onto runeArr
			for _, char := range newTempStr {
				runeArr = append(runeArr, char)
			}

		} else {
			// Append character to runeArr
			runeArr = append(runeArr, char)
		}
	}

	// Convert Rune Array back into String
	for _, char := range runeArr {
		decompressedStr = decompressedStr + string(char)
	}

	return decompressedStr
}
