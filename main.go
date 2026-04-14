package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"strings"
)

var words = []string{
	"apple", "banana", "cherry", "dragon", "eagle", "falcon", "guitar", "harbor",
	"island", "jungle", "knight", "lemon", "mango", "noble", "ocean", "panda",
	"queen", "river", "sunset", "tiger", "umbrella", "violet", "walnut", "xenon",
	"yellow", "zebra", "anchor", "breeze", "castle", "diamond", "ember", "forest",
	"galaxy", "horizon", "ivory", "jasper", "karma", "lunar", "marble", "nebula",
	"oracle", "phoenix", "quartz", "ruby", "silver", "thunder", "unity", "velvet",
	"winter", "crystal", "arctic", "blaze", "cosmic", "dusk", "eclipse", "flame",
}

func main() {
	words := flag.Int("w", 4, "Number of words")
	sep := flag.String("s", "-", "Separator")
	num := flag.Bool("n", false, "Add number")
	cap := flag.Bool("c", false, "Capitalize")
	flag.Parse()
	
	phrase := generatePhrase(*words, *sep, *num, *cap)
	fmt.Println(phrase)
}

func generatePhrase(wordCount int, sep string, addNum, capFirst bool) string {
	b := make([]byte, 2)
	rand.Read(b)
	seed := binary.LittleEndian.Uint16(b)
	
	result := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		idx := int(uint16(seed+uint16(i)*0x7f3a)) % len(words)
		w := words[idx]
		if capFirst {
			w = strings.Title(w)
		}
		result[i] = w
	}
	
	phrase := strings.Join(result, sep)
	if addNum {
		n := seed % 1000
		phrase = fmt.Sprintf("%s%d", phrase, n)
	}
	return phrase
}
