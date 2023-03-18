// Copyright (c) 2023 Aton-Kish
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stringZero  string
	stringHello string = "hello"

	int0 int
	int1 int = 1

	float64_0_0 float64
	float64_1_0 float64 = 1.0
)

func TestPointerString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		pointer *string
	}{
		{name: "happy path: (empty)", value: stringZero, pointer: &stringZero},
		{name: "happy path: hello", value: stringHello, pointer: &stringHello},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Pointer(tt.value)
			assert.Equal(t, tt.pointer, actual)
		})
	}
}

func TestPointerInt(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		pointer *int
	}{
		{name: "happy path: 0", value: int0, pointer: &int0},
		{name: "happy path: 1", value: int1, pointer: &int1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Pointer(tt.value)
			assert.Equal(t, tt.pointer, actual)
		})
	}
}

func TestPointerFloat64(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		pointer *float64
	}{
		{name: "happy path: 0.0", value: float64_0_0, pointer: &float64_0_0},
		{name: "happy path: 1.0", value: float64_1_0, pointer: &float64_1_0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Pointer(tt.value)
			assert.Equal(t, tt.pointer, actual)
		})
	}
}

func TestToValueString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		pointer *string
	}{
		{name: "happy path: nil", value: stringZero, pointer: nil},
		{name: "happy path: (empty)", value: stringZero, pointer: &stringZero},
		{name: "happy path: hello", value: stringHello, pointer: &stringHello},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValue(tt.pointer)
			assert.Equal(t, tt.value, actual)
		})
	}
}

func TestToValueInt(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		pointer *int
	}{
		{name: "happy path: nil", value: int0, pointer: nil},
		{name: "happy path: 0", value: int0, pointer: &int0},
		{name: "happy path: 1", value: int1, pointer: &int1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValue(tt.pointer)
			assert.Equal(t, tt.value, actual)
		})
	}
}

func TestToValueFloat64(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		pointer *float64
	}{
		{name: "happy path: nil", value: float64_0_0, pointer: nil},
		{name: "happy path: 0.0", value: float64_0_0, pointer: &float64_0_0},
		{name: "happy path: 1.0", value: float64_1_0, pointer: &float64_1_0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValue(tt.pointer)
			assert.Equal(t, tt.value, actual)
		})
	}
}

func TestPointerMapString(t *testing.T) {
	tests := []struct {
		name       string
		valueMap   map[string]string
		pointerMap map[string]*string
	}{
		{
			name: "happy path",
			valueMap: map[string]string{
				"zero":  stringZero,
				"hello": stringHello,
			},
			pointerMap: map[string]*string{
				"zero":  &stringZero,
				"hello": &stringHello,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PointerMap(tt.valueMap)
			assert.Equal(t, tt.pointerMap, actual)
		})
	}
}

func TestPointerMapInt(t *testing.T) {
	tests := []struct {
		name       string
		valueMap   map[string]int
		pointerMap map[string]*int
	}{
		{
			name: "happy path",
			valueMap: map[string]int{
				"0": int0,
				"1": int1,
			},
			pointerMap: map[string]*int{
				"0": &int0,
				"1": &int1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PointerMap(tt.valueMap)
			assert.Equal(t, tt.pointerMap, actual)
		})
	}
}

func TestPointerMapFloat64(t *testing.T) {
	tests := []struct {
		name       string
		valueMap   map[string]float64
		pointerMap map[string]*float64
	}{
		{
			name: "happy path",
			valueMap: map[string]float64{
				"0.0": float64_0_0,
				"1.0": float64_1_0,
			},
			pointerMap: map[string]*float64{
				"0.0": &float64_0_0,
				"1.0": &float64_1_0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PointerMap(tt.valueMap)
			assert.Equal(t, tt.pointerMap, actual)
		})
	}
}

func TestToValueMapString(t *testing.T) {
	tests := []struct {
		name       string
		valueMap   map[string]string
		pointerMap map[string]*string
	}{
		{
			name: "happy path: including nil",
			valueMap: map[string]string{
				"nil":   stringZero,
				"zero":  stringZero,
				"hello": stringHello,
			},
			pointerMap: map[string]*string{
				"nil":   nil,
				"zero":  &stringZero,
				"hello": &stringHello,
			},
		},
		{
			name: "happy path: not including nil",
			valueMap: map[string]string{
				"zero":  stringZero,
				"hello": stringHello,
			},
			pointerMap: map[string]*string{
				"zero":  &stringZero,
				"hello": &stringHello,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValueMap(tt.pointerMap)
			assert.Equal(t, tt.valueMap, actual)
		})
	}
}

func TestToValueMapInt(t *testing.T) {
	tests := []struct {
		name       string
		valueMap   map[string]int
		pointerMap map[string]*int
	}{
		{
			name: "happy path: including nil",
			valueMap: map[string]int{
				"nil": int0,
				"0":   int0,
				"1":   int1,
			},
			pointerMap: map[string]*int{
				"nil": nil,
				"0":   &int0,
				"1":   &int1,
			},
		},
		{
			name: "happy path: not including nil",
			valueMap: map[string]int{
				"0": int0,
				"1": int1,
			},
			pointerMap: map[string]*int{
				"0": &int0,
				"1": &int1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValueMap(tt.pointerMap)
			assert.Equal(t, tt.valueMap, actual)
		})
	}
}

func TestToValueMapFloat64(t *testing.T) {
	tests := []struct {
		name       string
		valueMap   map[string]float64
		pointerMap map[string]*float64
	}{
		{
			name: "happy path: including nil",
			valueMap: map[string]float64{
				"nil": float64_0_0,
				"0.0": float64_0_0,
				"1.0": float64_1_0,
			},
			pointerMap: map[string]*float64{
				"nil": nil,
				"0.0": &float64_0_0,
				"1.0": &float64_1_0,
			},
		},
		{
			name: "happy path: not including nil",
			valueMap: map[string]float64{
				"0.0": float64_0_0,
				"1.0": float64_1_0,
			},
			pointerMap: map[string]*float64{
				"0.0": &float64_0_0,
				"1.0": &float64_1_0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValueMap(tt.pointerMap)
			assert.Equal(t, tt.valueMap, actual)
		})
	}
}

func TestPointerSliceString(t *testing.T) {
	tests := []struct {
		name         string
		valueSlice   []string
		pointerSlice []*string
	}{
		{
			name: "happy path",
			valueSlice: []string{
				stringZero,
				stringHello,
			},
			pointerSlice: []*string{
				&stringZero,
				&stringHello,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PointerSlice(tt.valueSlice)
			assert.Equal(t, tt.pointerSlice, actual)
		})
	}
}

func TestPointerSliceInt(t *testing.T) {
	tests := []struct {
		name         string
		valueSlice   []int
		pointerSlice []*int
	}{
		{
			name: "happy path",
			valueSlice: []int{
				int0,
				int1,
			},
			pointerSlice: []*int{
				&int0,
				&int1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PointerSlice(tt.valueSlice)
			assert.Equal(t, tt.pointerSlice, actual)
		})
	}
}

func TestPointerSliceFloat64(t *testing.T) {
	tests := []struct {
		name         string
		valueSlice   []float64
		pointerSlice []*float64
	}{
		{
			name: "happy path",
			valueSlice: []float64{
				float64_0_0,
				float64_1_0,
			},
			pointerSlice: []*float64{
				&float64_0_0,
				&float64_1_0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PointerSlice(tt.valueSlice)
			assert.Equal(t, tt.pointerSlice, actual)
		})
	}
}

func TestToValueSliceString(t *testing.T) {
	tests := []struct {
		name         string
		valueSlice   []string
		pointerSlice []*string
	}{
		{
			name: "happy path: including nil",
			valueSlice: []string{
				stringZero,
				stringZero,
				stringHello,
			},
			pointerSlice: []*string{
				nil,
				&stringZero,
				&stringHello,
			},
		},
		{
			name: "happy path: not including nil",
			valueSlice: []string{
				stringZero,
				stringHello,
			},
			pointerSlice: []*string{
				&stringZero,
				&stringHello,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValueSlice(tt.pointerSlice)
			assert.Equal(t, tt.valueSlice, actual)
		})
	}
}

func TestToValueSliceInt(t *testing.T) {
	tests := []struct {
		name         string
		valueSlice   []int
		pointerSlice []*int
	}{
		{
			name: "happy path: including nil",
			valueSlice: []int{
				int0,
				int0,
				int1,
			},
			pointerSlice: []*int{
				nil,
				&int0,
				&int1,
			},
		},
		{
			name: "happy path: not including nil",
			valueSlice: []int{
				int0,
				int1,
			},
			pointerSlice: []*int{
				&int0,
				&int1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValueSlice(tt.pointerSlice)
			assert.Equal(t, tt.valueSlice, actual)
		})
	}
}

func TestToValueSliceFloat64(t *testing.T) {
	tests := []struct {
		name         string
		valueSlice   []float64
		pointerSlice []*float64
	}{
		{
			name: "happy path: including nil",
			valueSlice: []float64{
				float64_0_0,
				float64_0_0,
				float64_1_0,
			},
			pointerSlice: []*float64{
				nil,
				&float64_0_0,
				&float64_1_0,
			},
		},
		{
			name: "happy path: not including nil",
			valueSlice: []float64{
				float64_0_0,
				float64_1_0,
			},
			pointerSlice: []*float64{
				&float64_0_0,
				&float64_1_0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToValueSlice(tt.pointerSlice)
			assert.Equal(t, tt.valueSlice, actual)
		})
	}
}
