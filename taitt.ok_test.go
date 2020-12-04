package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t T a i t t
//  - - - - - - - - - -
//
//  Test Taitt function.
//
//  Called:  Taitt, vvd, viv
//
//  This revision:  2013 August 7
//
func TestTaitt(t *testing.T) {
	const fname = "Taitt"

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2 float64, b3 error)
	}{
		{"cgo", CgoTaitt},
		{"go", GoTaitt},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		t1, t2, err := test.fn(2453750.5, 0.892482639)

		vvd(t, t1, 2453750.5, 1e-6, tname, "t1")
		vvd(t, t2, 0.892855139, 1e-12, tname, "t2")
		errT(t, nil, err, tname)
	}
}

func BenchmarkTaitt(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (b1, b2 float64, l3 error)
	}{
		{"cgo", CgoTaitt},
		{"go", GoTaitt},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(2453750.5, 0.892482639)
			}
		})
	}
}
