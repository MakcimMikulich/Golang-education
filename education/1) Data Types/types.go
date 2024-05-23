package main

import "fmt"

func main() {
    // Булевый тип
    var b bool = true
    fmt.Println(b)

    // Целые числа
    var i int = 42
    fmt.Println(i)
    
    var u uint = 42
    fmt.Println(u)
    
    // Числа с плавающей точкой
    var f float64 = 3.14
    fmt.Println(f)
    
    // Комплексные числа
    var c complex128 = complex(1, 2)
    fmt.Println(c)
    
    // Строки
    var s string = "Hello, Go!"
    fmt.Println(s)
    
    // Массивы
    arr := [3]int{1, 2, 3}
    fmt.Println(arr)
    
    // Срезы
    var slice []int = []int{1, 2, 3}
    slice := []int{1, 2, 3}
    fmt.Println(slice)

    slice := make([]int, 5) // Срез длиной 5 и емкостью 5
    slice := make([]int, 5, 10) // Срез длиной 5 и емкостью 10

    slice = append(slice, 5)
    slice = append(slice, 6, 7, 8)


    copySlice := make([]int, len(slice))
    copy(copySlice, slice)
    fmt.Println("Copied slice:", copySlice)
    
    // Карты
    var m map[string]int = map[string]int{"one": 1, "two": 2}
    fmt.Println(m)
    
    // Структуры
    type Person struct {
        Name string
        Age  int
    }
    var p Person = Person{"Alice", 30}
    fmt.Println(p)
    
    // Указатели
    var ptr *int = &i
    fmt.Println(ptr)
    fmt.Println(*ptr)
    
    // Интерфейсы
    type Describer interface {
        Describe() string
    }
    
    func (p Person) Describe() string {
        return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
    }
    
    var d Describer = p
    fmt.Println(d.Describe())
}