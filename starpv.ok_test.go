package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t S t a r p v
//  - - - - - - - - - - -
//
//  Test Starpv function.
//
//  Called:  Starpv, vvd, viv
//
//  This revision:  2017 March 15
//
func TestStarpv(t *testing.T) {
	const fname = "Starpv"
	var pv [2][3]float64
	var err en.ErrNum
	var ra, dec, pmr, pmd, px, rv float64

	ra = 0.01686756
	dec = -1.093989828
	pmr = -1.78323516e-5
	pmd = 2.336024047e-6
	px = 0.74723
	rv = -21.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			[2][3]float64, en.ErrNum)
	}{
		{"cgo", CgoStarpv},
		{"go", GoStarpv},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		pv, err = test.fn(ra, dec, pmr, pmd, px, rv)

		vvd(t, pv[0][0], 126668.5912743160601, 1e-10,
			tname, "11")
		vvd(t, pv[0][1], 2136.792716839935195, 1e-12,
			tname, "12")
		vvd(t, pv[0][2], -245251.2339876830091, 1e-10,
			tname, "13")

		vvd(t, pv[1][0], -0.4051854008955659551e-2, 1e-13,
			tname, "21")
		vvd(t, pv[1][1], -0.6253919754414777970e-2, 1e-15,
			tname, "22")
		vvd(t, pv[1][2], 0.1189353714588109341e-1, 1e-13,
			tname, "23")

		errT(t, nil, err, tname, "0")
	}
}

func BenchmarkStarpv(b *testing.B) {
	var ra, dec, pmr, pmd, px, rv float64

	ra = 0.01686756
	dec = -1.093989828
	pmr = -1.78323516e-5
	pmd = 2.336024047e-6
	px = 0.74723
	rv = -21.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64) (
			[2][3]float64, en.ErrNum)
	}{
		{"cgo", CgoStarpv},
		{"go", GoStarpv},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(ra, dec, pmr, pmd, px, rv)
			}
		})
	}
}
