package cmd

import (
	"encoding/json"
	"fmt"
	"mygin/tools"
	"reflect"
)

type Results struct {
	Code   int    `json:"code,omitempty" xml:"code,omitempty" mytag:"code1"` //转换为json时 键 Code 变为 code 如果为 0 则忽略  转换为xml时 同样的规则
	Method string `json:"method"`                                            //转换为json时 键 Method 变为 method
	Msg    string `json:"msg,omitempty"`                                     //转换为json时 键 Msg 变为 msg 如果为 空字符串 则忽略
}

func syncCreateStruct() {
	tmpStuct := []reflect.StructField{
		{
			Name: "Height",
			Type: tools.GetTypeByName("string"),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: tools.GetTypeByName("int"),
			Tag:  `json:"age"`,
		},
		{
			Name: "Test",
			Type: tools.GetTypeByName("json.RawMessage"),
			Tag:  `json:"test"`,
		},
	}
	typ := reflect.StructOf(tmpStuct)
	fmt.Printf("%v", typ)
}

func testJSON() {
	results := Results{
		Method: "POST",
		Msg:    "",
	}
	jsonStr, _ := json.Marshal(results)
	fmt.Println(string(jsonStr))

	result1s := Results{Code: 200}
	t := reflect.TypeOf(result1s)         //利用反射获得 tag 信息
	field, found := t.FieldByName("Code") //获得struct中Code的tag信息
	fmt.Println(found)                    //获得结果  没有key 则为false
	fmt.Println(field.Tag.Get("mytag"))   //获得的内容

	fmt.Println(reflect.TypeOf(string("")))
	fmt.Println(reflect.TypeOf(int64(0)))
	fmt.Println(reflect.TypeOf(json.RawMessage("")))

	fmt.Println(tools.GetTypeByName("string"))

	syncCreateStruct()
}
