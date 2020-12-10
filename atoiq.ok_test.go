package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t A t o i q
//  - - - - - - - - - -
//
//  Test Atoiq function.
//
//  Called:  Apio13, iauAtoiq, vvd
//
//  This revision:  2013 October 4
//
func TestAtoiq(t *testing.T) {
	const fname = "Atoiq"
	var utc1, utc2, dut1, elong, phi, hm, xp, yp, phpa, tc, rh, wl,
		ob1, ob2, ri, di float64
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

	tests := []struct {
		ref string
		fn  func(a1 string, a2, a3 float64, a4 ASTROM) (
			c1, c2 float64)
		fnAssist func(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10,
			a11, a12 float64, a13 ASTROM) (ASTROM, error)
	}{
		{"cgo", CgoAtoiq, CgoApio13},
		{"go", GoAtoiq, GoApio13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom, err = test.fnAssist(utc1, utc2, dut1, elong,
			phi, hm, xp, yp, phpa, tc, rh, wl, astrom)
		if err != nil {
			t.Errorf("%s error: %s", tname, err)
		}

		ob1 = 2.710085107986886201
		ob2 = 0.1717653435758265198
		ri, di = test.fn("R", ob1, ob2, astrom)
		vvd(t, ri, 2.710121574449135955, 1e-12,
			tname, "R/ri")
		vvd(t, di, 0.1729371839114567725, 1e-12,
			tname, "R/di")

		ob1 = -0.09247619879782006106
		ob2 = 0.1717653435758265198
		ri, di = test.fn("H", ob1, ob2, astrom)
		vvd(t, ri, 2.710121574449135955, 1e-12,
			tname, "H/ri")
		vvd(t, di, 0.1729371839114567725, 1e-12,
			tname, "H/di")

		ob1 = 0.09233952224794989993
		ob2 = 1.407758704513722461
		ri, di = test.fn("A", ob1, ob2, astrom)
		vvd(t, ri, 2.710121574449135955, 1e-12,
			tname, "A/ri")
		vvd(t, di, 0.1729371839114567728, 1e-12,
			tname, "A/di")
	}
}

func BenchmarkAtoiq(b *testing.B) {
	const fname = "Atoiq"
	var utc1, utc2, dut1, elong, phi, hm, xp, yp, phpa, tc, rh, wl,
		ob1, ob2 float64
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

	tests := []struct {
		ref string
		fn  func(a1 string, a2, a3 float64, a4 ASTROM) (
			c1, c2 float64)
		fnAssist func(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10,
			a11, a12 float64, a13 ASTROM) (ASTROM, error)
	}{
		{"cgo", CgoAtoiq, CgoApio13},
		{"go", GoAtoiq, GoApio13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom, err = test.fnAssist(utc1, utc2, dut1, elong,
			phi, hm, xp, yp, phpa, tc, rh, wl, astrom)
		if err != nil {
			b.Errorf("%s error: %s", tname, err)
		}
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {

				ob1 = 2.710085107986886201
				ob2 = 0.1717653435758265198
				test.fn("R", ob1, ob2, astrom)

				ob1 = -0.09247619879782006106
				ob2 = 0.1717653435758265198
				test.fn("H", ob1, ob2, astrom)

				ob1 = 0.09233952224794989993
				ob2 = 1.407758704513722461
				test.fn("A", ob1, ob2, astrom)
			}
		})
	}
}
