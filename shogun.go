package main

import (
	"fmt"
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
		"源頼朝": 1,
		"源頼家": 2,
		"源実朝": 3,
		"藤原頼経": 4,
		"藤原頼嗣": 5,
		"宗尊親王": 6,
		"惟康親王": 7,
		"久明親王": 8,
		"守邦親王": 9,
	}
	)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}
	first := os.Args[1]
	val, ok := tokugawa[first]
	if ok {
		var dai string
		dai = fmt.Sprintf("%dth", val)
		if val == 1 || val == 11 {
			dai = fmt.Sprintf("%dst", val)
		} else if val == 2 || val == 12 {
			dai =fmt.Sprintf("%dnd",val)
		} else if val == 3 || val == 13 {
			dai = fmt.Sprintf("%drd", val)
		}
		fmt.Printf("%s (the %s Shogun)\n", first, dai)
		os.Exit(0)
	}
	os.Exit(2)
}
