/*
A library for generating short shareable alphanumeric codes from integers for url shorteners and the like

The readme can be found at https://github.com/joho/alphakey

Usage

To get going, just import this

		import "github.com/joho/alphakey"
You can either call a couple of the helper functions directly

		key := alphakey.GetKeyForInt(198898042) // returns "ninja"
		num := alphakey.GetIntForKey("dgt")     // returns 1337

Or if you want a bit more control you can instantiate your own converter and control the alphabet used and any offset you might want

		converter := &alphakey.Converter{
		  alphakey.UnambiguousLowercaseAlphabet, // all lower case, no l or o, because humans
		  497,                                   // an arbitrary offset so early sharers don't get a or b
		}

		key := converter.GetKey(1988483) // returns "ninja"
		num := converter.GetInt("emq")   // returns 1337
*/
package alphakey

import (
	"strings"
)

// My go-to alphabet for generating shareable codes: all lower case, no l or o (which can be mixed up with 1 and 0)
const UnambiguousLowercaseAlphabet = "abcdefghijkmnpqrstuv"

// The full alphanumeric set of chars to use for a code (which is used by default)
const FullAlphanumericAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// By default we start counting from zero
const defaultOffset = 0

// The type that does all the work. You only want your own instance of this if you're using a custom
// alphabet and/or offset
type KeyConverter struct {
	Alphabet string
	Offset   int
}

// Generates a key based on the alphabet & offset set in your converter
func (k *KeyConverter) GetKey(num int) string {
	num = num + k.Offset

	base := len(k.Alphabet)

	if num == 0 {
		return k.Alphabet[:1]
	}

	s := []string{}
	for num > 0 {
		index := num % base
		s = append(s, k.Alphabet[index:index+1])
		num = num / base
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, "")
}

// can turn a key generated by the same converter back into an int
func (k *KeyConverter) GetInt(key string) int {
	i := 0
	base := len(k.Alphabet)
	for _, char := range strings.Split(key, "") {
		indexOfChar := strings.Index(k.Alphabet, char)
		i = i*base + indexOfChar
	}
	return i - k.Offset
}

// DefaultConverter - aptly named I think
var DefaultConverter = &KeyConverter{
	UnambiguousLowercaseAlphabet,
	defaultOffset,
}

// return a "key" for a given int using the default alphabet
func GetKeyForInt(integer int) string {
	return DefaultConverter.GetKey(integer)
}

// take a "key" generated by the default converter and turn it back into an int
func GetIntForKey(key string) int {
	return DefaultConverter.GetInt(key)
}
