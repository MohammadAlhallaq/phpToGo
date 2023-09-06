package interfaces

type Programmer interface {
	Speak() string
}

type PhpProgrammer struct {
}

func (p *PhpProgrammer) Speak() string {
	return "this is php"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "this is java"

}
