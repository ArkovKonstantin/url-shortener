package shortener

import (
	"fmt"
	"math"
	"strconv"
)

func convertBase(val string, base, toBase int) (string, error) {
	alphabet := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-._~") // len = 66
	//fmt.Println(len(alphabet))
	idxs := make(map[rune]int)
	for i, v := range alphabet {
		idxs[v] = i
	}
	var dec int
	if base != 10 {
		runes := []rune(val)
		for i, j := len(runes)-1, 0; i >= 0; i, j = i-1, j+1 {
			char := runes[i]
			v := idxs[char]
			p := int(math.Pow(float64(base), float64(j)))
			dec += v * p
		}
	} else {
		dec, _ = strconv.Atoi(val)
	}
	var res []rune
	if toBase != 10 {
		var m int
		mods := make([]int, 0)
		for ; dec >= toBase; {
			dec, m = dec/toBase, dec%toBase
			mods = append(mods, m)
		}
		mods = append(mods, dec)
		for i := len(mods) - 1; i >= 0; i-- {
			x := mods[i]
			r := alphabet[x]
			res = append(res, r)
		}
		return string(res), nil
	}else{
		return strconv.Itoa(dec), nil
	}
}

func main() {
	s, _ := convertBase("10000023423", 10, 66)
	fmt.Println(s)
}
