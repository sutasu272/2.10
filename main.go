package main

import "fmt"

func main() {
	// Задание 1
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}
	// Задание 2
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	// Задание 3
	var number int = 5
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", number * i,)
	}
	fmt.Println()
	// Задание 4
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// Задание 5
	var number2 int
	fmt.Scan(&number2)
	count := 0
	for number2 > 0 {
		number2 /= 10
		count++
	}
	fmt.Println("Количество цифр:", count)
	// Задание 6
	text := "Developer"
	for _, char := range text {
		fmt.Printf("\n%c", char)
	}
	fmt.Println()
	// Задание 7
	balance := 3000
	var number3 int
	for {fmt.Print("Введите число от 0 до 3: ")
	fmt.Scan(&number3)
	switch number3 {
		case 0:
			fmt.Println("Выход из программы.")
			return
		case 1:
			fmt.Printf("Баланс: %d\n", balance)
		case 2:
			balance += 500
		case 3:
			balance -= 200	
	}
}
	
}
