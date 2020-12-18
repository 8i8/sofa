package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t P v u p
//  - - - - - - - - -
//
//  Test Pvup function.
//
//  Called:  Pvup, vvd
//
//  This revision:  2013 August 7
//
func TestPvup(t *testing.T) {
	const fname = "Pvup"
	var p [3]float64
	var pv [2][3]float64

	pv[0][0] = 126668.5912743160734
	pv[0][1] = 2136.792716839935565
	pv[0][2] = -245251.2339876830229

	pv[1][0] = -0.4051854035740713039e-2
	pv[1][1] = -0.6253919754866175788e-2
	pv[1][2] = 0.1189353719774107615e-1

	tests := []struct {
		ref string
		fn  func(float64, [2][3]float64) [3]float64
	}{
		{"cgo", CgoPvup},
		{"go", GoPvup},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		p = test.fn(2920.0, pv)

		vvd(t, p[0], 126656.7598605317105, 1e-12, tname, "1")
		vvd(t, p[1], 2118.531271155726332, 1e-12, tname, "2")
		vvd(t, p[2], -245216.5048590656190, 1e-12, tname, "3")
	}
}

func BenchmarkPvup(b *testing.B) {
	var pv [2][3]float64

	pv[0][0] = 126668.5912743160734
	pv[0][1] = 2136.792716839935565
	pv[0][2] = -245251.2339876830229

	pv[1][0] = -0.4051854035740713039e-2
	pv[1][1] = -0.6253919754866175788e-2
	pv[1][2] = 0.1189353719774107615e-1

	tests := []struct {
		ref string
		fn  func(float64, [2][3]float64) [3]float64
	}{
		{"cgo", CgoPvup},
		{"go", GoPvup},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2920.0, pv)
			}
		})
	}
}
