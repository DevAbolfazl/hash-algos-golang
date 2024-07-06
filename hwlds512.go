/*
	HWLDS (Hash with length-dependent security)
	512 Bits
*/

package hash_algos

import (
	"encoding/hex"
	"errors"
)

func hwlds512(inputString string) string {
	if inputString == "" {
		_ = errors.New("HWLDS Hash Error: Empty input")
	}
	var result []byte = []byte(inputString)
	var helpers = []byte{0b00000000}
	for i := 0; i < 128-len(result); i++ {
		result = append(result, helpers[0])
	}
	var x int = 0
	for y := range result {
		if x == len(helpers) {
			x = 0
		}
		result[y] ^= helpers[x]&result[y] | (result[y] << 2)
		for z := len(result) - 1; z >= 0; z-- {
			helpers = append(helpers, result[y]^result[z]<<helpers[x])
			for h := range helpers {
				helpers[h] = helpers[h] ^ helpers[x+1]
			}
			for h := range helpers {
				result[x] = result[y] ^ helpers[h]
			}
			result[z] = result[z] ^ (helpers[x] | result[y])
			result[y] ^= result[z]
		}
		x++
	}
	return hex.EncodeToString(result[len(result)-65 : len(result)-1])
}
