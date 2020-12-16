package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t N u m 0 0 b
//  - - - - - - - - - - -
//
//  Test Num00b function.
//
//  Called:  Num00b, vvd
//
//  This revision:  2013 August 7
//
func TestNum00b(t *testing.T) {
	const fname = "Num00b"
	var rmatn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNum00b},
		{"go", GoNum00b},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rmatn = test.fn(2400000.5, 53736)

		vvd(t, rmatn[0][0], 0.9999999999536069682, 1e-12,
			tname, "11")
		vvd(t, rmatn[0][1], 0.8837746144871248011e-5, 1e-12,
			tname, "12")
		vvd(t, rmatn[0][2], 0.3831488838252202945e-5, 1e-12,
			tname, "13")

		vvd(t, rmatn[1][0], -0.8837590456632304720e-5, 1e-12,
			tname, "21")
		vvd(t, rmatn[1][1], 0.9999999991354692733, 1e-12,
			tname, "22")
		vvd(t, rmatn[1][2], -0.4063198798559591654e-4, 1e-12,
			tname, "23")

		vvd(t, rmatn[2][0], -0.3831847930134941271e-5, 1e-12,
			tname, "31")
		vvd(t, rmatn[2][1], 0.4063195412258168380e-4, 1e-12,
			tname, "32")
		vvd(t, rmatn[2][2], 0.9999999991671806225, 1e-12,
			tname, "33")
	}
}

func BenchmarkNum00b(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) [3][3]float64
	}{
		{"cgo", CgoNum00b},
		{"go", GoNum00b},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736)
			}
		})
	}
}
