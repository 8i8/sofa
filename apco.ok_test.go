package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A p c o
//  - - - - - - - - -
//
//  Test Apco function.
//
//  Called:  Apco, vvd
//
//  This revision:  2017 March 15
//
func TestApco(t *testing.T) {
	const fname = "Apco"
	var date1, date2, x, y, s, theta, elong,
		phi, hm, xp, yp, sp, refa, refb float64
	var ebpv [2][3]float64
	var ehp [3]float64

	date1 = 2456384.5
	date2 = 0.970031644
	ebpv[0][0] = -0.974170438
	ebpv[0][1] = -0.211520082
	ebpv[0][2] = -0.0917583024
	ebpv[1][0] = 0.00364365824
	ebpv[1][1] = -0.0154287319
	ebpv[1][2] = -0.00668922024
	ehp[0] = -0.973458265
	ehp[1] = -0.209215307
	ehp[2] = -0.0906996477
	x = 0.0013122272
	y = -2.92808623e-5
	s = 3.05749468e-8
	theta = 3.14540971
	elong = -0.527800806
	phi = -1.2345856
	hm = 2738.0
	xp = 2.47230737e-7
	yp = 1.82640464e-6
	sp = -3.01974337e-11
	refa = 0.000201418779
	refb = -2.36140831e-7

	tests := []struct {
		ref string
		fn  func(a1, a2 float64,
			a3 [2][3]float64, a4 [3]float64,
			a5, a6, a7, a8 float64,
			a9, a10, a11 float64,
			a12, a13, a14 float64,
			a15, a16 float64) (b1 ASTROM)
	}{
		{"cgo", CgoApco},
		{"go", GoApco},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom := test.fn(date1, date2,
			ebpv, ehp,
			x, y, s, theta,
			elong, phi, hm,
			xp, yp, sp,
			refa, refb)

		vvd(t, astrom.pmt, 13.25248468622587269, 1e-11,
			tname, "pmt")
		vvd(t, astrom.eb[0], -0.9741827110630322720, 1e-12,
			tname, "eb(1)")
		vvd(t, astrom.eb[1], -0.2115130190135344832, 1e-12,
			tname, "eb(2)")
		vvd(t, astrom.eb[2], -0.09179840186949532298, 1e-12,
			tname, "eb(3)")
		vvd(t, astrom.eh[0], -0.9736425571689739035, 1e-12,
			tname, "eh(1)")
		vvd(t, astrom.eh[1], -0.2092452125849330936, 1e-12,
			tname, "eh(2)")
		vvd(t, astrom.eh[2], -0.09075578152243272599, 1e-12,
			tname, "eh(3)")
		vvd(t, astrom.em, 0.9998233241709957653, 1e-12,
			tname, "em")
		vvd(t, astrom.v[0], 0.2078704992916728762e-4, 1e-16,
			tname, "v(1)")
		vvd(t, astrom.v[1], -0.8955360107151952319e-4, 1e-16,
			tname, "v(2)")
		vvd(t, astrom.v[2], -0.3863338994288951082e-4, 1e-16,
			tname, "v(3)")
		vvd(t, astrom.bm1, 0.9999999950277561236, 1e-12,
			tname, "bm1")
		vvd(t, astrom.bpn[0][0], 0.9999991390295159156, 1e-12,
			tname, "bpn(1,1)")
		vvd(t, astrom.bpn[1][0], 0.4978650072505016932e-7, 1e-12,
			tname, "bpn(2,1)")
		vvd(t, astrom.bpn[2][0], 0.1312227200000000000e-2, 1e-12,
			tname, "bpn(3,1)")
		vvd(t, astrom.bpn[0][1], -0.1136336653771609630e-7, 1e-12,
			tname, "bpn(1,2)")
		vvd(t, astrom.bpn[1][1], 0.9999999995713154868, 1e-12,
			tname, "bpn(2,2)")
		vvd(t, astrom.bpn[2][1], -0.2928086230000000000e-4, 1e-12,
			tname, "bpn(3,2)")
		vvd(t, astrom.bpn[0][2], -0.1312227200895260194e-2, 1e-12,
			tname, "bpn(1,3)")
		vvd(t, astrom.bpn[1][2], 0.2928082217872315680e-4, 1e-12,
			tname, "bpn(2,3)")
		vvd(t, astrom.bpn[2][2], 0.9999991386008323373, 1e-12,
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
		vvd(t, astrom.eral, 2.617608903969802566, 1e-12,
			tname, "eral")
		vvd(t, astrom.refa, 0.2014187790000000000e-3, 1e-15,
			tname, "refa")
		vvd(t, astrom.refb, -0.2361408310000000000e-6, 1e-18,
			tname, "refb")
	}
}

func BenchmarkApco(b *testing.B) {
	var date1, date2, x, y, s, theta, elong,
		phi, hm, xp, yp, sp, refa, refb float64
	var ebpv [2][3]float64
	var ehp [3]float64

	date1 = 2456384.5
	date2 = 0.970031644
	ebpv[0][0] = -0.974170438
	ebpv[0][1] = -0.211520082
	ebpv[0][2] = -0.0917583024
	ebpv[1][0] = 0.00364365824
	ebpv[1][1] = -0.0154287319
	ebpv[1][2] = -0.00668922024
	ehp[0] = -0.973458265
	ehp[1] = -0.209215307
	ehp[2] = -0.0906996477
	x = 0.0013122272
	y = -2.92808623e-5
	s = 3.05749468e-8
	theta = 3.14540971
	elong = -0.527800806
	phi = -1.2345856
	hm = 2738.0
	xp = 2.47230737e-7
	yp = 1.82640464e-6
	sp = -3.01974337e-11
	refa = 0.000201418779
	refb = -2.36140831e-7

	tests := []struct {
		ref string
		fn  func(a1, a2 float64,
			a3 [2][3]float64, a4 [3]float64,
			a5, a6, a7, a8 float64,
			a9, a10, a11 float64,
			a12, a13, a14 float64,
			a15, a16 float64) (b1 ASTROM)
	}{
		{"cgo", CgoApco},
		{"go", GoApco},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(date1, date2,
					ebpv, ehp,
					x, y, s, theta,
					elong, phi, hm,
					xp, yp, sp,
					refa, refb)
			}
		})
	}
}
