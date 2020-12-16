package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A p i o 1 3
//  - - - - - - - - - - -
//
//  Test Apio13 function.
//
//  Called:  Apio13, vvd, viv
//
//  This revision:  2013 October 4
//
func TestApio13(t *testing.T) {
	const fname = "Apio13"
	var utc1, utc2, dut1, elong, phi,
		hm, xp, yp, phpa, tc, rh, wl float64
	var astrom ASTROM

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
		fn  func(a1, a2, a3, a4, a5, a6,
			a7, a8, a9, a10, a11, a12 float64,
			a13 ASTROM) (ASTROM, error)
	}{
		{"cgo", CgoApio13},
		{"go", GoApio13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		astr, err := test.fn(utc1, utc2, dut1, elong, phi, hm, xp, yp,
			phpa, tc, rh, wl, astrom)

		vvd(t, astr.along, -0.5278008060301974337, 1e-12,
			tname, "along")
		vvd(t, astr.xpl, 0.1133427418174939329e-5, 1e-17,
			tname, "xpl")
		vvd(t, astr.ypl, 0.1453347595745898629e-5, 1e-17,
			tname, "ypl")
		vvd(t, astr.sphi, -0.9440115679003211329, 1e-12,
			tname, "sphi")
		vvd(t, astr.cphi, 0.3299123514971474711, 1e-12,
			tname, "cphi")
		vvd(t, astr.diurab, 0.5135843661699913529e-6, 1e-12,
			tname, "diurab")
		vvd(t, astr.eral, 2.617608909189066140, 1e-12,
			tname, "eral")
		vvd(t, astr.refa, 0.2014187785940396921e-3, 1e-15,
			tname, "refa")
		vvd(t, astr.refb, -0.2361408314943696227e-6, 1e-18,
			tname, "refb")
		errT(t, nil, err, tname, "err")
	}
}

func BenchmarkApio13(b *testing.B) {
	var utc1, utc2, dut1, elong, phi,
		hm, xp, yp, phpa, tc, rh, wl float64
	var astrom ASTROM

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
		fn  func(a1, a2, a3, a4, a5, a6,
			a7, a8, a9, a10, a11, a12 float64,
			a13 ASTROM) (ASTROM, error)
	}{
		{"cgo", CgoApio13},
		{"go", GoApio13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(utc1, utc2, dut1,
					elong, phi, hm, xp, yp,
					phpa, tc, rh, wl, astrom)
			}
		})
	}
}
