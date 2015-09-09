package main

import (
	"fmt"
	"strconv"
)

type row [10]int

// multiplication table
type mTable [10]row

var d mTable = mTable{
	row{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	row{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
	row{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
	row{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
	row{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
	row{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
	row{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
	row{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
	row{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
	row{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
}

// permutation table
type pTable [8]row

var p pTable = pTable{
	row{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	row{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
	row{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
	row{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
	row{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
	row{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
	row{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
	row{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
}

var inv row = row{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}

func generateVerhoeff(num string) string {
	ln := len(num)
	c := 0
	for i := 0; i < ln; i++ {
		c = d[c][p[((i + 1) % 8)][int(num[ln-i-1]-'0')]]
	}
	return strconv.Itoa(inv[c])
}

func validateVerhoeff(num string) bool {
	ln := len(num)
	c := 0
	for i := 0; i < ln; i++ {
		c = d[c][p[(i % 8)][int(num[ln-i-1]-'0')]]
	}
	return c == 0
}

func main() {
	num := "999999"
	check := generateVerhoeff(num)

	checked := fmt.Sprintf("%s%s", num, check)
	ok := validateVerhoeff(checked)
	fmt.Printf("generateVerhoeff(\"%s\") checksum is %v and validateVerhoeff(\"%s\") is %v \n", num, check, checked, ok)
}
