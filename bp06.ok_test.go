package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t B p 0 6
//  - - - - - - - - -
//
//  Test Bp06 function.
//
//  Called:  Bp06, vvd
//
//  This revision:  2013 August 7
//
func TestBp06(t *testing.T) {
	const fname = "Bp06"
	var rb, rp, rbp [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2, c3 [3][3]float64)
	}{
		{"cgo", CgoBp06},
		{"go", GoBp06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		rb, rp, rbp = test.fn(2400000.5, 50123.9999)

		vvd(t, rb[0][0], 0.9999999999999942497, 1e-12,
			tname, "rb11")
		vvd(t, rb[0][1], -0.7078368960971557145e-7, 1e-14,
			tname, "rb12")
		vvd(t, rb[0][2], 0.8056213977613185606e-7, 1e-14,
			tname, "rb13")
		vvd(t, rb[1][0], 0.7078368694637674333e-7, 1e-14,
			tname, "rb21")
		vvd(t, rb[1][1], 0.9999999999999969484, 1e-12,
			tname, "rb22")
		vvd(t, rb[1][2], 0.3305943742989134124e-7, 1e-14,
			tname, "rb23")
		vvd(t, rb[2][0], -0.8056214211620056792e-7, 1e-14,
			tname, "rb31")
		vvd(t, rb[2][1], -0.3305943172740586950e-7, 1e-14,
			tname, "rb32")
		vvd(t, rb[2][2], 0.9999999999999962084, 1e-12,
			tname, "rb33")

		vvd(t, rp[0][0], 0.9999995504864960278, 1e-12,
			tname, "rp11")
		vvd(t, rp[0][1], 0.8696112578855404832e-3, 1e-14,
			tname, "rp12")
		vvd(t, rp[0][2], 0.3778929293341390127e-3, 1e-14,
			tname, "rp13")
		vvd(t, rp[1][0], -0.8696112560510186244e-3, 1e-14,
			tname, "rp21")
		vvd(t, rp[1][1], 0.9999996218880458820, 1e-12,
			tname, "rp22")
		vvd(t, rp[1][2], -0.1691646168941896285e-6, 1e-14,
			tname, "rp23")
		vvd(t, rp[2][0], -0.3778929335557603418e-3, 1e-14,
			tname, "rp31")
		vvd(t, rp[2][1], -0.1594554040786495076e-6, 1e-14,
			tname, "rp32")
		vvd(t, rp[2][2], 0.9999999285984501222, 1e-12,
			tname, "rp33")

		vvd(t, rbp[0][0], 0.9999995505176007047, 1e-12,
			tname, "rbp11")
		vvd(t, rbp[0][1], 0.8695404617348208406e-3, 1e-14,
			tname, "rbp12")
		vvd(t, rbp[0][2], 0.3779735201865589104e-3, 1e-14,
			tname, "rbp13")
		vvd(t, rbp[1][0], -0.8695404723772031414e-3, 1e-14,
			tname, "rbp21")
		vvd(t, rbp[1][1], 0.9999996219496027161, 1e-12,
			tname, "rbp22")
		vvd(t, rbp[1][2], -0.1361752497080270143e-6, 1e-14,
			tname, "rbp23")
		vvd(t, rbp[2][0], -0.3779734957034089490e-3, 1e-14,
			tname, "rbp31")
		vvd(t, rbp[2][1], -0.1924880847894457113e-6, 1e-14,
			tname, "rbp32")
		vvd(t, rbp[2][2], 0.9999999285679971958, 1e-12,
			tname, "rbp33")
	}
}

func BenchmarkBp06(b *testing.B) {
}
