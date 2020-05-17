package main

import "testing"

func BenchmarkAddItemUseAppend(b *testing.B) {
	list := make([]*Point, b.N, b.N)
	for n := 0; n < b.N; n++ {
		AddItemUseAppend(&list)
	}
}

func BenchmarkAddItemUseIndex(b *testing.B) {
	list := make([]*Point, b.N, b.N)
	for n := 0; n < b.N; n++ {
		AddItemUseIndex(&list, n)
	}
}
