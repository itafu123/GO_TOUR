package main

import "testing"

func TestMe1(t *testing.T){
	if(me(1, 2) != 3){
		t.Error("fail")
	}
}

// 寫錯的test
func TestMe2(t *testing.T){
	if(me(1, 2) != 4){
		t.Error("fail")
	}
}