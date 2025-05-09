package builder

// ConcreteBuilder 是Builder接口的具体实现
type ConcreteBuilder struct {
	computer *Computer
}

// NewConcreteBuilder 初始化并返回具体的建造者实例
func NewConcreteBuilder() Builder {
	return &ConcreteBuilder{computer: &Computer{}}
}

func (b *ConcreteBuilder) SetCPU(cpu string) Builder {
	b.computer.CPU = cpu
	return b
}

func (b *ConcreteBuilder) SetRAM(ram int) Builder {
	b.computer.RAM = ram
	return b
}

func (b *ConcreteBuilder) SetStorage(storage int) Builder {
	b.computer.Storage = storage
	return b
}

func (b *ConcreteBuilder) SetGPU(gpu string) Builder {
	b.computer.GPU = gpu
	return b
}

func (b *ConcreteBuilder) Build() *Computer {
	return b.computer
}
