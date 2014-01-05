package alphakey

import (
	"strings"
)

const UnambiguousLowercaseAlphabet = "abcdefghijkmnpqrstuv"
const FullAlphanumericAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const defaultOffset = 0

type KeyConverter struct {
	Alphabet string
	Offset   int
}

var DefaultConverter = &KeyConverter{
	UnambiguousLowercaseAlphabet,
	defaultOffset,
}

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

func (k *KeyConverter) GetInt(key string) int {
	i := 0
	base := len(k.Alphabet)
	for _, char := range strings.Split(key, "") {
		indexOfChar := strings.Index(k.Alphabet, char)
		i = i*base + indexOfChar
	}
	return i - k.Offset
}

func GetKeyForInt(integer int) string {
	return DefaultConverter.GetKey(integer)
}

func GetIntForKey(key string) int {
	return DefaultConverter.GetInt(key)
}
