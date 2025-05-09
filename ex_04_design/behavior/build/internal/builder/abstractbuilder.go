// BuildPart 根据传入的部件名称返回部件字符串
package builder

// Builder 接口定义了所有构建方法
type Builder interface {
	SetCPU(cpu string) Builder
	SetRAM(ram int) Builder
	SetStorage(storage int) Builder
	SetGPU(gpu string) Builder
	Build() *Computer
}
