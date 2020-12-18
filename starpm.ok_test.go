package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t S t a r p m
//  - - - - - - - - - - -
//
//  Test Starpm function.
//
//  Called:  Starpm, vvd, viv
//
//  This revision:  2017 March 15
//
func TestStarpm(t *testing.T) {
	const fname = "Starpm"
	var ra2, dec2, pmr2, pmd2, px2, rv2 float64
	var err en.ErrNum
	var ra1, dec1, pmr1, pmd1, px1, rv1 float64

	ra1 = 0.01686756
	dec1 = -1.093989828
	pmr1 = -1.78323516e-5
	pmd1 = 2.336024047e-6
	px1 = 0.74723
	rv1 = -21.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5,
			a6, a7, a8, a9, a10 float64) (
			c1, c2, c3, c4, c5, c6 float64, c7 en.ErrNum)
	}{
		{"cgo", CgoStarpm},
		{"go", GoStarpm},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		ra2, dec2, pmr2, pmd2, px2, rv2, err = test.fn(
			ra1, dec1, pmr1, pmd1, px1, rv1,
			2400000.5, 50083.0, 2400000.5, 53736.0)

		vvd(t, ra2, 0.01668919069414256149, 1e-13,
			tname, "ra")
		vvd(t, dec2, -1.093966454217127897, 1e-13,
			tname, "dec")
		vvd(t, pmr2, -0.1783662682153176524e-4, 1e-17,
			tname, "pmr")
		vvd(t, pmd2, 0.2338092915983989595e-5, 1e-17,
			tname, "pmd")
		vvd(t, px2, 0.7473533835317719243, 1e-13,
			tname, "px")
		vvd(t, rv2, -21.59905170476417175, 1e-11,
			tname, "rv")

		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkStarpm(b *testing.B) {
	var ra1, dec1, pmr1, pmd1, px1, rv1 float64

	ra1 = 0.01686756
	dec1 = -1.093989828
	pmr1 = -1.78323516e-5
	pmd1 = 2.336024047e-6
	px1 = 0.74723
	rv1 = -21.6

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5,
			a6, a7, a8, a9, a10 float64) (
			c1, c2, c3, c4, c5, c6 float64, c7 en.ErrNum)
	}{
		{"cgo", CgoStarpm},
		{"go", GoStarpm},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(ra1, dec1, pmr1, pmd1, px1, rv1,
					2400000.5, 50083.0, 2400000.5,
					53736.0)
			}
		})
	}
}
