package main

import (
	"encoding/json" // JSON 包
	"fmt"
)

// 默认使用字段名作为 JSON 的键名
type response1 struct {
	Page   int
	Fruits []string
}

// `json` 标签可以自定义键名(json输出时使用)
type response2 struct {
	State  bool     `json:"state"`
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	strs := []string{"two", "one", "three"}
	ints := []int{3, 2, 4}
	maps := map[string]string{"k1": "v1，中文", "k2": "v2", "k3": "v3"}

	strsJson, _ := json.Marshal(strs) // 将字符串切片转换为 JSON 格式的byte切片
	intsJson, _ := json.Marshal(ints)
	mapsJson, _ := json.Marshal(maps)

	fmt.Println("1", strs)             // 打印字符串切片
	fmt.Println("2", strsJson)         // 打印JSON格式的 byte切片
	fmt.Println("3", string(strsJson)) // 打印JSON格式的 字符串

	fmt.Println("4", string(intsJson))
	fmt.Println("5", string(mapsJson))

	res1 := response1{
		Page:   1,
		Fruits: []string{"apple", "peach"},
	}
	res1Json, _ := json.Marshal(res1)
	fmt.Println("6", string(res1Json))

	res2 := response2{
		State:  true,
		Page:   1,
		Fruits: []string{"apple", "peach"},
	}
	res2Json, _ := json.Marshal(res2)
	fmt.Println("7", string(res2Json))

	// 反序列化 - 将JSON字符串转换为Go数据结构
	jsonStr := `{"num":6.13,"strs":["a","b"],"mapkey":{"submap":"value"}}`

	// map类型变量用于存储解析后的JSON数据
	var data map[string]interface{}

	// 使用json.Unmarshal将JSON字符串解析为Go数据结构，第一个参数是JSON字符串的字节切片，第二个参数是目标变量的指针
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		panic(err)
	}
	fmt.Println("8", data)

	// 类型断言 - 从map中获取特定类型的值
	num := data["num"].(float64)
	str := data["strs"].([]interface{})
	mapkey := data["mapkey"].(map[string]interface{})

	fmt.Println("9", num)
	fmt.Println("10", str)
	fmt.Println("11", str[0].(string))
	fmt.Println("12", mapkey["submap"].(string))

	// 反序列化到结构体 - 将JSON直接解析为预定义的结构体
	jsonResStr := `{"state":false, "page":1, "fruits":["apple", "peach"]}` // 定义一个与response2结构体匹配的JSON字符串
	res := response2{}
	json.Unmarshal([]byte(jsonResStr), &res) // 将JSON解析到结构体中，字段会自动匹配
	fmt.Println("13", res)
	fmt.Println("14", res.State)
	fmt.Println("15", res.Fruits)

}
