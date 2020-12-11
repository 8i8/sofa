package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t N u m a t
//  - - - - - - - - - -
//
//  Test Numat function.
//
//  Called:  Numat, vvd
//
//  This revision:  2013 August 7
//
func TestNumat(t *testing.T) {
	const fname = "Numat"
	var epsa, dpsi, deps float64
	var rmatn [3][3]float64

	epsa = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5
	deps = 0.4063239174001678826e-4

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) [3][3]float64
	}{
		{"cgo", CgoNumat},
		{"go", GoNumat},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rmatn = test.fn(epsa, dpsi, deps)

		vvd(t, rmatn[0][0], 0.9999999999536227949, 1e-12,
			tname, "11")
		vvd(t, rmatn[0][1], 0.8836239320236250577e-5, 1e-12,
			tname, "12")
		vvd(t, rmatn[0][2], 0.3830833447458251908e-5, 1e-12,
			tname, "13")

		vvd(t, rmatn[1][0], -0.8836083657016688588e-5, 1e-12,
			tname, "21")
		vvd(t, rmatn[1][1], 0.9999999991354654959, 1e-12,
			tname, "22")
		vvd(t, rmatn[1][2], -0.4063240865361857698e-4, 1e-12,
			tname, "23")

		vvd(t, rmatn[2][0], -0.3831192481833385226e-5, 1e-12,
			tname, "31")
		vvd(t, rmatn[2][1], 0.4063237480216934159e-4, 1e-12,
			tname, "32")
		vvd(t, rmatn[2][2], 0.9999999991671660407, 1e-12,
			tname, "33")
	}
}

func BenchmarkNumat(b *testing.B) {
	var epsa, dpsi, deps float64

	epsa = 0.4090789763356509900
	dpsi = -0.9630909107115582393e-5
	deps = 0.4063239174001678826e-4

	tests := []struct {
		ref string
		fn  func(a1, a2, a3 float64) [3][3]float64
	}{
		{"cgo", CgoNumat},
		{"go", GoNumat},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(epsa, dpsi, deps)
			}
		})
	}
}
