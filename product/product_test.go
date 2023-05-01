package product_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-generics/product"
)

func TestProduct_WithInts(t *testing.T) {
	t.Parallel()
	got := product.Product(2, 3)

	want := 6

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestProduct_WithFloats(t *testing.T) {
	t.Parallel()
	got := product.Product(2.2, 3.3)

	want := 7.26

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestProduct_WithComplex(t *testing.T) {
	t.Parallel()
	got := product.Product(2i, 3i)

	want := -6 + 0i

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
