package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 基本 switch
	day := "Mon"
	switch day {
	case "Mon":
		fmt.Println("Monday")
	case "Tue":
		fmt.Println("Tuesday")
	case "Wed":
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// 2. 多个值的 case
	switch day {
	case "Sat", "Sun":
		fmt.Println("Weekend!")
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("Weekday")
	}

	// 3. 带初始化语句的 switch
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// 4. 无表达式的 switch（替代 if-else 链）
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	// 5. 更复杂的条件
	x, y := 10, 20
	switch {
	case x > 0 && y > 0:
		fmt.Println("Both positive")
	case x < 0 && y < 0:
		fmt.Println("Both negative")
	case x == 0 || y == 0:
		fmt.Println("At least one is zero")
	default:
		fmt.Println("Mixed signs")
	}

	// 6. fallthrough 关键字
	num := 1
	fmt.Println("\nFallthrough demo:")
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two or after one")
	case 3:
		fmt.Println("Three")
	}

	// 7. 类型 switch
	fmt.Println("\nType switch demo:")
	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(3.14)
	checkType([]int{1, 2, 3})
}

// checkType 使用类型 switch 判断接口变量的实际类型
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	default:
		fmt.Printf("Unknown type: %T, value: %v\n", v, v)
	}
}
