package builder

// Director 封装复杂的建造流程
type Director struct {
	builder Builder
}

// NewDirector 返回导演实例
func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

// BuildGamingComputer 标准游戏电脑配置
func (d *Director) BuildGamingComputer() *Computer {
	return d.builder.
		SetCPU("Intel i9").
		SetRAM(32).
		SetStorage(1024).
		SetGPU("RTX 4090").
		Build()
}

// BuildOfficeComputer 标准办公电脑配置
func (d *Director) BuildOfficeComputer() *Computer {
	return d.builder.
		SetCPU("Intel i5").
		SetRAM(16).
		SetStorage(512).
		SetGPU("Integrated GPU").
		Build()
}
