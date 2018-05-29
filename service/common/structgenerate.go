package common

import (
	"reflect"
)

type StructBuilder struct {
	field []reflect.StructField
}

func NewStructBuilder() *StructBuilder {
	return &StructBuilder{}
}

func (b *StructBuilder) AddField(fname string, ftype reflect.Type) {
	b.field = append(
		b.field,
		reflect.StructField{
			Name: fname,
			Type: ftype,
		})
}

func (b *StructBuilder) Build() Struct {
	strct := reflect.StructOf(b.field)
	index := make(map[string]int)
	for i := 0; i < strct.NumField(); i++ {
		index[strct.Field(i).Name] = i
	}
	return Struct{strct, index}
}

type Struct struct {
	strct reflect.Type
	index map[string]int
}

func (s *Struct) NewInstance() *Instance {
	instance := reflect.New(s.strct).Elem()
	return &Instance{instance, s.index}
}

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (i *Instance) Field(name string) reflect.Value {
	return i.internal.Field(i.index[name])
}

func (i *Instance) SetString(name, value string) {
	i.Field(name).SetString(value)
}

func (i *Instance) SetBool(name string, value bool) {
	i.Field(name).SetBool(value)
}

func (i *Instance) SetInt(name string, value int) {
	i.Field(name).SetInt(int64(value))
}

func (i *Instance) SetFloat(name string, value float64) {
	i.Field(name).SetFloat(value)
}

func (i *Instance) Value() interface{} {
	return i.internal.Interface()
}

func (i *Instance) Pointer() interface{} {
	return i.internal.Addr().Interface()
}

/*
func main() {
	b := NewStructBuilder()
	b.AddField("Name", reflect.TypeOf("")) //型情報が欲しいだけなので、reflect.TypeOf()の入力値はなんでもいい。
	b.AddField("Age", reflect.TypeOf(123))
	person := b.Build()

	i := person.NewInstance()
	i.SetString("Name", "gopher")
	i.SetInt("Age", 8)

	fmt.Println(i.Value())   // -> {gopher 8}
	fmt.Println(i.Pointer()) // -> &{gopher 8}
}
*/
