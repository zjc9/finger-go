package main

import (
	"reflect"
)

type Init struct {
	fingersMap map[string]reflect.Type
}

func (i *Init) init() {
	//var impMap = make(map[string]reflect.Type)
	//f := reflect.TypeOf(fingers.WebAnalyzer{})
	//f
}
