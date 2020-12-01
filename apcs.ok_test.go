package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A p c s
//  - - - - - - - - -
//
//  Test Apcs function.
//
//  Called:  Apcs, vvd
//
//  This revision:  2017 March 15
//
func TestApcs(t *testing.T) {
	const fname = "Apcs"

	var date1, date2 float64
	var pv, ebpv [2][3]float64
	var ehp [3]float64

	date1 = 2456384.5
	date2 = 0.970031644
	pv[0][0] = -1836024.09
	pv[0][1] = 1056607.72
	pv[0][2] = -5998795.26
	pv[1][0] = -77.0361767
	pv[1][1] = -133.310856
	pv[1][2] = 0.0971855934
	ebpv[0][0] = -0.974170438
	ebpv[0][1] = -0.211520082
	ebpv[0][2] = -0.0917583024
	ebpv[1][0] = 0.00364365824
	ebpv[1][1] = -0.0154287319
	ebpv[1][2] = -0.00668922024
	ehp[0] = -0.973458265
	ehp[1] = -0.209215307
	ehp[2] = -0.0906996477

	tests := []struct {
		ref string
		fn  func(a, b float64, c, d [2][3]float64, e [3]float64) ASTROM
	}{
		{"cgo", CgoApcs},
		{"go", GoApcs},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom := test.fn(date1, date2, pv, ebpv, ehp)

		vvd(t, astrom.pmt, 13.25248468622587269, 1e-11, tname, "pmt")
		vvd(t, astrom.eb[0], -0.9741827110629881886, 1e-12, tname, "eb(1)")
		vvd(t, astrom.eb[1], -0.2115130190136415986, 1e-12, tname, "eb(2)")
		vvd(t, astrom.eb[2], -0.09179840186954412099, 1e-12, tname, "eb(3)")
		vvd(t, astrom.eh[0], -0.9736425571689454706, 1e-12, tname, "eh(1)")
		vvd(t, astrom.eh[1], -0.2092452125850435930, 1e-12, tname, "eh(2)")
		vvd(t, astrom.eh[2], -0.09075578152248299218, 1e-12, tname, "eh(3)")
		vvd(t, astrom.em, 0.9998233241709796859, 1e-12, tname, "em")
		vvd(t, astrom.v[0], 0.2078704993282685510e-4, 1e-16, tname, "v(1)")
		vvd(t, astrom.v[1], -0.8955360106989405683e-4, 1e-16, tname, "v(2)")
		vvd(t, astrom.v[2], -0.3863338994289409097e-4, 1e-16, tname, "v(3)")
		vvd(t, astrom.bm1, 0.9999999950277561237, 1e-12, tname, "bm1")
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

func BenchmarkApcs(b *testing.B) {

	var date1, date2 float64
	var pv, ebpv [2][3]float64
	var ehp [3]float64

	date1 = 2456384.5
	date2 = 0.970031644
	pv[0][0] = -1836024.09
	pv[0][1] = 1056607.72
	pv[0][2] = -5998795.26
	pv[1][0] = -77.0361767
	pv[1][1] = -133.310856
	pv[1][2] = 0.0971855934
	ebpv[0][0] = -0.974170438
	ebpv[0][1] = -0.211520082
	ebpv[0][2] = -0.0917583024
	ebpv[1][0] = 0.00364365824
	ebpv[1][1] = -0.0154287319
	ebpv[1][2] = -0.00668922024
	ehp[0] = -0.973458265
	ehp[1] = -0.209215307
	ehp[2] = -0.0906996477

	tests := []struct {
		ref string
		fn  func(a, b float64, c, d [2][3]float64, e [3]float64) ASTROM
	}{
		{"cgo", CgoApcs},
		{"go", GoApcs},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(date1, date2, pv, ebpv, ehp)
			}
		})
	}
}
