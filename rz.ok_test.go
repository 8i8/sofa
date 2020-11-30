package sofa

import "testing"

//
//  - - - - - - -
//   T e s t R z
//  - - - - - - -
//
//  Test Rz function.
//
//  Called:  Rz, vvd
//
//  This revision:  2013 August 7
//
func TestRz(t *testing.T) {
	const fname = "Rz"
	var psi float64
	var r [3][3]float64

	psi = 0.3456789

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
		{"cgo", Rz},
		{"go", goRz},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rota := test.fn(psi, r)

		vvd(t, rota[0][0], 2.898197754208926769, 1e-12, tname, "11")
		vvd(t, rota[0][1], 3.500207892850427330, 1e-12, tname, "12")
		vvd(t, rota[0][2], 2.898197754208926769, 1e-12, tname, "13")

		vvd(t, rota[1][0], 2.144865911309686813, 1e-12, tname, "21")
		vvd(t, rota[1][1], 0.865184781897815993, 1e-12, tname, "22")
		vvd(t, rota[1][2], 2.144865911309686813, 1e-12, tname, "23")

		vvd(t, rota[2][0], 3.0, 1e-12, tname, "31")
		vvd(t, rota[2][1], 4.0, 1e-12, tname, "32")
		vvd(t, rota[2][2], 5.0, 1e-12, tname, "33")
	}
}

func BenchmarkRz(b *testing.B) {
	var psi float64
	var r [3][3]float64

	psi = 0.3456789

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
		{"cgo", Rz},
		{"go", goRz},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(psi, r)
			}
		})
	}
}
