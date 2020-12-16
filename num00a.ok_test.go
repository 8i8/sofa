package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t N u m 0 0 a
//  - - - - - - - - - - -
//
//  Test Num00a function.
//
//  Called:  Num00a, vvd
//
//  This revision:  2013 August 7
//
func TestNum00a(t *testing.T) {
	const fname = "Num00a"
	var rmatn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNum00a},
		{"go", GoNum00a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rmatn = test.fn(2400000.5, 53736.0)

		vvd(t, rmatn[0][0], 0.9999999999536227949, 1e-12,
			tname, "11")
		vvd(t, rmatn[0][1], 0.8836238544090873336e-5, 1e-12,
			tname, "12")
		vvd(t, rmatn[0][2], 0.3830835237722400669e-5, 1e-12,
			tname, "13")

		vvd(t, rmatn[1][0], -0.8836082880798569274e-5, 1e-12,
			tname, "21")
		vvd(t, rmatn[1][1], 0.9999999991354655028, 1e-12,
			tname, "22")
		vvd(t, rmatn[1][2], -0.4063240865362499850e-4, 1e-12,
			tname, "23")

		vvd(t, rmatn[2][0], -0.3831194272065995866e-5, 1e-12,
			tname, "31")
		vvd(t, rmatn[2][1], 0.4063237480216291775e-4, 1e-12,
			tname, "32")
		vvd(t, rmatn[2][2], 0.9999999991671660338, 1e-12,
			tname, "33")
	}
}

func BenchmarkNum00a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNum00a},
		{"go", GoNum00a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0)
			}
		})
	}
}
