package structs

import "testing"

func TestPerimter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// Table Driven Test
	// Kalo misal cara pengecekannya sama, jadiin struct aja trus for loop in

	// dalam context ini karna sama sama shape dan kita mau check apakah areanya udh bener
	// Hal ini akan mempermudah untuk testing interface, atau sesuatu yang memiliki kegunaan yang sama
	// Jadi misalkan ada yang baru implement "interface Shape", hanya masukkan ke shapes
	shapes := [...]struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range shapes {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
