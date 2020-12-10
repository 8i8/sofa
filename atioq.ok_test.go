package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t A t i o q
//  - - - - - - - - - -
//
//  Test Atioq function.
//
//  Called:  Apio13, iauAtioq, vvd, viv
//
//  This revision:  2013 October 4
//
func TestAtioq(t *testing.T) {
	const fname = "Atioq"
	var utc1, utc2, dut1, elong, phi, hm, xp, yp, phpa, tc, rh, wl,
		ri, di, aob, zob, hob, dob, rob float64
	var astrom ASTROM
	var err error

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
	ri = 2.710121572969038991
	di = 0.1729371367218230438

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 ASTROM) (c1, c2, c3,
			c4, c5 float64)
		fnAssist func(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10,
			a11, a12 float64, a13 ASTROM) (ASTROM, error)
	}{
		{"cgo", CgoAtioq, CgoApio13},
		{"go", GoAtioq, GoApio13},
	}

	for _, test := range tests {
		astrom, err = test.fnAssist(utc1, utc2, dut1, elong,
			phi, hm, xp, yp, phpa, tc, rh, wl, astrom)
		if err != nil {
			t.Errorf("%s error: %s", fname, err)
		}
		tname := fname + " " + test.ref

		aob, zob, hob, dob, rob = test.fn(ri, di, astrom)

		vvd(t, aob, 0.09233952224794989993, 1e-12,
			tname, "aob")
		vvd(t, zob, 1.407758704513722461, 1e-12,
			tname, "zob")
		vvd(t, hob, -0.09247619879782006106, 1e-12,
			tname, "hob")
		vvd(t, dob, 0.1717653435758265198, 1e-12,
			tname, "dob")
		vvd(t, rob, 2.710085107986886201, 1e-12,
			tname, "rob")
	}
}

func BenchmarkAtioq(b *testing.B) {
	const fname = "Atioq"
	var utc1, utc2, dut1, elong, phi, hm, xp, yp, phpa, tc, rh, wl,
		ri, di float64
	var astrom ASTROM
	var err error

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
	ri = 2.710121572969038991
	di = 0.1729371367218230438

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 ASTROM) (c1, c2, c3,
			c4, c5 float64)
		fnAssist func(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10,
			a11, a12 float64, a13 ASTROM) (ASTROM, error)
	}{
		{"cgo", CgoAtioq, CgoApio13},
		{"go", GoAtioq, GoApio13},
	}

	for _, test := range tests {
		astrom, err = test.fnAssist(utc1, utc2, dut1, elong,
			phi, hm, xp, yp, phpa, tc, rh, wl, astrom)
		if err != nil {
			b.Errorf("%s error: %s", fname, err)
		}
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(ri, di, astrom)
			}
		})
	}
}
