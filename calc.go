package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var Res int
	var Tes bool
	Intnum := 0
	Intten := 0
	Inthon := 0
	Inttho := 0
	Rum := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	Rumrev := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"}
	Rumrevten := map[int]string{10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC"}
	Rumrevhon := map[int]string{100: "C", 200: "CC", 300: "CCC", 400: "CD", 500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM"}
	Rumrevtho := map[int]string{1000: "M", 2000: "MM", 3000: "MMM"}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	line := strings.Split(input, " ")
	for index, value := range line {
		if value == "+" {
			Res = sum(co(line[index-1]), co(line[index+1]))
		} else if value == "*" {
			Res = mn(co(line[index-1]), co(line[index+1]))
		} else if value == "/" {
			Res = del(co(line[index-1]), co(line[index+1]))
		} else if value == "-" {
			Res = mi(co(line[index-1]), co(line[index+1]))
		}
	}
	if Tes != false && Res > 0 {
		fmt.Println(coret(res))

	}
}
func co(a string) (b int) {

	if _, ok := Rum[a]; ok {
		tes = true
		return Rum[a]
	}
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
func coret(a int) string {
	if a > 4000 {
		fmt.Println("Too much")
	} else if a/1000 > 0 {
		Inttho = a / 1000
	} else if (a/1000)/100 > 0 {
		Inthon = (a / 1000) / 100
	} else if ((a/1000)/100)/10 > 0 {
		Intten = ((a / 1000) / 100) / 10
	} else {
		Intnum = ((a / 1000) / 100) % 10
	}

}
func sum(a int, b int) int {
	Res := a + b
	return Res
}

func mn(a int, b int) int {
	Res := a * b
	return Res
}

func del(a int, b int) int {
	Res := a / b
	return Res
}
func mi(a int, b int) int {
	Res := a - b
	return Res
}
