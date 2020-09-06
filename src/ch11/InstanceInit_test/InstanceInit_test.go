package InstanceInit_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestInstanceInit(t *testing.T) {
	e1 := Employee{"1", "Bob", 20}
	t.Log(e1)
	t.Logf("e1 is %T", &e1)

	e2 := Employee{Name: "Mike", Age: 20}
	t.Log(e2)
	t.Logf("e2 is %T", &e2)

	e3 := new(Employee) // 注意这里返回的引用/指针，相当于 e3 := &Employee{}
	e3.Id = "2"         // 与其他主要编程语言的差异：通过实例的指针访问成员不需要使用->
	e3.Name = "Rose"
	e3.Age = 30
	t.Log(e3)
	t.Logf("e3 is %T", e3)
}

// 行为/方法定义
// 第一种定义方法在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String1() string {
	// 会产生对象对象复制
	fmt.Printf("Address2 is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s - Name: %s - Age: %d", e.Id, e.Name, e.Age)
}

// 通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String2() string {
	// 不会产生对象复制
	fmt.Printf("Address5 is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s / Name: %s / Age: %d", e.Id, e.Name, e.Age)
}
func TestInstanceFunc(t *testing.T) {
	e1 := Employee{"1", "Bob", 20}
	fmt.Printf("Address1 is %x\n", unsafe.Pointer(&e1.Name))
	t.Log(e1.String1())
	e2 := &Employee{"2", "Bob1", 28}
	fmt.Printf("Address3 is %x\n", unsafe.Pointer(&e2.Name))
	t.Log(e2.String1())

	e3 := new(Employee)
	e3.Id = "2"
	e3.Name = "Rose"
	fmt.Printf("Address4 is %x\n", unsafe.Pointer(&e3.Name))
	t.Log(e3.String2())
}
