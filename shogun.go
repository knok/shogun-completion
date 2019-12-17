// shogun.go
//
// Copyright 2013, 2015 NOKUBI Takatsugu <knok@daionet.gr.jp>
//
// Copying and distribution of this file, with or without modification,
// are permitted in any medium without royalty provided the copyright
// notice and this notice are preserved.  This file is offered as-is,

package main

import (
	"fmt"
	"io"
	"os"
)

var (
	tokugawa = map[string]int{
		"徳川家康": 1,
		"徳川秀忠": 2,
		"徳川家光": 3,
		"徳川家綱": 4,
		"徳川綱吉": 5,
		"徳川家宣": 6,
		"徳川家継": 7,
		"徳川吉宗": 8,
		"徳川家重": 9,
		"徳川家治": 10,
		"徳川家斉": 11,
		"徳川家慶": 12,
		"徳川家定": 13,
		"徳川家茂": 14,
		"徳川慶喜": 15,
	}

	kamakura = map[string]int{
		"源頼朝":  1,
		"源頼家":  2,
		"源実朝":  3,
		"藤原頼経": 4,
		"藤原頼嗣": 5,
		"宗尊親王": 6,
		"惟康親王": 7,
		"久明親王": 8,
		"守邦親王": 9,
	}
)

func Tokugawa(name string) (string, error) {
	val, ok := tokugawa[name]
	if !ok {
		return "", fmt.Errorf("そんな将軍はいない: %v", name)
	}
	dai := Dai(val)
	return fmt.Sprintf("%s (the %s Shogun)", name, dai), nil
}

func Kamakura(name string) (string, error) {
	val, ok := kamakura[name]
	if !ok {
		return "", fmt.Errorf("そんな将軍はいない: %v", name)
	}
	dai := Dai(val)
	return fmt.Sprintf("%s (the %s Shogun)", name, dai), nil
}

func Dai(num int) string {
	var suffix = "th"
	x := num % 10
	if x == 1 {
		suffix = "st"
	} else if x == 2 {
		suffix = "nd"
	} else if x == 3 {
		suffix = "rd"
	}
	return fmt.Sprintf("%d%s", num, suffix)
}

var (
	stdout io.Writer = os.Stdout
	stderr io.Writer = os.Stderr
)

func run(args []string) int {
	if len(args) == 0 {
		return 1
	}
	first := args[0]
	second := ""
	var fn func(string) (string, error)

	if first == "kamakura" {
		if len(args) != 2 {
			return 1
		}
		second = args[1]
		fn = Kamakura
	} else {
		if len(args) != 1 {
			return 1
		}
		second = args[0]
		fn = Tokugawa
	}

	result, err := fn(second)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 2
	}
	fmt.Fprintln(stdout, result)
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
