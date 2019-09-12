package httptest

import "testing"

func BenchmarkFasthttpGet(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		FasthttpGet()
	}

}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Get()
	}
}
