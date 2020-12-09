package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A p i o
//  - - - - - - - - -
//
//  Test Apio function.
//
//  Called:  Apio, vvd
//
//  This revision:  2013 October 3
//
func TestApio(t *testing.T) {
	const fname = "Apio"
	var sp, theta, elong, phi, hm, xp, yp, refa, refb float64
	var astrom ASTROM

	sp = -3.01974337e-11
	theta = 3.14540971
	elong = -0.527800806
	phi = -1.2345856
	hm = 2738.0
	xp = 2.47230737e-7
	yp = 1.82640464e-6
	refa = 0.000201418779
	refb = -2.36140831e-7

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7, a8, a9 float64,
			a10 ASTROM) ASTROM
	}{
		{"cgo", CgoApio},
		{"go", GoApio},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		astrom = test.fn(sp, theta,
			elong, phi, hm, xp, yp,
			refa, refb, astrom)

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
		vvd(t, astrom.diurab, 0.5135843661699913529e-6, 1e-12,
			tname, "diurab")
		vvd(t, astrom.eral, 2.617608903969802566, 1e-12,
			tname, "eral")
		vvd(t, astrom.refa, 0.2014187790000000000e-3, 1e-15,
			tname, "refa")
		vvd(t, astrom.refb, -0.2361408310000000000e-6, 1e-18,
			tname, "refb")
	}
}

func BenchmarkApio(b *testing.B) {
	var sp, theta, elong, phi, hm, xp, yp, refa, refb float64
	var astrom ASTROM

	sp = -3.01974337e-11
	theta = 3.14540971
	elong = -0.527800806
	phi = -1.2345856
	hm = 2738.0
	xp = 2.47230737e-7
	yp = 1.82640464e-6
	refa = 0.000201418779
	refb = -2.36140831e-7

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6, a7, a8, a9 float64,
			a10 ASTROM) ASTROM
	}{
		{"cgo", CgoApio},
		{"go", GoApio},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(sp, theta,
					elong, phi, hm, xp, yp,
					refa, refb, astrom)
			}
		})
	}
}
