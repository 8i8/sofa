package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t A t c o 1 3
//  - - - - - - - - - - -
//
//  Test Atco13 function.
//
//  Called:  Atco13, vvd, viv
//
//  This revision:  2017 March 15
//
func TestAtco13(t *testing.T) {
	const fname = "Atco13"
	var rc, dc, pr, pd, px, rv, utc1, utc2, dut1,
		elong, phi, hm, xp, yp, phpa, tc, rh, wl,
		aob, zob, hob, dob, rob, eo float64
	var err en.ErrNum

	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0
	utc1 = 2456384.5
	utc2 = 0.969254051
	dut1 = 0.1550675
	elong = -0.527800806
	phi = -1.2345856
	hm = 2738.0
	xp = 2.47230737e-7
	yp = 1.82640464e-6
	phpa = 731.0
	tc = 12.8
	rh = 0.59
	wl = 0.55

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11,
			a12, a13, a14, a15, a16, a17, a18 float64) (
			c1, c2, c3, c4, c5, c6 float64, c7 en.ErrNum)
	}{
		{"cgo", CgoAtco13},
		{"go", GoAtco13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		aob, zob, hob, dob, rob, eo, err = test.fn(
			rc, dc, pr, pd, px, rv, utc1, utc2, dut1, elong,
			phi, hm, xp, yp, phpa, tc, rh, wl)

		vvd(t, aob, 0.09251774485385390973, 1e-12,
			tname, "aob")
		vvd(t, zob, 1.407661405256671703, 1e-12,
			tname, "zob")
		vvd(t, hob, -0.09265154431430045141, 1e-12,
			tname, "hob")
		vvd(t, dob, 0.1716626560074556029, 1e-12,
			tname, "dob")
		vvd(t, rob, 2.710260453503366591, 1e-12,
			tname, "rob")
		vvd(t, eo, -0.003020548354802412839, 1e-14,
			tname, "eo")

		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkAtco13(b *testing.B) {
	var rc, dc, pr, pd, px, rv, utc1, utc2, dut1,
		elong, phi, hm, xp, yp, phpa, tc, rh, wl float64

	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0
	utc1 = 2456384.5
	utc2 = 0.969254051
	dut1 = 0.1550675
	elong = -0.527800806
	phi = -1.2345856
	hm = 2738.0
	xp = 2.47230737e-7
	yp = 1.82640464e-6
	phpa = 731.0
	tc = 12.8
	rh = 0.59
	wl = 0.55

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11,
			a12, a13, a14, a15, a16, a17, a18 float64) (
			c1, c2, c3, c4, c5, c6 float64, c7 en.ErrNum)
	}{
		{"cgo", CgoAtco13},
		{"go", GoAtco13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(rc, dc, pr, pd, px, rv, utc1,
					utc2, dut1, elong, phi, hm, xp, yp,
					phpa, tc, rh, wl)
			}
		})
	}
}
