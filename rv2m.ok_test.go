package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t R v 2 m
//  - - - - - - - - -
//
//  Test Rv2m function.
//
//  Called:  Rv2m, vvd
//
//  This revision:  2013 August 7
//
func TestRv2m(t *testing.T) {
	const fname = "Rv2m"
	var r [3][3]float64
	var w [3]float64

	w[0] = 0.0
	w[1] = 1.41371669
	w[2] = -1.88495559

	tests := []struct {
		ref string
		fn  func([3]float64) [3][3]float64
	}{
		{"cgo", CgoRv2m},
		{"go", GoRv2m},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		r = test.fn(w)

		vvd(t, r[0][0], -0.7071067782221119905, 1e-14, tname, "11")
		vvd(t, r[0][1], -0.5656854276809129651, 1e-14, tname, "12")
		vvd(t, r[0][2], -0.4242640700104211225, 1e-14, tname, "13")

		vvd(t, r[1][0], 0.5656854276809129651, 1e-14, tname, "21")
		vvd(t, r[1][1], -0.0925483394532274246, 1e-14, tname, "22")
		vvd(t, r[1][2], -0.8194112531408833269, 1e-14, tname, "23")

		vvd(t, r[2][0], 0.4242640700104211225, 1e-14, tname, "31")
		vvd(t, r[2][1], -0.8194112531408833269, 1e-14, tname, "32")
		vvd(t, r[2][2], 0.3854415612311154341, 1e-14, tname, "33")
	}
}

func BenchmarkRv2m(b *testing.B) {
	var w [3]float64

	w[0] = 0.0
	w[1] = 1.41371669
	w[2] = -1.88495559

	tests := []struct {
		ref string
		fn  func([3]float64) [3][3]float64
	}{
		{"cgo", CgoRv2m},
		{"go", GoRv2m},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(w)
			}
		})
	}
}
