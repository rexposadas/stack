package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyInt int

var i int

// Test working with basic maps.
func Test_typeConversion(t *testing.T) {
	mi := MyInt(12)

	// This is illegal because, even if they have the same underlying type, they are distinct.
	// i = mi

	// We need a type conversion.
	i = int(mi)
	assert.Equal(t, 12, i)

}

func Test_assertions(t *testing.T) {
	var w io.Writer = os.Stdout

	// value.(T)
	// "Take the value and make it this type"
	f := w.(*os.File) // works!
	assert.NotNil(t, f)

	// Will fail because os.Stdout is not a buffer
	// c:= w.(*bytes.Buffer)  // panics

	// To avoid the panic, you can use the second form:
	ff, ok := w.(*os.File)
	assert.True(t, ok)
	assert.NotNil(t, ff)

	b, ok := w.(*bytes.Buffer)
	assert.False(t, ok)
	assert.Nil(t, b)
}

type A struct {
	a string
	b []int
}

// Use the reflect package to interrogate types.
func Test_deepEqual(t *testing.T) {
	a := A{
		a: "a",
		b: []int{1, 2, 3},
	}

	b := A{
		a: "a",
		b: []int{1, 2, 3},
	}

	c := A{
		a: "a",
		b: []int{8}, // c is different from and b because of this value.
	}

	assert.True(t, reflect.DeepEqual(a, b))
	assert.False(t, reflect.DeepEqual(a, c))
}
