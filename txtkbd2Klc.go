// txtkbd2Klc project
// Copyright 2016 Philippe Quesnel
// Licensed under the Academic Free License version 3.0
// see README !!
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func getVickey(ch byte) string {
	switch ch {
	case ';', ':':
		return "OEM_1"
	case '+', '=':
		return "OEM_PLUS"
	case ',', '<':
		return "OEM_COMMA"
	case '-', '_':
		return "OEM_MINUS"
	case '.', '>':
		return "OEM_PERIOD"
	case '/', '?':
		return "OEM_2"
	case '`', '~':
		return "OEM_3"
	case '[', '{':
		return "OEM_4"
	case '\\', '|':
		return "OEM_5"
	case ']', '}':
		return "OEM_6"
	case '\'', '"':
		return "OEM_7"
	default:
		return "?"
	}
}
func main() {
	// read text file keyboard def (30 mains keys)
	if len(os.Args) != 2 {
		fmt.Println("parameters: inputKeyboardDefFilename")
		return
	}
	strBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	// remove all whitespace, end up with only the keys
	var re = regexp.MustCompile(`\s+`)
	strBytes = re.ReplaceAllLiteral(strBytes, nil)

	if len(strBytes) != 60 {
		fmt.Printf("expecting 60 characters, got %d\n", len(strBytes))
		os.Exit(-1)
	}

	// print back keyboard, shifted/non-shifted together
	for row := 0; row < 3; row += 1 {
		rowUp := strBytes[row*10:]
		rowLo := strBytes[row*10+30:]
		fmt.Print("// ")
		for charIdx := 0; charIdx < 10; charIdx += 1 {
			charUp := rowUp[charIdx]
			charLo := rowLo[charIdx]
			fmt.Printf(" %c%c", charLo, charUp)
		}
		fmt.Println("")
	}

	// output KLC entries
	usedVickeys := make(map[string]bool)
	oemVkIdx := 1
	for row := 0; row < 3; row += 1 {
		fmt.Println("")
		rowUp := strBytes[row*10:]
		rowLo := strBytes[row*10+30:]
		scanCode := 16 + row*14

		// all keys of this row
		for charIdx := 0; charIdx < 10; charIdx += 1 {
			charUp := rowUp[charIdx]
			charLo := rowLo[charIdx]

			// different output for symbols vs letters
			if strings.Contains("QAZWSXEDCRFVTGBYHNUJMIKOLP", string(charUp)) {
				fmt.Printf("%x\t%c\t\t%d\t%c\t%c",
					scanCode+charIdx, charUp, 1, charLo, charUp)
			} else {
				// symbols need a VK_ and output as unicode value (hex: 003e)
				vickey := getVickey(charLo)

				if usedVickeys[vickey] {
					fmt.Println("// this VK is already used ! manually fix it")
				}
				usedVickeys[vickey] = true

				fmt.Printf("%x\t%s\t%d\t%04x\t%04x",
					scanCode+charIdx, vickey, 0, charLo, charUp)
				oemVkIdx += 1
			}
			// end line for this key with comment showing the characters
			fmt.Printf("\t\t// %c %c\n", charLo, charUp)
		}
	}

	// output unused VKs as ref to fix dup VKs
	if len(usedVickeys) > 0 {
		fmt.Println("\nunused VKs (main 30 keys)")
		vksStr := ",./;"
		for _, ch := range vksStr {
			vickey := getVickey(byte(ch))
			if !usedVickeys[vickey] {
				fmt.Printf(" %c : %s\n", ch, vickey)
			}
		}

		fmt.Println("\nunused VKs (whole kbd)")
		vksStr = "+,-.;/`[\\]'"
		for _, ch := range vksStr {
			vickey := getVickey(byte(ch))
			if !usedVickeys[vickey] {
				fmt.Printf(" %c : %s\n", ch, vickey)
			}
		}
	}
}
