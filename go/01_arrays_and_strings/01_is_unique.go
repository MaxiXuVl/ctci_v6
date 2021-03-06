package arrays

// str.size = n
// alphabet(str) = a (size of alphabet)

// A: time = O(n), mem = O(a)
func AllUniqA(str string) bool {
	h := make(map[rune]bool)

	for _, c := range str {
		if h[c] {
			return false
		} else {
			h[c] = true
		}
	}

	return true
}

// C: time = O(n*log(n)), mem = O(1)
func AllUniqC(str string) bool {
	str = SortStringByCharacter(str)
	runes := StringToRuneSlice(str)

	for i := 0; i < len(runes)-1; i++ {
		if str[i] == str[i+1] {
			return false
		}
	}

	return true
}

// E: time = O(n), mem = O(1)
// Special case if alphabet <= 64 (for 64 bits systems), and alphabet is continuous
func AllUniqE(str string) bool {
	bits := 0

	for _, c := range str {
		bit := 1 << uint(c-'a')

		if bits&bit == 1 {
			return false
		} else {
			bits |= bit
		}
	}

	return true
}
