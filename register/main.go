package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var cntT, cntN int

	in := bufio.NewReader(os.Stdin)
	cntT = readInt(in)

	// Обход по наборам данных
	for t := 0; t < cntT; t++ {
		// Обход по попыткам регистрации
		cntN = readInt(in)
		setLogins := make(map[string]struct{}, cntN)
		for n := 0; n < cntN; n++ {
			login := strings.ToUpper(readLine(in))
			fmt.Println(checkLogin(login, &setLogins))
			// fmt.Println(login)
		}
		fmt.Println()
	}
}

func checkLogin(l string, setLogins *map[string]struct{}) string {
	// проверка длины
	if len(l) < 2 || len(l) > 24 {
		return "NO"
	} else {
		// проверка на уникальность
		if _, ok := (*setLogins)[l]; ok {
			return "NO"
		} else {
			(*setLogins)[l] = struct{}{}
			if l[0] == 45 {
				return "NO"
			}
			// return "YES"
			for i := 0; i < len(l); i++ {
				if !(l[i] == 45 || l[i] == 95 || (l[i] > 47 && l[i] < 58) || (l[i] > 64 && l[i] < 91) || (l[i] > 96 && l[i] < 123)) {
					return "NO"
				}
			}
			return "YES"
		}
	}
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLine(in *bufio.Reader) string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	return line
}
