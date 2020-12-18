package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t P m s a f e
//  - - - - - - - - - - -
//
//  Test Pmsafe function.
//
//  Called:  Pmsafe, vvd, viv
//
//  This revision:  2017 March 15
//
func TestPmsafe(t *testing.T) {
	const fname = "Pmasafe"
	var err en.ErrNum
	var ra1, dec1, pmr1, pmd1, px1, rv1, ep1a, ep1b, ep2a, ep2b,
		ra2, dec2, pmr2, pmd2, px2, rv2 float64

	ra1 = 1.234
	dec1 = 0.789
	pmr1 = 1e-5
	pmd1 = -2e-5
	px1 = 1e-2
	rv1 = 10.0
	ep1a = 2400000.5
	ep1b = 48348.5625
	ep2a = 2400000.5
	ep2b = 51544.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6,
			a7, a8, a9, a10 float64) (
			c1, c2, c3, c4, c5, c6 float64, c7 en.ErrNum)
	}{
		{"cgo", CgoPmsafe},
		{"go", GoPmsafe},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		ra2, dec2, pmr2, pmd2, px2, rv2, err = test.fn(
			ra1, dec1, pmr1, pmd1, px1, rv1,
			ep1a, ep1b, ep2a, ep2b)

		vvd(t, ra2, 1.234087484501017061, 1e-12,
			tname, "ra2")
		vvd(t, dec2, 0.7888249982450468567, 1e-12,
			tname, "dec2")
		vvd(t, pmr2, 0.9996457663586073988e-5, 1e-12,
			tname, "pmr2")
		vvd(t, pmd2, -0.2000040085106754565e-4, 1e-16,
			tname, "pmd2")
		vvd(t, px2, 0.9999997295356830666e-2, 1e-12,
			tname, "px2")
		vvd(t, rv2, 10.38468380293920069, 1e-10,
			tname, "rv2")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkPmsafe(b *testing.B) {
	var ra1, dec1, pmr1, pmd1, px1, rv1,
		ep1a, ep1b, ep2a, ep2b float64

	ra1 = 1.234
	dec1 = 0.789
	pmr1 = 1e-5
	pmd1 = -2e-5
	px1 = 1e-2
	rv1 = 10.0
	ep1a = 2400000.5
	ep1b = 48348.5625
	ep2a = 2400000.5
	ep2b = 51544.5

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6,
			a7, a8, a9, a10 float64) (
			c1, c2, c3, c4, c5, c6 float64, c7 en.ErrNum)
	}{
		{"cgo", CgoPmsafe},
		{"go", GoPmsafe},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(ra1, dec1, pmr1, pmd1, px1, rv1,
					ep1a, ep1b, ep2a, ep2b)
			}
		})
	}
}
