package flyweight

import "sync"

// 大致的思路就是将可变和不可变的变量进行拆分,从而使得不可变的这一块可以被多个地方引用(同一个进程,用户态地址相同,那么数据也相同)

// 假设当前有a、b、c 三个班级

// 学校人员抽象方法
type School interface {
	Name() string
	Position() string
	Class
}

// 班级抽象接口
type Class interface {
	Class() string
}

// 班级结构
type class struct {
	name string
}

func (c *class) Class() string {
	return c.name
}

// 默认工厂
var defaultFactory = &classFactory{m: map[string]Class{
	"A班": newClass("A班"),
}}

type classFactory struct {
	m map[string]Class
	sync.Mutex
}

func newClass(className string) Class {
	return &class{
		name: className,
	}
}

func GetClass(className string) Class {
	var val Class
	var ok bool

	// tip: 这一块懒得用,如果是真想延迟初始化,那么建议用sync.map的逻辑,双map处理。否则锁开支还是很大的。如果能用常量加数组那就更快了。
	//defaultFactory.Lock()
	if val, ok = defaultFactory.m[className]; ok {
		return val
	}

	//val =&class{
	//    name: className,
	//}
	//defaultFactory.m[className] = val
	//defaultFactory.Unlock()
	return val
}

// 学生结构
type student struct {
	class Class
	name  string
}

func (s student) Position() string {
	return "学生"
}
func (s student) Name() string {
	return s.name
}
func (s student) Class() string {
	return s.class.Class()
}
func NewStudent(name string, class Class) School {
	return student{
		class: class,
		name:  name,
	}

}

type teacher struct {
	class Class
	name  string
}

func (s teacher) Position() string {
	return "教师"
}
func (s teacher) Name() string {
	return s.name
}
func (s teacher) Class() string {
	return s.class.Class()
}
func NewTeacher(name string, class Class) School {
	return teacher{
		class: class,
		name:  name,
	}

}
