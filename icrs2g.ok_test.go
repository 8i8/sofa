package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t I c r s 2 g
//  - - - - - - - - - - -
//
//  Test Icrs2g function.
//
//  Called:  Icrs2g, vvd
//
//  This revision:  2015 January 30
//
func TestIcrs2g(t *testing.T) {
	const fname = "Icrs2g"
	var dr, dd, dl, db float64

	dr = 5.9338074302227188048671087
	dd = -1.1784870613579944551540570

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoIcrs2g},
		{"go", GoIcrs2g},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		dl, db = test.fn(dr, dd)

		vvd(t, dl, 5.5850536063818546461558, 1e-14, tname, "L")
		vvd(t, db, -0.7853981633974483096157, 1e-14, tname, "B")
	}
}

func BenchmarkIcrs2g(b *testing.B) {
	var dr, dd float64

	dr = 5.9338074302227188048671087
	dd = -1.1784870613579944551540570

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2 float64)
	}{
		{"cgo", CgoIcrs2g},
		{"go", GoIcrs2g},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(dr, dd)
			}
		})
	}
}
