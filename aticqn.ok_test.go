package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A t i c q n
//  - - - - - - - - - - -
//
//  Test Aticqn function.
//
//  Called:  Apci13, iauAticqn, vvd
//
//  This revision:  2017 March 15
//
func TestAticqn(t *testing.T) {
	const fname = "Aticqn"
	var date1, date2, ri, di, rc, dc float64
	var astrom ASTROM
	body := make([]LDBODY, 3)

	date1 = 2456165.5
	date2 = 0.401182685
	ri = 2.709994899247599271
	di = 0.1728740720983623469
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
		fn  func(a1, a2 float64, a3 ASTROM, a4 int, a5 []LDBODY) (
			c1, c2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAticqn, CgoApci13},
		{"go", GoAticqn, GoApci13},
	}

	for _, test := range tests {
		astrom, _ = test.fnAssist(date1, date2, astrom)
		tname := fname + " " + test.ref

		rc, dc = test.fn(ri, di, astrom, 3, body)

		vvd(t, rc, 2.709999575033027333, 1e-12, tname, "rc")
		vvd(t, dc, 0.1739999656316469990, 1e-12, tname, "dc")
	}
}

func BenchmarkAticqn(b *testing.B) {
	var date1, date2, ri, di float64
	var astrom ASTROM
	body := make([]LDBODY, 3)

	date1 = 2456165.5
	date2 = 0.401182685
	ri = 2.709994899247599271
	di = 0.1728740720983623469
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
		fn  func(a1, a2 float64, a3 ASTROM, a4 int, a5 []LDBODY) (
			c1, c2 float64)
		fnAssist func(a1, a2 float64, a3 ASTROM) (ASTROM, float64)
	}{
		{"cgo", CgoAticqn, CgoApci13},
		{"go", GoAticqn, GoApci13},
	}

	for _, test := range tests {
		astrom, _ = test.fnAssist(date1, date2, astrom)
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(ri, di, astrom, 3, body)
			}
		})
	}
}
