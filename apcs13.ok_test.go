package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t A p c s 1 3
//  - - - - - - - - - - -
//
//  Test Apcs13 function.
//
//  Called:  Apcs13, vvd
//
//  This revision:  2017 March 15
//
func TestApcs13(t *testing.T) {
	const fname = "Apcs13"
	var date1, date2 float64
	var pv [2][3]float64
	var astr ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	pv[0][0] = -6241497.16
	pv[0][1] = 401346.896
	pv[0][2] = -1251136.04
	pv[1][0] = -29.264597
	pv[1][1] = -455.021831
	pv[1][2] = 0.0266151194

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [2][3]float64,
			a4 ASTROM) ASTROM
	}{
		{"cgo", CgoApcs13},
		{"go", GoApcs13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom := test.fn(date1, date2, pv, astr)

		vvd(t, astrom.pmt, 12.65133794027378508, 1e-11,
			tname, "pmt")

		vvd(t, astrom.eb[0], 0.9012691529025250644, 1e-12,
			tname, "eb(1)")
		vvd(t, astrom.eb[1], -0.4173999812023194317, 1e-12,
			tname, "eb(2)")
		vvd(t, astrom.eb[2], -0.1809906511146429670, 1e-12,
			tname, "eb(3)")

		vvd(t, astrom.eh[0], 0.8939939101760130792, 1e-12,
			tname, "eh(1)")
		vvd(t, astrom.eh[1], -0.4111053891734021478, 1e-12,
			tname, "eh(2)")
		vvd(t, astrom.eh[2], -0.1782336880636997374, 1e-12,
			tname, "eh(3)")

		vvd(t, astrom.em, 1.010428384373491095, 1e-12,
			tname, "em")

		vvd(t, astrom.v[0], 0.4279877294121697570e-4, 1e-16,
			tname, "v(1)")
		vvd(t, astrom.v[1], 0.7963255087052120678e-4, 1e-16,
			tname, "v(2)")
		vvd(t, astrom.v[2], 0.3517564013384691531e-4, 1e-16,
			tname, "v(3)")

		vvd(t, astrom.bm1, 0.9999999952947980978, 1e-12,
			tname, "bm1")

		vvd(t, astrom.bpn[0][0], 1, 0, tname, "bpn(1,1)")
		vvd(t, astrom.bpn[1][0], 0, 0, tname, "bpn(2,1)")
		vvd(t, astrom.bpn[2][0], 0, 0, tname, "bpn(3,1)")
		vvd(t, astrom.bpn[0][1], 0, 0, tname, "bpn(1,2)")
		vvd(t, astrom.bpn[1][1], 1, 0, tname, "bpn(2,2)")
		vvd(t, astrom.bpn[2][1], 0, 0, tname, "bpn(3,2)")
		vvd(t, astrom.bpn[0][2], 0, 0, tname, "bpn(1,3)")
		vvd(t, astrom.bpn[1][2], 0, 0, tname, "bpn(2,3)")
		vvd(t, astrom.bpn[2][2], 1, 0, tname, "bpn(3,3)")
	}
}

func BenchmarkApcs13(b *testing.B) {
	var date1, date2 float64
	var pv [2][3]float64
	var astr ASTROM

	date1 = 2456165.5
	date2 = 0.401182685
	pv[0][0] = -6241497.16
	pv[0][1] = 401346.896
	pv[0][2] = -1251136.04
	pv[1][0] = -29.264597
	pv[1][1] = -455.021831
	pv[1][2] = 0.0266151194

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 [2][3]float64,
			a4 ASTROM) ASTROM
	}{
		{"cgo", CgoApcs13},
		{"go", GoApcs13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(date1, date2, pv, astr)
			}
		})
	}
}
