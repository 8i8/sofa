package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A t c i q n
//  - - - - - - - - - - -
//
//  Test Atciqn function.
//
//  Called:  Apci13, iauAtciqn, vvd
//
//  This revision:  2017 March 15
//
func TestAtciqn(t *testing.T) {
	const fname = "Atciqn"
	var date1, date2, rc, dc, pr, pd, px, rv float64
	var astrom ASTROM
	b := make([]LDBODY, 3)

	date1 = 2456165.5
	date2 = 0.401182685
	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0
	b[0].bm = 0.00028574
	b[0].dl = 3e-10
	b[0].pv[0][0] = -7.81014427
	b[0].pv[0][1] = -5.60956681
	b[0].pv[0][2] = -1.98079819
	b[0].pv[1][0] = 0.0030723249
	b[0].pv[1][1] = -0.00406995477
	b[0].pv[1][2] = -0.00181335842
	b[1].bm = 0.00095435
	b[1].dl = 3e-9
	b[1].pv[0][0] = 0.738098796
	b[1].pv[0][1] = 4.63658692
	b[1].pv[0][2] = 1.9693136
	b[1].pv[1][0] = -0.00755816922
	b[1].pv[1][1] = 0.00126913722
	b[1].pv[1][2] = 0.000727999001
	b[2].bm = 1.0
	b[2].dl = 6e-6
	b[2].pv[0][0] = -0.000712174377
	b[2].pv[0][1] = -0.00230478303
	b[2].pv[0][2] = -0.00105865966
	b[2].pv[1][0] = 6.29235213e-6
	b[2].pv[1][1] = -3.30888387e-7
	b[2].pv[1][2] = -2.96486623e-7

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64, a7 ASTROM,
			a8 int, a9 []LDBODY) (a10, a11 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAtciqn, CgoApci13},
		{"go", GoAtciqn, GoApci13},
	}

	for _, test := range tests {
		astrom, _ := test.fnAssist(date1, date2, astrom)
		tname := fname + " " + test.ref

		ri, di := test.fn(rc, dc, pr, pd, px, rv, astrom, 3, b)

		vvd(t, ri, 2.710122008104983335, 1e-12, tname, "ri")
		vvd(t, di, 0.1729371916492767821, 1e-12, tname, "di")
	}
}

func BenchmarkAtciqn(b *testing.B) {
	var date1, date2, rc, dc, pr, pd, px, rv float64
	var astrom ASTROM
	body := make([]LDBODY, 3)

	date1 = 2456165.5
	date2 = 0.401182685
	rc = 2.71
	dc = 0.174
	pr = 1e-5
	pd = 5e-6
	px = 0.1
	rv = 55.0
	body[0].bm = 0.00028574
	body[0].dl = 3e-10
	body[0].pv[0][0] = -7.81014427
	body[0].pv[0][1] = -5.60956681
	body[0].pv[0][2] = -1.98079819
	body[0].pv[1][0] = 0.0030723249
	body[0].pv[1][1] = -0.00406995477
	body[0].pv[1][2] = -0.00181335842
	body[1].bm = 0.00095435
	body[1].dl = 3e-9
	body[1].pv[0][0] = 0.738098796
	body[1].pv[0][1] = 4.63658692
	body[1].pv[0][2] = 1.9693136
	body[1].pv[1][0] = -0.00755816922
	body[1].pv[1][1] = 0.00126913722
	body[1].pv[1][2] = 0.000727999001
	body[2].bm = 1.0
	body[2].dl = 6e-6
	body[2].pv[0][0] = -0.000712174377
	body[2].pv[0][1] = -0.00230478303
	body[2].pv[0][2] = -0.00105865966
	body[2].pv[1][0] = 6.29235213e-6
	body[2].pv[1][1] = -3.30888387e-7
	body[2].pv[1][2] = -2.96486623e-7

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4, a5, a6 float64, a7 ASTROM,
			a8 int, a9 []LDBODY) (a10, a11 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAtciqn, CgoApci13},
		{"go", GoAtciqn, GoApci13},
	}

	for _, test := range tests {
		astrom, _ := test.fnAssist(date1, date2, astrom)
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(rc, dc, pr, pd, px, rv,
					astrom, 3, body)
			}
		})
	}
}
