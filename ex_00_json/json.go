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

	fmt.Println(strs)             // 打印字符串切片
	fmt.Println(strsJson)         // 打印JSON格式的 byte切片
	fmt.Println(string(strsJson)) // 打印JSON格式的 字符串

	fmt.Println(string(intsJson))
	fmt.Println(string(mapsJson))

	res1 := response1{
		Page:   1,
		Fruits: []string{"apple", "peach"},
	}
	res1Json, _ := json.Marshal(res1)
	fmt.Println(string(res1Json))

	res2 := response2{
		State:  true,
		Page:   1,
		Fruits: []string{"apple", "peach"},
	}
	res2Json, _ := json.Marshal(res2)
	fmt.Println(string(res2Json))

	// 反序列化
	jsonStr := `{"num":6.13,"strs":["a","b"],"mapkey":{"submap":"value"}}`

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	num := data["num"].(float64)
	str := data["strs"].([]interface{})
	mapkey := data["mapkey"].(map[string]interface{})

	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(str[0].(string))
	fmt.Println(mapkey["submap"].(string))

	// 反序列化到结构体
	jsonResStr := `{"state":false, "page":1, "fruits":["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(jsonResStr), &res)
	fmt.Println(res)
	fmt.Println(res.State)
	fmt.Println(res.Fruits)

}
