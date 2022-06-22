package main

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name     string `bilibili:"BILIBILI_NAME"`
	PublicWX string `bilibili:"BILIBILI_PUBLICWX"`
}

func PrintTag(ptr interface{}) {
	reType := reflect.TypeOf(ptr)
	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		panic("gg")
		return
	}
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag

		laberTag := tag.Get("bilibili")
		fmt.Println(laberTag)
	}

}

func main() {
	userinfo := &UserInfo{"其味无穷二无群", "cncf_kuqweqw"}
	PrintTag(userinfo)
}
