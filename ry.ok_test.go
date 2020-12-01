package sofa

import "testing"

//
//  - - - - - - -
//   T e s t R y
//  - - - - - - -
//
//  Test Ry function.
//
//  Called:  Ry, vvd
//
//  This revision:  2013 August 7
//
func TestRy(t *testing.T) {
	const fname = "Ry"
	var theta float64
	var r [3][3]float64

	theta = 0.3456789

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
		fn  func(float64, [3][3]float64) [3][3]float64
	}{
		{"cgo", CgoRy},
		{"go", GoRy},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rota := test.fn(theta, r)

		vvd(t, rota[0][0], 0.8651847818978159930, 1e-12, tname, "11")
		vvd(t, rota[0][1], 1.467194920539316554, 1e-12, tname, "12")
		vvd(t, rota[0][2], 0.1875137911274457342, 1e-12, tname, "13")

		vvd(t, rota[1][0], 3, 1e-12, tname, "21")
		vvd(t, rota[1][1], 2, 1e-12, tname, "22")
		vvd(t, rota[1][2], 3, 1e-12, tname, "23")

		vvd(t, rota[2][0], 3.500207892850427330, 1e-12, tname, "31")
		vvd(t, rota[2][1], 4.779889022262298150, 1e-12, tname, "32")
		vvd(t, rota[2][2], 5.381899160903798712, 1e-12, tname, "33")
	}
}

func BenchmarkRy(b *testing.B) {

	var theta float64
	var r [3][3]float64

	theta = 0.3456789

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
		fn  func(float64, [3][3]float64) [3][3]float64
	}{
		{"cgo", CgoRy},
		{"go", GoRy},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(theta, r)
			}
		})
	}
}
