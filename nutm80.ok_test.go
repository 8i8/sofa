package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t N u t m 8 0
//  - - - - - - - - - - -
//
//  Test Nutm80 function.
//
//  Called:  Nutm80, vvd
//
//  This revision:  2013 August 7
//
func TestNutm80(t *testing.T) {
	const fname = "Nutm80"
	var rmatn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNutm80},
		{"go", GoNutm80},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rmatn = test.fn(2400000.5, 53736.0)

		vvd(t, rmatn[0][0], 0.9999999999534999268, 1e-12,
			tname, "11")
		vvd(t, rmatn[0][1], 0.8847935789636432161e-5, 1e-12,
			tname, "12")
		vvd(t, rmatn[0][2], 0.3835906502164019142e-5, 1e-12,
			tname, "13")

		vvd(t, rmatn[1][0], -0.8847780042583435924e-5, 1e-12,
			tname, "21")
		vvd(t, rmatn[1][1], 0.9999999991366569963, 1e-12,
			tname, "22")
		vvd(t, rmatn[1][2], -0.4060052702727130809e-4, 1e-12,
			tname, "23")

		vvd(t, rmatn[2][0], -0.3836265729708478796e-5, 1e-12,
			tname, "31")
		vvd(t, rmatn[2][1], 0.4060049308612638555e-4, 1e-12,
			tname, "32")
		vvd(t, rmatn[2][2], 0.9999999991684415129, 1e-12,
			tname, "33")
	}
}

func BenchmarkNutm80(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNutm80},
		{"go", GoNutm80},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
