package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"math/rand"
	"time"
)

var weathers = []string{"sunny", "windy", "cloudy", "rainy"}

var beforeSayHello = func(ctx *cli.Context) error {
	fmt.Println("Before command sayHello...")
	return nil
}

func HelloCmd() *cli.Command {
	return &cli.Command{
		Name:    "hello",        // 命令名称，执行时需要指定
		Aliases: []string{"ho"}, // 命令别名，简化名称
		Usage:   "Say hello to you",
		Before:  beforeSayHello,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "your `NAME`", Required: true},
		},
		Subcommands: cli.Commands{reportWeatherCmd(), complainWeatherCmd()}, // 子命令
		Action:      sayHello,                                               // 具体命令的执行逻辑
	}
}

func sayHello(ctx *cli.Context) error {
	name := ctx.String("n")
	fmt.Println("hello,", name, "!")
	return nil
}

func reportWeatherCmd() *cli.Command {
	return &cli.Command{
		Name:    "report-weather", // 命令名称，执行时需要指定
		Aliases: []string{"rw"},   // 命令别名，简化名称
		Usage:   "Report the weather today",
		Before: func(ctx *cli.Context) error {
			fmt.Println("Before subcommand of sayHello weatherCmd...")
			return nil
		},
		Flags:    []cli.Flag{},
		Action:   reportWeather,
		Category: "weather",
	}
}

func complainWeatherCmd() *cli.Command {
	return &cli.Command{
		Name:    "complain-weather", // 命令名称，执行时需要指定
		Aliases: []string{"cw"},     // 命令别名，简化名称
		Usage:   "Complains the weather today",
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Flags:    []cli.Flag{},
		Action:   complainWeather,
		Category: "weather",
	}
}

func reportWeather(ctx *cli.Context) error {
	name := ctx.String("n")
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	weather := weathers[rd.Intn(len(weathers))]
	fmt.Printf("hello %s, today is a %s day!\n", name, weather)
	if ctx.Bool("i") {
		fmt.Println("this is verbose info")
	}
	return nil
}

func complainWeather(ctx *cli.Context) error {
	name := ctx.String("n")
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	weather := weathers[rd.Intn(len(weathers))]
	fmt.Printf("hello %s, today is a %s day, oh, i don't like it!\n", name, weather)
	if ctx.Bool("i") {
		fmt.Println("this is verbose info")
	}
	return nil
}
