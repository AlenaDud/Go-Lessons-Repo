package main

import (
	"fmt"
	"strconv"
)

func main() {
	ageLong()
	ageShort()
	name()
	temperature()
	sunny()
	evaluate()
	shuffleNum()
	evalArea()
	totalSum()
	pointerCount()
	defaultStrVal()
	defaultVal()
	typeNumCoercion()
	typeStrCoercion()
}

// 1. Объявите переменную age типа int и присвойте ей значение 30.
func ageLong() {
	var age int
	age = 30
	fmt.Println(age)
}

func ageShort() {
	age := 30
	fmt.Println(age)
}

// 2. Объявите переменную name типа string и присвойте ей ваше имя.
func name() {
	var name string
	name = "Алёна"
	fmt.Println(name)
}

// 3. Объявите переменную temperature типа float64 и присвойте ей значение 25.5.
func temperature() {
	var temperature float64
	temperature = 25.5
	fmt.Println(temperature)
}

// 4. Объявите переменную isSunny типа bool и присвойте ей значение true.
func sunny() {
	var isSunny bool
	isSunny = true
	fmt.Println(isSunny)
}

// 5. Объявите переменную x и присвойте ей значение 10. Затем умножьте x на 5 и присвойте результат переменной y.
func evaluate() {
	var x, y int
	x = 10
	y = x * 5
	fmt.Println(x, y)
}

// 6. Объявите переменные a и b, затем обменяйте их значения без использования дополнительных переменных.
func shuffleNum() {
	var a, b = 1, 2
	a, b = b, a
}

/*
7. Объявите переменные width и height, представляющие ширину и высоту прямоугольника.
Затем вычислите площадь прямоугольника и присвойте результат переменной area.
*/
func evalArea() {
	var width, height, area int
	area = width * height
	fmt.Println(area)
}

// 8. Создайте переменную total и используйте арифметические операторы для подсчета суммы чисел от 1 до 10.
func totalSum() {
	var total int
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println(total)
}

/*
 9. Объявите переменную count типа int и присвойте ей значение 42. Затем
    создайте указатель на count и выведите значение, на которое он указывает.
*/
func pointerCount() {
	var count int
	count = 42
	var countPointer *int
	countPointer = &count
	fmt.Println(*countPointer)
}

/*
10. Объявите переменную defaultValue типа string и не присваивайте ей значение.
Затем выведите значение defaultValue. Что будет выведено?
*/
func defaultStrVal() {
	var defaultValue string
	fmt.Printf("%#v", defaultValue)
	// stdout: ""
}

/*
11. Объявите переменную zeroInt типа int и переменную zeroBool типа bool.
Затем выведите их значения. Что будет выведено?
*/
func defaultVal() {
	var zeroInt int
	var zeroBool bool
	fmt.Printf("%#v", zeroInt)
	fmt.Printf("%#v", zeroBool)
	// stdout: 0, false
}

/*
12. Объявите переменную number типа float64 и присвойте ей значение 42.5.
Затем приведите её к типу int и выведите результат.
*/
func typeNumCoercion() {
	var number float64
	var numberInt int
	number = 42.5
	numberInt = int(number)
	fmt.Println(numberInt)
}

/*
13. Объявите переменную stringValue типа string с текстом "123".
Затем преобразуйте её в число и выполните арифметическую операцию.
*/
func typeStrCoercion() {
	var stringValue string
	stringValue = "123"
	stringNumber, _ := strconv.Atoi(stringValue)
	stringNumber *= stringNumber
	fmt.Println(stringNumber)
}
