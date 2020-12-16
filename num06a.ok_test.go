package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t N u m 0 6 a
//  - - - - - - - - - - -
//
//  Test Num06a function.
//
//  Called:  Num06a, vvd
//
//  This revision:  2013 August 7
//
func TestNum06a(t *testing.T) {
	const fname = "Num06a"
	var rmatn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNum06a},
		{"go", GoNum06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rmatn = test.fn(2400000.5, 53736)

		vvd(t, rmatn[0][0], 0.9999999999536227668, 1e-12,
			tname, "11")
		vvd(t, rmatn[0][1], 0.8836241998111535233e-5, 1e-12,
			tname, "12")
		vvd(t, rmatn[0][2], 0.3830834608415287707e-5, 1e-12,
			tname, "13")

		vvd(t, rmatn[1][0], -0.8836086334870740138e-5, 1e-12,
			tname, "21")
		vvd(t, rmatn[1][1], 0.9999999991354657474, 1e-12,
			tname, "22")
		vvd(t, rmatn[1][2], -0.4063240188248455065e-4, 1e-12,
			tname, "23")

		vvd(t, rmatn[2][0], -0.3831193642839398128e-5, 1e-12,
			tname, "31")
		vvd(t, rmatn[2][1], 0.4063236803101479770e-4, 1e-12,
			tname, "32")
		vvd(t, rmatn[2][2], 0.9999999991671663114, 1e-12,
			tname, "33")
	}
}

func BenchmarkNum06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNum06a},
		{"go", GoNum06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736)
			}
		})
	}
}
