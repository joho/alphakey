package alphakey

import (
	"testing"
)

func verifyKeysConvertBack(t *testing.T, testKeyGenerator func(int) (string, int)) {
	for i := 0; i <= 1000; i++ {
		key, regeneratedInteger := testKeyGenerator(i)

		if i != regeneratedInteger {
			t.Errorf("Generated key '%v' from %v, but it converted back to %v", key, i, regeneratedInteger)
		}
	}
}

func TestPublicFuncs(t *testing.T) {
	verifyKeysConvertBack(t, func(num int) (string, int) {
		key := GetKeyForInt(num)
		return key, GetIntForKey(key)
	})
}

func TestDefaultCoverter(t *testing.T) {
	verifyKeysConvertBack(t, func(num int) (string, int) {
		key := DefaultConverter.GetKey(num)
		regeneratedInteger := DefaultConverter.GetInt(key)

		return key, regeneratedInteger
	})
}

func TestCustomConverter(t *testing.T) {
	converter := &KeyConverter{
		"abcdefghijkl",
		200,
	}
	verifyKeysConvertBack(t, func(num int) (string, int) {
		key := converter.GetKey(num)
		regeneratedInteger := converter.GetInt(key)

		return key, regeneratedInteger
	})
}
