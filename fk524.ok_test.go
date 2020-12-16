package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t F k 5 2 4
//  - - - - - - - - - -
//
//  Test Fk524 function.
//
//  Called:  Fk524, vvd
//
//  This revision:  2018 December 6
//
func TestFk524(t *testing.T) {
	const fname = "Fk524"
	var r2000, d2000, dr2000, dd2000, p2000, v2000,
		r1950, d1950, dr1950, dd1950, p1950, v1950 float64

	r2000 = 0.8723503576487275595
	d2000 = -0.7517076365138887672
	dr2000 = 0.2019447755430472323e-4
	dd2000 = 0.3541563940505160433e-5
	p2000 = 0.1559
	v2000 = 86.87

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoFk524},
		{"go", GoFk524},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		r1950, d1950, dr1950, dd1950, p1950, v1950 = test.fn(
			r2000, d2000, dr2000, dd2000, p2000, v2000)

		vvd(t, r1950, 0.8636359659799603487, 1e-13,
			tname, "r1950")
		vvd(t, d1950, -0.7550281733160843059, 1e-13,
			tname, "d1950")
		vvd(t, dr1950, 0.2023628192747172486e-4, 1e-17,
			tname, "dr1950")
		vvd(t, dd1950, 0.3624459754935334718e-5, 1e-18,
			tname, "dd1950")
		vvd(t, p1950, 0.1560079963299390241, 1e-13,
			tname, "p1950")
		vvd(t, v1950, 86.79606353469163751, 1e-11,
			tname, "v1950")
	}
}

func BenchmarkFk524(b *testing.B) {
	var r2000, d2000, dr2000, dd2000, p2000, v2000 float64

	r2000 = 0.8723503576487275595
	d2000 = -0.7517076365138887672
	dr2000 = 0.2019447755430472323e-4
	dd2000 = 0.3541563940505160433e-5
	p2000 = 0.1559
	v2000 = 86.87

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			c1, c2, c3, c4, c5, c6 float64)
	}{
		{"cgo", CgoFk524},
		{"go", GoFk524},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(r2000, d2000, dr2000, dd2000,
					p2000, v2000)
			}
		})
	}
}
