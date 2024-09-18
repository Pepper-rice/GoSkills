package main

import "fmt"

type rect struct {
	width, height int
}

// area() 是一个拥有 *rect 类型的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 还可以通过值类型定义方法
func (r rect) perim() int {
	return (r.width + r.height) * 2
}

// 使用指针接收者的方法，可以修改接收者所指向的实际值
func (r *rect) scale(factor int) {
	r.width *= factor
	r.height *= factor
}

// 使用值接收者的方法，不会修改接收者的实际值（因为是副本）
func (r rect) scaleValue(factor int) {
	r.width *= factor
	r.height *= factor
}

// Go 会自动处理值和指针之间的转换，想要避免调用方法时产生一个拷贝。都可以使用指针调用方法
func main() {
	r := rect{width: 10, height: 20}

	fmt.Println("area:", r.area())   // 此调用Go会自动将 r 转换为指针类型 &r，因为 area 方法的接收者是指针类型（*rect）
	fmt.Println("perim:", r.perim()) // 此调用不需要转换，因为 perim 方法的接收者是值类型（rect）
	fmt.Printf("Address of r: %p\n", &r)

	pr := &r
	fmt.Println("area:", pr.area())   // 由于 pr 本身就是 *rect 类型，所以不需要转换
	fmt.Println("perim:", pr.perim()) // Go会自动将指针解引用为 rect 值类型，因为 perim 方法的接收者是值类型
	fmt.Printf("Address of pr: %p\n", pr)

	// 调用指针接收者方法
	r.scale(2)
	fmt.Println("After scaling with pointer receiver:", r) // r的值被修改了

	// 调用值接收者方法
	r.scaleValue(2)
	fmt.Println("After scaling with value receiver:", r) // r真实的值没有改变

	// 通过指针调用值接收者方法（Go自动解引用）
	pr.scaleValue(2)
	fmt.Println("After scaling with value receiver (pointer call):", r) // r真实的值没有改变

	// 通过指针调用指针接收者方法
	pr.scale(2)
	fmt.Println("After scaling with pointer receiver (pointer call):", r) // r的值被修改了
}
