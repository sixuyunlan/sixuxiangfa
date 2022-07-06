package main

import "fmt"

func main() {
	b := map[int]string{
		20081234: "lina",
		20081314: "silou",
		20091478: "qwee",
		20011234: "s12ASQ",
	}
	delete(b, 20081234)
	var key string

	fmt.Scanln(&key)

	fmt.Println(b)
	for k, v := range b {
		fmt.Println("k is", k, "v is", v)
		fmt.Printf("key为：%v value为%v \n", k, v)
	}

	a := map[string]map[int]string{
		"第一班": {
			101: "丽丽",
			102: "鹿鹿",
			103: "欢欢",
		},
		"第二班": {
			201: "王尼玛",
			202: "张老三",
			203: "李扒皮",
		},
	}

	for k1, v1 := range a {
		fmt.Println(k1, v1)
	}
	for k2, v2 := range a {
		fmt.Println(k2)
		for k3, v3 := range v2 {
			fmt.Println(k3, v3)
		}
	}
	cha, zhao := a["第一班"]
	fmt.Println(cha)
	fmt.Println(zhao)
}
