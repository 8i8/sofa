package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A p c o 1 3
//  - - - - - - - - - - -
//
//  Test Apco13 function.
//
//  Called:  Apco13, vvd, viv
//
//  This revision:  2017 March 15
//
func TestApco13(t *testing.T) {
	const fname = "Apco13"
	var utc1, utc2, dut1, elong, phi, hm,
		xp, yp, phpa, tc, rh, wl float64
	var astr ASTROM

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
			a13 ASTROM) (ASTROM, float64, error)
	}{
		{"cgo", CgoApco13},
		{"go", GoApco13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom, eo, err := test.fn(utc1, utc2, dut1, elong,
			phi, hm, xp, yp, phpa, tc, rh, wl, astr)

		vvd(t, astrom.pmt, 13.25248468622475727, 1e-11,
			tname, "pmt")
		vvd(t, astrom.eb[0], -0.9741827107320875162, 1e-12,
			tname, "eb(1)")
		vvd(t, astrom.eb[1], -0.2115130190489716682, 1e-12,
			tname, "eb(2)")
		vvd(t, astrom.eb[2], -0.09179840189496755339, 1e-12,
			tname, "eb(3)")
		vvd(t, astrom.eh[0], -0.9736425572586935247, 1e-12,
			tname, "eh(1)")
		vvd(t, astrom.eh[1], -0.2092452121603336166, 1e-12,
			tname, "eh(2)")
		vvd(t, astrom.eh[2], -0.09075578153885665295, 1e-12,
			tname, "eh(3)")
		vvd(t, astrom.em, 0.9998233240913898141, 1e-12,
			tname, "em")
		vvd(t, astrom.v[0], 0.2078704994520489246e-4, 1e-16,
			tname, "v(1)")
		vvd(t, astrom.v[1], -0.8955360133238868938e-4, 1e-16,
			tname, "v(2)")
		vvd(t, astrom.v[2], -0.3863338993055887398e-4, 1e-16,
			tname, "v(3)")
		vvd(t, astrom.bm1, 0.9999999950277561004, 1e-12,
			tname, "bm1")
		vvd(t, astrom.bpn[0][0], 0.9999991390295147999, 1e-12,
			tname, "bpn(1,1)")
		vvd(t, astrom.bpn[1][0], 0.4978650075315529277e-7, 1e-12,
			tname, "bpn(2,1)")
		vvd(t, astrom.bpn[2][0], 0.001312227200850293372, 1e-12,
			tname, "bpn(3,1)")
		vvd(t, astrom.bpn[0][1], -0.1136336652812486604e-7, 1e-12,
			tname, "bpn(1,2)")
		vvd(t, astrom.bpn[1][1], 0.9999999995713154865, 1e-12,
			tname, "bpn(2,2)")
		vvd(t, astrom.bpn[2][1], -0.2928086230975367296e-4, 1e-12,
			tname, "bpn(3,2)")
		vvd(t, astrom.bpn[0][2], -0.001312227201745553566, 1e-12,
			tname, "bpn(1,3)")
		vvd(t, astrom.bpn[1][2], 0.2928082218847679162e-4, 1e-12,
			tname, "bpn(2,3)")
		vvd(t, astrom.bpn[2][2], 0.9999991386008312212, 1e-12,
			tname, "bpn(3,3)")
		vvd(t, astrom.along, -0.5278008060301974337, 1e-12,
			tname, "along")
		vvd(t, astrom.xpl, 0.1133427418174939329e-5, 1e-17,
			tname, "xpl")
		vvd(t, astrom.ypl, 0.1453347595745898629e-5, 1e-17,
			tname, "ypl")
		vvd(t, astrom.sphi, -0.9440115679003211329, 1e-12,
			tname, "sphi")
		vvd(t, astrom.cphi, 0.3299123514971474711, 1e-12,
			tname, "cphi")
		vvd(t, astrom.diurab, 0, 0,
			tname, "diurab")
		vvd(t, astrom.eral, 2.617608909189066140, 1e-12,
			tname, "eral")
		vvd(t, astrom.refa, 0.2014187785940396921e-3, 1e-15,
			tname, "refa")
		vvd(t, astrom.refb, -0.2361408314943696227e-6, 1e-18,
			tname, "refb")
		vvd(t, eo, -0.003020548354802412839, 1e-14,
			tname, "eo")
			errT(t, nil, err, tname, "err")
	}
}

func BenchmarkApco13(b *testing.B) {
	var utc1, utc2, dut1, elong, phi, hm,
		xp, yp, phpa, tc, rh, wl float64
	var astr ASTROM

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
			a13 ASTROM) (ASTROM, float64, error)
	}{
		{"cgo", CgoApco13},
		{"go", GoApco13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(utc1, utc2, dut1,
					elong, phi, hm, xp, yp, phpa,
					tc, rh, wl, astr)
			}
		})
	}
}
