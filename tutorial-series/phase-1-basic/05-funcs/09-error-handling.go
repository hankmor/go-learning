package main

import (
	"errors"
	"fmt"
	"os"
)

// 定义自定义错误
var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
)

// divide 基本错误处理
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// divideWithWrap 使用 %w 包装错误
func divideWithWrap(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("failed to divide %d by %d: %w", a, b, errors.New("division by zero"))
	}
	return a / b, nil
}

// User 用户结构体
type User struct {
	ID   int
	Name string
	Age  int
}

// findUser 模拟查找用户
func findUser(id int) (*User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid user id %d: %w", id, ErrInvalidInput)
	}
	if id > 100 {
		return nil, fmt.Errorf("user %d: %w", id, ErrNotFound)
	}
	return &User{ID: id, Name: "User" + fmt.Sprint(id), Age: 20}, nil
}

// ValidationError 自定义错误类型
type ValidationError struct {
	Field string
	Value interface{}
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s (value: %v)",
		e.Field, e.Msg, e.Value)
}

// validateAge 验证年龄
func validateAge(age int) error {
	if age < 0 || age > 150 {
		return &ValidationError{
			Field: "age",
			Value: age,
			Msg:   "must be between 0 and 150",
		}
	}
	return nil
}

// validateUser 验证用户信息
func validateUser(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if user.Name == "" {
		return &ValidationError{
			Field: "name",
			Value: user.Name,
			Msg:   "cannot be empty",
		}
	}
	if err := validateAge(user.Age); err != nil {
		return fmt.Errorf("user validation failed: %w", err)
	}
	return nil
}

// readConfig 模拟读取配置文件
func readConfig(filename string) error {
	_, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", filename, err)
	}
	return nil
}

// processUser 处理用户，演示错误链
func processUser(id int) error {
	user, err := findUser(id)
	if err != nil {
		return fmt.Errorf("process user failed: %w", err)
	}

	if err := validateUser(user); err != nil {
		return fmt.Errorf("process user failed: %w", err)
	}

	fmt.Printf("Processing user: %+v\n", user)
	return nil
}

// errorIsDemo 演示 errors.Is 的使用
func errorIsDemo() {
	fmt.Println("\n=== errors.Is Demo ===")

	// 测试 findUser
	_, err := findUser(200)
	if err != nil {
		// 判断是否是特定错误
		if errors.Is(err, ErrNotFound) {
			fmt.Println("User not found, creating new one...")
		} else if errors.Is(err, ErrInvalidInput) {
			fmt.Println("Invalid input provided")
		} else {
			fmt.Println("Unexpected error:", err)
		}
	}

	// 测试 processUser
	err = processUser(200)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Process failed: user not found")
		}
	}
}

// errorAsDemo 演示 errors.As 的使用
func errorAsDemo() {
	fmt.Println("\n=== errors.As Demo ===")

	user := &User{ID: 1, Name: "", Age: 200}
	err := validateUser(user)

	if err != nil {
		// 尝试将错误转换为 ValidationError
		var validationErr *ValidationError
		if errors.As(err, &validationErr) {
			fmt.Printf("Validation error: field=%s, value=%v, msg=%s\n",
				validationErr.Field, validationErr.Value, validationErr.Msg)
		} else {
			fmt.Println("Not a validation error:", err)
		}
	}
}

// multipleErrorsDemo 演示处理多个错误
func multipleErrorsDemo() {
	fmt.Println("\n=== Multiple Errors Demo ===")

	var errs []error

	// 收集多个错误
	if err := validateAge(-1); err != nil {
		errs = append(errs, err)
	}

	if err := validateAge(200); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		fmt.Println("Found", len(errs), "errors:")
		for i, err := range errs {
			fmt.Printf("  %d. %v\n", i+1, err)
		}
	}
}

// panicRecoverDemo 演示 panic 和 recover
func panicRecoverDemo() {
	fmt.Println("\n=== Panic and Recover Demo ===")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("something went wrong")
	fmt.Println("This will not be printed")
}

func main() {
	// 1. 基本错误处理
	fmt.Println("=== Basic Error Handling ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 2. 错误包装
	fmt.Println("\n=== Error Wrapping ===")
	_, err = divideWithWrap(10, 0)
	if err != nil {
		fmt.Println("Wrapped error:", err)
	}

	// 3. 自定义错误
	fmt.Println("\n=== Custom Errors ===")
	err = validateAge(200)
	if err != nil {
		fmt.Println("Validation error:", err)
	}

	// 4. errors.Is
	errorIsDemo()

	// 5. errors.As
	errorAsDemo()

	// 6. 多个错误
	multipleErrorsDemo()

	// 7. panic 和 recover
	panicRecoverDemo()
	fmt.Println("Program continues after panic recovery")

	// 8. 最佳实践
	fmt.Println("\n=== Best Practices ===")
	fmt.Println("1. 总是检查错误，不要忽略")
	fmt.Println("2. 使用 %w 包装错误以保留错误链")
	fmt.Println("3. 使用 errors.Is 和 errors.As 判断错误类型")
	fmt.Println("4. 自定义错误类型要实现 Error() 方法")
	fmt.Println("5. panic 只用于不可恢复的错误")
	fmt.Println("6. 在库代码中返回错误，不要 panic")
}
