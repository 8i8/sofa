package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A t o i 1 3
//  - - - - - - - - - - -
//
//  Test Atoi13 function.
//
//  Called:  Atoi13, vvd, viv
//
//  This revision:  2013 October 3
//
func TestAtoi13(t *testing.T) {
	const fname = "Atoi13"
	var utc1, utc2, dut1, elong, phi, hm, xp, yp, phpa, tc, rh, wl,
		ob1, ob2, ri, di float64
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

	tests := []struct {
		ref string
		fn  func(a1 string, a2, a3, a4, a5, a6, a7, a8, a9, a10,
			a11, a12, a13, a14, a15 float64) (
			c1, c2 float64, c3 error)
	}{
		{"cgo", CgoAtoi13},
		{"go", GoAtoi13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		ob1 = 2.710085107986886201
		ob2 = 0.1717653435758265198
		ri, di, err = test.fn("R", ob1, ob2, utc1, utc2, dut1,
			elong, phi, hm, xp, yp, phpa, tc, rh, wl)
		vvd(t, ri, 2.710121574449135955, 1e-12, tname, "R/ri")
		vvd(t, di, 0.1729371839114567725, 1e-12, tname, "R/di")
		errT(t, nil, err, tname+" R/J")

		ob1 = -0.09247619879782006106
		ob2 = 0.1717653435758265198
		ri, di, err = test.fn("H", ob1, ob2, utc1, utc2, dut1,
			elong, phi, hm, xp, yp, phpa, tc, rh, wl)
		vvd(t, ri, 2.710121574449135955, 1e-12, tname, "H/ri")
		vvd(t, di, 0.1729371839114567725, 1e-12, tname, "H/di")
		errT(t, nil, err, tname+" H/J")

		ob1 = 0.09233952224794989993
		ob2 = 1.407758704513722461
		ri, di, err = test.fn("A", ob1, ob2, utc1, utc2, dut1,
			elong, phi, hm, xp, yp, phpa, tc, rh, wl)
		vvd(t, ri, 2.710121574449135955, 1e-12, tname, "A/ri")
		vvd(t, di, 0.1729371839114567728, 1e-12, tname, "A/di")
		errT(t, nil, err, tname+" A/J")
	}
}

func BenchmarkAtoi13(b *testing.B) {
	var utc1, utc2, dut1, elong, phi, hm, xp, yp, phpa, tc, rh, wl,
		ob1, ob2 float64

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
		fn  func(a1 string, a2, a3, a4, a5, a6, a7, a8, a9, a10,
			a11, a12, a13, a14, a15 float64) (
			c1, c2 float64, c3 error)
	}{
		{"cgo", CgoAtoi13},
		{"go", GoAtoi13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ob1 = 2.710085107986886201
				ob2 = 0.1717653435758265198
				test.fn("R", ob1, ob2, utc1, utc2, dut1,
				elong, phi, hm, xp, yp, phpa, tc, rh,
				wl)

				ob1 = -0.09247619879782006106
				ob2 = 0.1717653435758265198
				test.fn("H", ob1, ob2, utc1, utc2, dut1,
				elong, phi, hm, xp, yp, phpa, tc, rh,
				wl)

				ob1 = 0.09233952224794989993
				ob2 = 1.407758704513722461
				test.fn("A", ob1, ob2, utc1, utc2, dut1,
				elong, phi, hm, xp, yp, phpa, tc, rh,
				wl)
			}
		})
	}
}
