package sofa

import "testing"

//
//  - - - - - - -
//   T e s t T r
//  - - - - - - -
//
//  Test Tr function.
//
//  Called:  Tr, vvd
//
//  This revision:  2013 August 7
//
func TestTr(t *testing.T) {
	const fname = "Tr"
	var r [3][3]float64

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	tests := []struct {
		ref string
		fn  func([3][3]float64) [3][3]float64
	}{
		{"cgo", CgoTr},
		{"go", GoTr},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rt := test.fn(r)

		vvd(t, rt[0][0], 2.0, 0.0, tname, "11")
		vvd(t, rt[0][1], 3.0, 0.0, tname, "12")
		vvd(t, rt[0][2], 3.0, 0.0, tname, "13")

		vvd(t, rt[1][0], 3.0, 0.0, tname, "21")
		vvd(t, rt[1][1], 2.0, 0.0, tname, "22")
		vvd(t, rt[1][2], 4.0, 0.0, tname, "23")

		vvd(t, rt[2][0], 2.0, 0.0, tname, "31")
		vvd(t, rt[2][1], 3.0, 0.0, tname, "32")
		vvd(t, rt[2][2], 5.0, 0.0, tname, "33")
	}
}

func BenchmarkTr(b *testing.B) {
	var r [3][3]float64

	r[0][0] = 2.0
	r[0][1] = 3.0
	r[0][2] = 2.0

	r[1][0] = 3.0
	r[1][1] = 2.0
	r[1][2] = 3.0

	r[2][0] = 3.0
	r[2][1] = 4.0
	r[2][2] = 5.0

	tests := []struct {
		ref string
		fn  func([3][3]float64) [3][3]float64
	}{
		{"cgo", CgoTr},
		{"go", GoTr},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(r)
			}
		})
	}
}
