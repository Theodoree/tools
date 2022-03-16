package factory

import "errors"

// factory
type PeopleEnum uint64

const (
    Teacher PeopleEnum = iota + 1
    Student
    Professor
    Principal
)

var WrongEnumType = errors.New("错误的枚举类型")

var (
    _teacherFactory   = teacherFactory{}
    _studentFactory   = studentFactory{}
    _professorFactory = professorFactory{}
    _principalFactory = principalFactory{}
)

type People interface {
    Eat()
    Say()
}

type PeopleFactory interface {
    Create() People
}

func Say(enum PeopleEnum) error {
    factory, err := choiceFactory(enum)
    if err != nil {
        return err
    }

    instance := factory.Create()
    instance.Say()
    return nil
}

func choiceFactory(enum PeopleEnum) (PeopleFactory, error) {
    switch enum {
    case Teacher:
        return _teacherFactory, nil
    case Student:
        return _studentFactory, nil
    case Professor:
        return _professorFactory, nil
    case Principal:
        return _principalFactory, nil
    default:
        return nil, WrongEnumType
    }

}

type teacher struct{}

func (t teacher) Say() { println("teacher") }
func (t teacher) Eat() {}

type teacherFactory struct{}

func (t teacherFactory) Create() People {
    return teacher{}
}

type student struct{}

func (t student) Say() { println("student") }
func (t student) Eat() {}

type studentFactory struct{}

func (t studentFactory) Create() People {
    return student{}
}

type professor struct{}

func (t professor) Say() { println("professor") }
func (t professor) Eat() {}

type professorFactory struct{}

func (t professorFactory) Create() People {
    return professor{}
}

type principal struct{}

func (t principal) Say() { println("principal") }
func (t principal) Eat() {}

type principalFactory struct{}

func (t principalFactory) Create() People {
    return principal{}
}
