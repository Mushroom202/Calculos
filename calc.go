package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var Res int
	var ResRum string
	var typerum bool
	Rum := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	Rumrev := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"}
	Rumrevten := map[int]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC", 10: "C"}
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	line := strings.Split(input, " ")
	if len(line) > 3 {
		log.Fatal("Ошибка: Больше 2 чисел или больше 1 математического действия")
	}
	for _, value := range line {
		if co(line[0]) > 10 && co(line[0]) < 0 && co(line[2]) > 10 && co(line[2]) < 0 {
			log.Fatal("Ошибка: Число меньше или больше 10")
		}
		if value == "+" {
			Res = sum(co(line[0]), co(line[2]))
		} else if value == "*" {
			Res = mn(co(line[0]), co(line[2]))
		} else if value == "/" {
			Res = del(co(line[0]), co(line[2]))
		} else if value == "-" {
			Res = mi(co(line[0]), co(line[2]))
		}
		_, ok := Rum[line[0]]
		_, ok1 := Rum[line[2]]
		if (ok1 && ok == false) && (ok1 || ok != true) {
			log.Fatal("Ошибка, только римские")
		} else if ok1 && ok == true {
			typerum = true

		} else {
			fmt.Println("int", Res)
		}
	}
	if typerum == true {
		if Res <= 0 {
			log.Fatal("Ошибка: для римских цифр нельзя использовать цифры меньше или равно 0")
		}
		if Res/10 > 0 {
			ResRum += Rumrevten[Res/10]
			ResRum += Rumrev[Res%10]
		} else {
			ResRum += Rumrev[Res%10]
		}
		fmt.Println(ResRum)
	}

}
func co(a string) (b int) {
	Rum := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	if _, ok := Rum[a]; ok {
		return Rum[a]
	}
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
	}
	return b
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
