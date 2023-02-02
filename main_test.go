package main

import (
	protobuf "benchmark/gen/proto"
	"math/rand"
	"testing"

	"google.golang.org/protobuf/proto"
)

func BenchmarkEmpty(b *testing.B) {
	dataEmpty := &protobuf.Container{
		Id:      1,
		OtherId: new(int32),
		Payload: []*protobuf.Payload{},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(dataEmpty)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

func Benchmark10kb(b *testing.B) {
	bytes := make([]byte, 10240)
	_, _ = rand.Read(bytes)
	data := &protobuf.Container{
		Id:      1,
		OtherId: new(int32),
		Payload: []*protobuf.Payload{{
			Data: bytes,
		}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(data)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

func Benchmark100kb(b *testing.B) {
	bytes := make([]byte, 102400)
	_, _ = rand.Read(bytes)
	data := &protobuf.Container{
		Id:      1,
		OtherId: new(int32),
		Payload: []*protobuf.Payload{{
			Data: bytes,
		}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(data)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

func Benchmark1000kb(b *testing.B) {
	bytes := make([]byte, 1024000)
	_, _ = rand.Read(bytes)
	data := &protobuf.Container{
		Id:      1,
		OtherId: new(int32),
		Payload: []*protobuf.Payload{{
			Data: bytes,
		}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(data)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

func Benchmark10x10kb(b *testing.B) {
	data := &protobuf.Container{
		Id:      1,
		OtherId: new(int32),
		Payload: []*protobuf.Payload{},
	}
	for i := 0; i < 10; i++ {
		bytes := make([]byte, 10240)
		_, _ = rand.Read(bytes)
		data.Payload = append(data.Payload, &protobuf.Payload{Data: bytes})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(data)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

func Benchmark100x10kb(b *testing.B) {
	data := &protobuf.Container{
		Id:      1,
		OtherId: new(int32),
		Payload: []*protobuf.Payload{},
	}
	for i := 0; i < 100; i++ {
		bytes := make([]byte, 10240)
		_, _ = rand.Read(bytes)
		data.Payload = append(data.Payload, &protobuf.Payload{Data: bytes})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(data)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}
