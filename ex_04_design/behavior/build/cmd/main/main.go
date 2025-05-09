package main

import (
	"fmt"
	"github.com/learning-go/ex_04_design/behavior/build/internal/builder"
)

func main() {
	// 自定义构建电脑
	customComputer := builder.NewConcreteBuilder().
		SetCPU("AMD Ryzen 9 7945").
		SetRAM(32).
		SetStorage(2048).
		SetGPU("Radeon RX 6900 XT").
		Build()
	fmt.Printf("Custom Computer: %+v\n", customComputer)

	// 通过Director构建电脑
	concreteBuilder := builder.NewConcreteBuilder()
	director := builder.NewDirector(concreteBuilder)

	//构建游戏电脑
	gamingComputer := director.BuildGamingComputer()
	fmt.Printf("Gaming Computer: %+v\n", gamingComputer)
	//构建办公电脑
	officeComputer := director.BuildOfficeComputer()
	fmt.Printf("Office Computer: %+v\n", officeComputer)
}
