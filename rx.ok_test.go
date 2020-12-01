package sofa

import "testing"

//
//  - - - - - - -
//   T e s t R x
//  - - - - - - -
//
//  Test Rx function.
//
//  Called:  Rx, vvd
//
//  This revision:  2013 August 7
//
func TestRx(t *testing.T) {
	const fname = "Rx"
	var phi float64
	var r [3][3]float64
	phi = 0.3456789

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
		{"cgo", CgoRx},
		{"go", GoRx},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rota := test.fn(phi, r)

		vvd(t, rota[0][0], 2.0, 0.0, tname, "11")
		vvd(t, rota[0][1], 3.0, 0.0, tname, "12")
		vvd(t, rota[0][2], 2.0, 0.0, tname, "13")

		vvd(t, rota[1][0], 3.839043388235612460,
			1e-12, tname, "21")
		vvd(t, rota[1][1], 3.237033249594111899,
			1e-12, tname, "22")
		vvd(t, rota[1][2], 4.516714379005982719,
			1e-12, tname, "23")

		vvd(t, rota[2][0], 1.806030415924501684,
			1e-12, tname, "31")
		vvd(t, rota[2][1], 3.085711545336372503,
			1e-12, tname, "32")
		vvd(t, rota[2][2], 3.687721683977873065,
			1e-12, tname, "33")
	}
}

func BenchmarkRx(b *testing.B) {
	var phi float64
	var r [3][3]float64
	phi = 0.3456789

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
		{"cgo", CgoRx},
		{"go", GoRx},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(phi, r)
			}
		})
	}
}
