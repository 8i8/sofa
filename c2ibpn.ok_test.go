package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t C 2 i b p n
//  - - - - - - - - - - -
//
//  Test C2ibpn function.
//
//  Called:  C2ibpn, vvd
//
//  This revision:  2013 August 7
//
func TestC2ibpn(t *testing.T) {
	const fname = "C2ibpn"
	var rbpn, rc2i [3][3]float64

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
		fn  func(a1, a2 float64, a3 [3][3]float64) [3][3]float64
	}{
		{"cgo", CgoC2ibpn},
		{"go", GoC2ibpn},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		rc2i = test.fn(2400000.5, 50123.9999, rbpn)

		vvd(t, rc2i[0][0], 0.9999994021664089977, 1e-12,
			tname, "11")
		vvd(t, rc2i[0][1], -0.3869195948017503664e-8, 1e-12,
			tname, "12")
		vvd(t, rc2i[0][2], -0.1093465511383285076e-2, 1e-12,
			tname, "13")

		vvd(t, rc2i[1][0], 0.5068413965715446111e-7, 1e-12,
			tname, "21")
		vvd(t, rc2i[1][1], 0.9999999990835075686, 1e-12,
			tname, "22")
		vvd(t, rc2i[1][2], 0.4281334246452708915e-4, 1e-12,
			tname, "23")

		vvd(t, rc2i[2][0], 0.1093465510215479000e-2, 1e-12,
			tname, "31")
		vvd(t, rc2i[2][1], -0.4281337229063151000e-4, 1e-12,
			tname, "32")
		vvd(t, rc2i[2][2], 0.9999994012499173103, 1e-12,
			tname, "33")
	}
}

func BenchmarkC2ibpn(b *testing.B) {
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
		fn  func(a1, a2 float64, a3 [3][3]float64) [3][3]float64
	}{
		{"cgo", CgoC2ibpn},
		{"go", GoC2ibpn},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 50123.9999, rbpn)
			}
		})
	}
}
