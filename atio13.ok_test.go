package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A t i o 1 3
//  - - - - - - - - - - -
//
//  Test Atio13 function.
//
//  Called:  Atio13, vvd, viv
//
//  This revision:  2013 October 3
//
func TestAtio13(t *testing.T) {
	const fname = "Atio13"
	var ri, di, utc1, utc2, dut1, elong, phi, hm, xp, yp,
		phpa, tc, rh, wl, aob, zob, hob, dob, rob float64
	var err error

	ri = 2.710121572969038991
	di = 0.1729371367218230438
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
			a12, a13, a14 float64) (
			c1, c2, c3, c4, c5 float64, c6 error)
	}{
		{"cgo", CgoAtio13},
		{"go", GoAtio13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		aob, zob, hob, dob, rob, err = test.fn(ri, di, utc1,
			utc2, dut1, elong, phi, hm, xp, yp, phpa, tc,
			rh, wl)

		vvd(t, aob, 0.09233952224794989993, 1e-12, tname, "aob")
		vvd(t, zob, 1.407758704513722461, 1e-12, tname, "zob")
		vvd(t, hob, -0.09247619879782006106, 1e-12, tname, "hob")
		vvd(t, dob, 0.1717653435758265198, 1e-12, tname, "dob")
		vvd(t, rob, 2.710085107986886201, 1e-12, tname, "rob")
		errT(t, nil, err, tname)
	}
}

func BenchmarkAtio13(b *testing.B) {
	var ri, di, utc1, utc2, dut1, elong, phi, hm, xp, yp,
		phpa, tc, rh, wl float64

	ri = 2.710121572969038991
	di = 0.1729371367218230438
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
			a12, a13, a14 float64) (
				c1, c2, c3, c4, c5 float64, c6 error)
	}{
		{"cgo", CgoAtio13},
		{"go", GoAtio13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(ri, di, utc1, utc2, dut1, elong,
				phi, hm, xp, yp, phpa, tc, rh, wl)
			}
		})
	}
}
