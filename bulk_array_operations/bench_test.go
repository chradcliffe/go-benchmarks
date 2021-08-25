package bulk_array_operations_test

import (
	"fmt"
	"math/rand"
	"time"
	"testing"
)

var result []byte

func init() {
	rand.Seed(time.Now().Unix())
}

func writeByAppend(input []byte) []byte {
	numElem := len(input)
	output := make([]byte, 0, numElem)
	for _, v := range input {
		output = append(output, v)
	}
	return output
}

func writeByIndex(input []byte) []byte {
	numElem := len(input)
	output := make([]byte, numElem)
	for i := 0; i < numElem; i++ {
		output[i] = input[i]	
	}
	return output
}

func writeByAppendNoPreallocation(input []byte) []byte {
	output := []byte{}
	for _, v := range input {
		output = append(output, v)
	}
	return output
}

func writeByAppendEstimatedSize(numElem int, input []byte) []byte {
	output := make([]byte, 0, numElem)
	for _, v := range input {
		output = append(output, v)
	}
	return output
}

func writeByIndexEstimatedSize(numElem int, input []byte) []byte {
	output := make([]byte, numElem)
	i := 0
	for ; i < numElem; i++ {
		output[i] = input[i]
	}

	actualLen := len(input)
	for ; i < actualLen; i++ {
		output = append(output, input[i])
	}
	return output
}

func benchmarkWriteByAppend(numElem int, b *testing.B) {
	input := make([]byte, numElem)
	var output []byte

	_, err := rand.Read(input)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
                output = writeByAppend(input)
        }
	result = output
}

func BenchmarkWriteByAppend_100(b *testing.B) { benchmarkWriteByAppend(100, b) }
func BenchmarkWriteByAppend_1000(b *testing.B) { benchmarkWriteByAppend(1000, b) }
func BenchmarkWriteByAppend_10000(b *testing.B) { benchmarkWriteByAppend(10000, b) }
func BenchmarkWriteByAppend_100000(b *testing.B) { benchmarkWriteByAppend(100000, b) }
func BenchmarkWriteByAppend_1000000(b *testing.B) { benchmarkWriteByAppend(1000000, b) }

func benchmarkWriteByAppendNoPreallocation(numElem int, b *testing.B) {
	input := make([]byte, numElem)
	var output []byte

	_, err := rand.Read(input)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
                output = writeByAppendNoPreallocation(input)
        }
	result = output
}

func BenchmarkWriteByAppendNoPreallocation_100(b *testing.B) { benchmarkWriteByAppendNoPreallocation(100, b) }
func BenchmarkWriteByAppendNoPreallocation_1000(b *testing.B) { benchmarkWriteByAppendNoPreallocation(1000, b) }
func BenchmarkWriteByAppendNoPreallocation_10000(b *testing.B) { benchmarkWriteByAppendNoPreallocation(10000, b) }
func BenchmarkWriteByAppendNoPreallocation_100000(b *testing.B) { benchmarkWriteByAppendNoPreallocation(100000, b) }
func BenchmarkWriteByAppendNoPreallocation_1000000(b *testing.B) { benchmarkWriteByAppendNoPreallocation(1000000, b) }

func benchmarkWriteByIndex(numElem int, b *testing.B) {
	input := make([]byte, numElem)
	var output []byte

	_, err := rand.Read(input)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
                output = writeByIndex(input)
        }
	result = output
}

func BenchmarkWriteByIndex_100(b *testing.B) { benchmarkWriteByIndex(100, b) }
func BenchmarkWriteByIndex_1000(b *testing.B) { benchmarkWriteByIndex(1000, b) }
func BenchmarkWriteByIndex_10000(b *testing.B) { benchmarkWriteByIndex(10000, b) }
func BenchmarkWriteByIndex_100000(b *testing.B) { benchmarkWriteByIndex(100000, b) }
func BenchmarkWriteByIndex_1000000(b *testing.B) { benchmarkWriteByIndex(1000000, b) }

func benchmarkWriteByIndexEstimatedSize(numElem int, b *testing.B) {
	input := make([]byte, numElem)
	var output []byte

	_, err := rand.Read(input)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
                output = writeByIndexEstimatedSize(2*numElem/3, input)
        }
	result = output
}

func BenchmarkWriteByIndexEstimatedSize_100(b *testing.B) { benchmarkWriteByIndexEstimatedSize(100, b) }
func BenchmarkWriteByIndexEstimatedSize_1000(b *testing.B) { benchmarkWriteByIndexEstimatedSize(1000, b) }
func BenchmarkWriteByIndexEstimatedSize_10000(b *testing.B) { benchmarkWriteByIndexEstimatedSize(10000, b) }
func BenchmarkWriteByIndexEstimatedSize_100000(b *testing.B) { benchmarkWriteByIndexEstimatedSize(100000, b) }
func BenchmarkWriteByIndexEstimatedSize_1000000(b *testing.B) { benchmarkWriteByIndexEstimatedSize(1000000, b) }

func benchmarkWriteByAppendEstimatedSize(numElem int, b *testing.B) {
	input := make([]byte, numElem)
	var output []byte

	_, err := rand.Read(input)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
                output = writeByAppendEstimatedSize(2*numElem/3, input)
        }
	result = output
}

func BenchmarkWriteByAppendEstimatedSize_100(b *testing.B) { benchmarkWriteByAppendEstimatedSize(100, b) }
func BenchmarkWriteByAppendEstimatedSize_1000(b *testing.B) { benchmarkWriteByAppendEstimatedSize(1000, b) }
func BenchmarkWriteByAppendEstimatedSize_10000(b *testing.B) { benchmarkWriteByAppendEstimatedSize(10000, b) }
func BenchmarkWriteByAppendEstimatedSize_100000(b *testing.B) { benchmarkWriteByAppendEstimatedSize(100000, b) }
func BenchmarkWriteByAppendEstimatedSize_1000000(b *testing.B) { benchmarkWriteByAppendEstimatedSize(1000000, b) }


