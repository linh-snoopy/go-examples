package main

import (
	"testing"
)

func TestSumS(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	expected := 21
	actual := Sum(numbers)
	if actual != expected {
		t.Errorf("Get %d instead of %d", actual, expected)
	}
}

//func TestSumF(t *testing.T) {
//	numbers := []int{1,2,3,4,5,6}
//	expected := 12
//	actual := Sum(numbers)
//	if actual != expected {
//		t.Errorf("Get %d instead of %d", actual, expected)
//	}
//}

func benchmarkSum(numbers []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(numbers)
	}
}

func BenchmarkSum1(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 23, 3, 5, 5, 7, 87, 5, 2}
	benchmarkSum(numbers, b)
}

func BenchmarkSum2(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 5, 7, 87, 5, 2}
	benchmarkSum(numbers, b)
}

func BenchmarkSum3(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 23, 3, 5, 5, 7, 87, 5, 2, 4, 45, 7, 9, 2, 2, 5, 7, 8, 3, 21, 4, 4, 76, 2, 45, 7, 8, 3, 4}
	benchmarkSum(numbers, b)
}

func BenchmarkSum34(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 23, 3, 5, 5, 7, 87, 5, 9, 2, 2, 5, 7, 8, 3, 21, 4, 4, 76, 2, 45, 7, 8, 3, 4}
	benchmarkSum(numbers, b)
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
