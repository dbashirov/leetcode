package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var cntT, cntN int

	in := bufio.NewReader(os.Stdin)
	cntT = readInt(in)

	// Обход по наборам данных
	for t := 0; t < cntT; t++ {
		// Обход записей
		cntN = readInt(in)
		ab := make(map[string][]string)
		for n := 0; n < cntN; n++ {
			name, tel := readLog(in)
			//fmt.Printf("name=%s tel=%s", name, tel)
			if arrTel, ok := ab[name]; ok {
				index := contained(&arrTel, tel)
				if index != -1 {
					// continue
					arrTel = removeByIndex(arrTel, index)
					arrTel = append(arrTel, tel)
				} else {
					if len(arrTel) >= 5 {
						arrTel = arrTel[1:]
					}
					//fmt.Printf("1: %v len=%d\n", arrTel, len(arrTel))
					arrTel = append(arrTel, tel)
					ab[name] = arrTel
				}
			} else {
				//fmt.Printf("2: %v\n", arrTel)
				arrTel = append(arrTel, tel)
				ab[name] = arrTel
				//fmt.Printf("name=%s tel=%s arrTel=%s\n", name, tel, arrTel)
			}

		}

		// Сортировка
		keys := make([]string, 0, len(ab))
		for k := range ab {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// Вывод результата

		for _, k := range keys {
			fmt.Printf("%s: %d ", k, len(ab[k]))
			for i := len(ab[k]) - 1; i >= 0; i-- {
				fmt.Printf("%s ", ab[k][i])
			}
			fmt.Printf("\n")
		}
		fmt.Println()
	}
}

func contained(arr *[]string, s string) int {
	for k, v := range *arr {
		if s == v {
			return k
		}
	}
	return -1
}

func removeByIndex(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
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

func readLog(in *bufio.Reader) (string, string) {
	numbs := strings.Split(readLine(in), " ")
	return numbs[0], numbs[1]
}
