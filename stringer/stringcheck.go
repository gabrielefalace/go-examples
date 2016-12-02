package stringer

import "strings"

func CheckDirect(tofind, text string) bool {
	if len(tofind) != len(text) {
		return false
	}
	commonLength := len(text)

	for occurrenceIndex, occurrenceChar := range text {

		if occurrenceChar == rune(tofind[0]) {

			toFindIndex := 0
			textIndex := occurrenceIndex
			matchedChars := 0

			for j := 0; j < commonLength; j++ {

				// compare & match - character by character
				if tofind[toFindIndex] == text[textIndex] {
					matchedChars++
				} else {
					break
				}

				// increment on first string
				toFindIndex++

				// increment on second string (circular)
				textIndex = textIndex + 1
				if textIndex == commonLength {
					textIndex = 0
				}
			}

			// len(text) consecutive chars matched means string fully match
			if matchedChars == commonLength {
				return true
			}
		}
	}
	return false
}

func CheckConcat(toFind, text string) bool {
	return strings.Contains(text+text, toFind)
}
