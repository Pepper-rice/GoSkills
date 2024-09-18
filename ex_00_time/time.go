package main

import (
	"fmt"
	"time"
)

// 自定义时间格式常量
const (
	lay1 = "2006年01月02日 15时04分05秒"      // 中文格式
	lay2 = "2006-01-02 15:04:05"        // 常用格式
	lay3 = "2006-01-02 15:04:05.999999" // 精确到微秒的格式
)

func main() {
	// 1. 获取当前时间
	now := time.Now() // Now() 获取当前系统时间
	fmt.Println("当前时间 (默认格式):", now)

	// 2. 获取时间的各个部分
	fmt.Println("当前年份:", now.Year())
	fmt.Println("当前月份 (英文):", now.Month())
	fmt.Println("当前日期 (几号):", now.Day())
	fmt.Println("当前小时:", now.Hour())

	// 3. 获取时间戳
	fmt.Println("当前时间的秒级时间戳:", now.Unix())       // 返回从1970年到现在的秒数
	fmt.Println("当前时间的毫秒级时间戳:", now.UnixMilli()) // 返回从1970年到现在的毫秒数
	fmt.Println("当前时间的纳秒级时间戳:", now.UnixNano())  // 返回从1970年到现在的纳秒数

	// 4. 使用RFC3339标准格式输出时间
	fmt.Println("当前时间 (RFC3339标准格式):", now.Format(time.RFC3339))

	// 5. 使用自定义格式化时间输出
	dt1 := now.Format(lay1)
	dt2 := now.Format(lay2)
	dt3 := now.Format(lay3)
	fmt.Println("当前时间 (中文格式):", dt1)
	fmt.Println("当前时间 (常用格式):", dt2)
	fmt.Println("当前时间 (精确到微秒):", dt3)

	fmt.Println("---")

	// 6. 获取当前时间的时区信息
	zone, offset := now.Zone() // Zone() 获取当前时间的时区和偏移量
	fmt.Println("当前时区:", zone)
	fmt.Println("当前时区偏移秒数:", offset)

	// 7. 时区转换
	cstZone := time.FixedZone("CST", 8*3600) // 创建一个固定的时区 CST
	fmt.Println("当前时间 (中文格式):", now.Format(lay1))
	fmt.Println("当前时间转换为CST时区 (中文格式):", now.In(cstZone).Format(lay1))

	fmt.Println("---")

	// 8. 字符串解析成时间对象
	dtStr := "2017年04月07日 16时30分47秒"
	utcTime, _ := time.Parse(lay1, dtStr) // Parse() 返回 UTC 时间
	fmt.Println("解析的 UTC 时间 (常用格式):", utcTime.Format(lay2))
	fmt.Println("解析的 UTC 时间戳 (秒级):", utcTime.Unix())

	// 9. 正确的时间解析方式
	cstTime, _ := time.ParseInLocation(lay1, dtStr, time.Local) // ParseInLocation() 返回本地时间
	fmt.Println("解析的 CST 时间 (常用格式):", cstTime.Format(lay2))
	fmt.Println("解析的 CST 时间戳 (秒级):", cstTime.Unix())

	fmt.Println("---")

	// 10. 时间戳转字符串，字符串转时间戳
	unixstamp := 1609527845
	uTime := time.Unix(int64(unixstamp), 0)
	fmt.Println("Unix 时间戳转字符串 (中文格式):", uTime.Format(lay1))

	// 字符串转回时间戳
	uStr := uTime.Format(lay1)
	us, _ := time.ParseInLocation(lay1, uStr, time.Local)
	fmt.Println("字符串转回 Unix 时间戳 (中文格式):", us.Format(lay1))
	fmt.Println("字符串转回 Unix 时间戳 (秒级):", us.Unix())
}
