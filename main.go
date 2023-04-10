package main

import (
	"Lesson11/Lists"
	"fmt"
	"log"
	"strconv"
)

func main() {

	Humo, Others := Lists.Sort()

	/*
		fmt.Println("Список карточек Хумо:")
		fmt.Printf("%v\n\n", Humo)
		fmt.Println("Список карточек Других банков:")
		fmt.Printf("%v\n\n", Others)
	*/

	var input string
	fmt.Println("Введите счёт карты!")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("Неверный формат ввода!", err)
	}
	fmt.Println("Введенное значение:", input)
	if len(input) != 9 {
		log.Fatal("Длинна не соответствует формату счёта карты! (Количество символов счёта = 9)")
	}
	if input[4] != 45 {
		log.Fatal(`Отсутствует знак "-"`)
	}
	numCheck, err := strconv.Atoi(input[0:4])
	if err != nil {
		fmt.Println("Ошибка ввода")
		fmt.Println(numCheck)
	}
	numCheck2, err := strconv.Atoi(input[5:9])
	if err != nil {
		fmt.Println("Ошибка ввода")
		fmt.Println(numCheck2)
	}
	if v, exists := (Humo[input]); exists {
		fmt.Print(v)
		fmt.Println("Такая карта уже существует ")
	} else if v, exists := (Others[input]); exists {
		fmt.Print(v)
		fmt.Println("Такая карта уже существует ")
	} else {
		Lists.CardList1 = append(Lists.CardList1, input)

		if input[3] >= 53 {
			fmt.Printf(" Новая карта Хумо была добавлена\n")
			Lists.Sort()

		} else {
			fmt.Printf(" Новая карта другого банка была добавлена\n")
			Lists.Sort()

		}

	}

}

/*
func hw() {
	var T int
	fmt.Println("Введите количество наборов")
	fmt.Scan(&T)
	for i := 1; i <= T; i++ {
		fmt.Println("Набор №", i)
		fmt.Println("Введите слово из 4 латинских букв:")
		var input string
		fmt.Scan(&input)
		for _, v := range input {
			switch v {
			case int32(input[1]):
				if int32(input[2]) == int32(input[3]) {
					fmt.Println("yes")
					fmt.Println(string(v), string(input[1]), " & ", string(input[2]), string(input[3]))
				}
			case int32(input[2]):
				if int32(input[1]) == int32(input[3]) {
					fmt.Println("yes")
					fmt.Println(string(v), string(input[2]), " & ", string(input[1]), string(input[3]))
				}
			case int32(input[3]):
				if int32(input[2]) == int32(input[1]) {
					fmt.Println("yes")
					fmt.Println(string(v), string(input[3]), " & ", string(input[2]), string(input[1]))
				}
			default:
				fmt.Println("no")
			}
		}
	}
}

func cw() {
	var T int
	m := make(map[uint8]string)
	fmt.Println("Введите количество наборов")
	fmt.Scan(&T)
	for i := 1; i <= T; i++ {
		fmt.Println("Набор №", i)
		fmt.Println("Введите слово из 4 латинских букв:")
		var input string
		fmt.Scan(&input)
		m[input[0]] = "first"
		m[input[1]] = "second"
		m[input[2]] = "third"
		m[input[3]] = "fourth"
		fmt.Println(m)
		if len(m) == 2 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
*/
