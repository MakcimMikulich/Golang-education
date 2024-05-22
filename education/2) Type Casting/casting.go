package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	convInNumber()
	convInString()
}

func convInNumber(){
   var str string = "5"
   num, _ := strconv.Atoi(str)
   fmt.Println(num * 2) // вывод 10
}

func convInString(){
	var num int = 5
	str := strconv.Itoa(num)
	fmt.Println(str+str) // вывод 55
 }


 func convInInt(){
	var a float64 = 6.6
	b := int(a) // приведение (casting) вещественного числа к целому
	fmt.Println(b) // Вывод 6
 }

 func convInFloat(){
	var a int = 6
	b := float64(a) // приведение (casting) целого числа к вещественному
	fmt.Println(b + 0.4) // Вывод 6.4
 }


 func divisionFloat(){
	var n int = 5 
	var floatDivider float32 = 2.0
	fmt.Println(float32(n) / floatDivider) // 2.5
 }

 var b float64 = math.Abs(b) // модуль числа

var a float64 = math.Pow(b, 2) // возводим переменную b в степень 2

var a float64 = math.Sqrt(b) // корень квадратный

var b float64 = math.Min(a, b) // минимум из чисел а и b (для максимума math.Max(a, b))