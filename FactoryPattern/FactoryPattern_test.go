package FactoryPattern

import (
	"testing"
	"fmt"
	"reflect"
)

func TestFactoryPattern(t *testing.T)  {
	factory := ShapeFactory{}
	//var shape Shape
	shape := factory.GetShape(CircleName)
	if c,ok := shape.(Circle);!ok {
		fmt.Println(reflect.TypeOf(c))
		t.Errorf("Input is %s but output is not.Type is %v", CircleName,reflect.TypeOf(c))
	}
	shape = factory.GetShape(RectangleName)
	if r,ok := shape.(Circle);!ok {
		fmt.Println(r)
		t.Errorf("Input is %s but output is not", RectangleName)
	}
}
