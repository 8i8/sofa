package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t B p n 2 x y
//  - - - - - - - - - - -
//
//  Test Bpn2xy function.
//
//  Called:  Bpn2xy, vvd
//
//  This revision:  2013 August 7
//
func TestBpn2xy(t *testing.T) {
	const fname = "Bpn2xy"
	var rbpn [3][3]float64

	rbpn[0][0] = 9.999962358680738e-1
	rbpn[0][1] = -2.516417057665452e-3
	rbpn[0][2] = -1.093569785342370e-3

	rbpn[1][0] = 2.516462370370876e-3
	rbpn[1][1] = 9.999968329010883e-1
	rbpn[1][2] = 4.006159587358310e-5

	rbpn[2][0] = 1.093465510215479e-3
	rbpn[2][1] = -4.281337229063151e-5
	rbpn[2][2] = 9.999994012499173e-1

	tests := []struct {
		ref string
		fn  func([3][3]float64) (a, b float64)
	}{
		{"cgo", CgoBpn2xy},
		{"go", GoBpn2xy},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		x, y := test.fn(rbpn)

		vvd(t, x, 1.093465510215479e-3, 1e-12, tname, "x")
		vvd(t, y, -4.281337229063151e-5, 1e-12, tname, "y")
	}
}

func BenchmarkBpn2xy(b *testing.B) {
	var rbpn [3][3]float64

	rbpn[0][0] = 9.999962358680738e-1
	rbpn[0][1] = -2.516417057665452e-3
	rbpn[0][2] = -1.093569785342370e-3

	rbpn[1][0] = 2.516462370370876e-3
	rbpn[1][1] = 9.999968329010883e-1
	rbpn[1][2] = 4.006159587358310e-5

	rbpn[2][0] = 1.093465510215479e-3
	rbpn[2][1] = -4.281337229063151e-5
	rbpn[2][2] = 9.999994012499173e-1

	tests := []struct {
		ref string
		fn  func([3][3]float64) (a, b float64)
	}{
		{"cgo", CgoBpn2xy},
		{"go", GoBpn2xy},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(rbpn)
			}
		})
	}
}
