package errid

import (
	"fmt"
	"reflect"
)

const ERR_NIL_DESC string = ""
const ERR_NIL_ID uint32 = 0

type Err struct {
	id   uint32
	desc string
}

func (e *Err) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("[%v]:%v", e.id, e.desc)
}

func (e *Err) Id() uint32 {
	if e == nil {
		return ERR_NIL_ID
	}
	return e.id
}

func (e *Err) Desc() string {
	if e == nil {
		return ERR_NIL_DESC
	}
	return e.desc
}

func NewMsg(id uint32, desc string) *Err {
	return &Err{id: id, desc: desc}
}

func GetId(e error) uint64 {
	if e == nil || reflect.ValueOf(e).Kind() != reflect.Ptr {
		return 0
	}
	v := reflect.ValueOf(e).Elem()
	if field := v.FieldByName("id"); field.IsValid() {
		back := field.Uint()
		return back
	}
	return 0
}
