package server

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkHelloWorldHandler(b *testing.B) {
	s := &Server{}
	req := httptest.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		s.HelloWorldHandler(rr, req)
	}
}

func BenchmarkHelloWorldHandlerParallel(b *testing.B) {
	s := &Server{}
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			s.HelloWorldHandler(rr, req)
		}
	})
}
