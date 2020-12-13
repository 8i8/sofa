package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t D t d b
//  - - - - - - - - -
//
//  Test Dtdb function.
//
//  Called:  Dtdb, vvd
//
//  This revision:  2013 August 7
//
func TestDtdb(t *testing.T) {
	const fname = "Dtdb"
	var dtdb float64

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) float64
	}{
		{"cgo", CgoDtdb},
		{"go", GoDtdb},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dtdb = test.fn(2448939.5, 0.123, 0.76543, 5.0123,
			5525.242, 3190.0)

		vvd(t, dtdb, -0.1280368005936998991e-2, 1e-15, tname,
			"")
	}
}

func BenchmarkDtdb(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) float64
	}{
		{"cgo", CgoDtdb},
		{"go", GoDtdb},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2448939.5, 0.123, 0.76543,
					5.0123, 5525.242, 3190.0)
			}
		})
	}
}
