package bufrw

import (
	"bytes"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func TestBufferReadWriteBool(t *testing.T) {
	tests := []bool{
		true,
		false,
	}
	for _, value := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteBool(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadBool(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
}

func TestBufferReadWriteBools(t *testing.T) {
	tests := [][]bool{
		{},
		{false},
		{false, true},
	}
	for _, values := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteBools(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadBools(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		values := make([]bool, n)
		for i := 0; i < n; i++ {
			values[i] = rand.Int()%2 == 0
		}
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteBools(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadBools(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
}

func TestBufferReadWriteByteValue(t *testing.T) {
	tests := make([]byte, 255)
	for i := 0; i < 255; i++ {
		tests[i] = byte(i)
	}
	for _, value := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteByteValue(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadByteValue(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
}

func TestBufferReadWriteByteValues(t *testing.T) {
	tests := [][]byte{
		{},
		{0, 1, 254, 255, 255},
	}
	for _, values := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteByteValues(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadByteValues(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		values := make([]byte, n)
		for i := 0; i < n; i++ {
			values[i] = byte(rand.Intn(256))
		}
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteByteValues(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadByteValues(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
}

func TestBufferReadWriteInt(t *testing.T) {
	tests := []int{
		math.MinInt32,
		-9999,
		-1,
		0,
		1,
		9999,
		math.MaxInt32,
	}
	for _, value := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteInt(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInt(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
	for i := 0; i < 1000; i++ {
		var buf Buffer
		var w bytes.Buffer
		value := int(rand.Int31())
		if err := buf.WriteInt(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInt(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
}

func TestBufferReadWriteInts(t *testing.T) {
	tests := [][]int{
		{},
		{math.MinInt32, -9999, -1, 0, 1, 9999, math.MaxInt32},
	}
	for _, values := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteInts(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInts(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		values := make([]int, n)
		for i := 0; i < n; i++ {
			values[i] = int(rand.Int31())
		}
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteInts(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInts(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
}

func TestBufferReadWriteInt64(t *testing.T) {
	tests := []int64{
		math.MinInt64,
		-9999,
		-1,
		0,
		1,
		9999,
		math.MaxInt64,
	}
	for _, value := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteInt64(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInt64(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
	for i := 0; i < 1000; i++ {
		var buf Buffer
		var w bytes.Buffer
		value := rand.Int63()
		if err := buf.WriteInt64(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInt64(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
}

func TestBufferReadWriteInt64s(t *testing.T) {
	tests := [][]int64{
		{},
		{math.MinInt64, -9999, -1, 0, 1, 9999, math.MaxInt64},
	}
	for _, values := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteInt64s(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInt64s(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		values := make([]int64, n)
		for i := 0; i < n; i++ {
			values[i] = rand.Int63()
		}
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteInt64s(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadInt64s(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
}

func TestBufferReadWriteFloat64(t *testing.T) {
	tests := []float64{
		math.MaxInt64,
		-9999,
		-1.5,
		-1,
		-0.01,
		0,
		0.01,
		1,
		1.5,
		9999,
		math.MaxInt64,
	}
	for _, value := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteFloat64(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadFloat64(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
	for i := 0; i < 1000; i++ {
		var buf Buffer
		var w bytes.Buffer
		value := rand.Float64()
		if err := buf.WriteFloat64(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadFloat64(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
}

func TestBufferReadWriteFloat64s(t *testing.T) {
	tests := [][]float64{
		{},
		{math.MaxInt64, -9999, -1.5, -1, -0.01, 0, 0.01, 1, 1.5, 9999, math.MaxInt64},
	}
	for _, values := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteFloat64s(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadFloat64s(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		values := make([]float64, n)
		for i := 0; i < n; i++ {
			values[i] = rand.Float64()
		}
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteFloat64s(&w, values...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadFloat64s(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, values) {
			t.Errorf("Write/Read %v = %v", values, got)
		}
	}
}

func TestBufferReadWriteString(t *testing.T) {
	tests := []string{
		"",
		"A",
		"ㄒ乇丂ㄒ",
	}
	for _, value := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteString(&w, value); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadString(&w)
		if err != nil {
			t.Fatal(err)
		}
		if got != value {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
}

func TestBufferReadWriteStrings(t *testing.T) {
	tests := [][]string{
		{},
		{""},
		{"", "A"},
		{"ㄒ乇", "丂ㄒ"},
	}
	for _, value := range tests {
		var buf Buffer
		var w bytes.Buffer
		if err := buf.WriteStrings(&w, value...); err != nil {
			t.Fatal(err)
		}
		got, err := buf.ReadStrings(&w)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, value) {
			t.Errorf("Write/Read %v = %v", value, got)
		}
	}
}
