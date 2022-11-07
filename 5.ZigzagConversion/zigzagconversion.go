package main

import "fmt"

func convert(s string, numRows int) string {
	mid := 0
	var rez string
	if numRows > 2 {
		mid = numRows - 2
	}
	for i:= 0; i<numRows; i++ {
		for j:=i; j < len(s); j+=numRows+mid {
			rez = rez + s[j:j+1]
			if i > 0 && i < numRows-1 {
				imid := (numRows - i - 1) * 2
				if j+imid >= len(s) {
					continue
				}
				rez = rez + s[j+imid:j+imid+1]	
			} 
		}
	}
	return rez
}

func main() {
	
	s := "PAYPALISHIRING"
	numRows := 4

	// s = "abcdef"

	fmt.Println(convert(s, numRows))
}