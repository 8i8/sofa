package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t P v s t a r
//  - - - - - - - - - - -
//
//  Test Pvstar function.
//
//  Called:  Pvstar, vvd, viv
//
//  This revision:  2017 March 15
//
func TestPvstar(t *testing.T) {
	const fname = "Pvstar"
	var ra, dec, pmr, pmd, px, rv float64
	var err error
	var pv [2][3]float64

	pv[0][0] = 126668.5912743160601
	pv[0][1] = 2136.792716839935195
	pv[0][2] = -245251.2339876830091

	pv[1][0] = -0.4051854035740712739e-2
	pv[1][1] = -0.6253919754866173866e-2
	pv[1][2] = 0.1189353719774107189e-1

	tests := []struct {
		ref string
		fn  func([2][3]float64) (c1, c2, c3, c4, c5, c6 float64,
			c7 error)
	}{
		{"cgo", CgoPvstar},
		{"go", GoPvstar},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref


		ra, dec, pmr, pmd, px, rv, err = test.fn(pv)

		vvd(t, ra, 0.1686756e-1, 1e-12, tname, "ra")
		vvd(t, dec, -1.093989828, 1e-12, tname, "dec")
		vvd(t, pmr, -0.1783235160000472788e-4, 1e-16,
			tname, "pmr")
		vvd(t, pmd, 0.2336024047000619347e-5, 1e-16,
			tname, "pmd")
		vvd(t, px, 0.74723, 1e-12, tname, "px")
		vvd(t, rv, -21.60000010107306010, 1e-11, tname, "rv")

		errT(t, nil, err, tname, "0")
	}
}

func BenchmarkPvstar(b *testing.B) {
	var pv [2][3]float64

	pv[0][0] = 126668.5912743160601
	pv[0][1] = 2136.792716839935195
	pv[0][2] = -245251.2339876830091

	pv[1][0] = -0.4051854035740712739e-2
	pv[1][1] = -0.6253919754866173866e-2
	pv[1][2] = 0.1189353719774107189e-1

	tests := []struct {
		ref string
		fn  func([2][3]float64) (c1, c2, c3, c4, c5, c6 float64,
			c7 error)
	}{
		{"cgo", CgoPvstar},
		{"go", GoPvstar},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(pv)
			}
		})
	}
}
