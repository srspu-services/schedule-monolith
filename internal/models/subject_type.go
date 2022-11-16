package model

type SubjectType int

const (
	Leacture SubjectType = iota + 1
	Practice
	Laboratory
)

func (s SubjectType) String() string {
	return [...]string{"Лекция", "Практика", "Лабораторная"}[s-1]
}

func (s SubjectType) EnumIndex() int {
	return int(s)
}
