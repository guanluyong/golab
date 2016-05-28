// Copyright 2016 By GuanLuyong.  All rights reserved.

// Package router is For web router init.
// RegisterUtil use reflect to simple working.
package router

import (
	//"fmt"
	"reflect"
)

// RegisterUtil used to init Router
func RegisterUtil(structRouter interface{}) {
	vr := reflect.ValueOf(structRouter)
	//tr := reflect.TypeOf(structRouter)
	for i := 0; i < vr.NumMethod(); i++ {
		vm := vr.Method(i)
		//	tm := tr.Method(i)
		//	fmt.Println(tm.Name)
		vm.Call([]reflect.Value{})
	}
}
