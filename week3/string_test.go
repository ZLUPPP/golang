package main

import "testing"

var resultString string
var resultBytes []byte

func BenchmarkWrongProcessPayload(b *testing.B) {
	data := []byte("   hello\rworld\rtest   ")

	b.ReportAllocs()

	for b.Loop() {
		resultString = wrongProcessPayload(data)
	}
}

func BenchmarkCorrectProcessPayload(b *testing.B) {
	data := []byte("   hello\rworld\rtest   ")

	b.ReportAllocs()

	for b.Loop() {
		resultBytes = correctProcessPayload(data)
	}
}
