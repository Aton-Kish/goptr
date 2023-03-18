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

func Pointer[T any](v T) (p *T) {
	return &v
}

func ToValue[T any](p *T) (v T) {
	if p == nil {
		return v
	}

	return *p
}

func PointerMap[K comparable, V any](vm map[K]V) (pm map[K]*V) {
	pm = make(map[K]*V, len(vm))

	for k, v := range vm {
		v := v
		pm[k] = Pointer(v)
	}

	return pm
}

func ToValueMap[K comparable, V any](pm map[K]*V) (vm map[K]V) {
	vm = make(map[K]V, len(pm))

	for k, p := range pm {
		p := p
		vm[k] = ToValue(p)
	}

	return vm
}

func PointerSlice[E any](vs []E) (ps []*E) {
	ps = make([]*E, len(vs))

	for i, v := range vs {
		v := v
		ps[i] = Pointer(v)
	}

	return ps
}

func ToValueSlice[E any](ps []*E) (vs []E) {
	vs = make([]E, len(ps))

	for i, p := range ps {
		p := p
		vs[i] = ToValue(p)
	}

	return vs
}
