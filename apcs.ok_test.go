package sofa

import "testing"

//
//  - - - - - - -
//   t _ a p c s
//  - - - - - - -
//
//  Test Apcs function.
//
//  Returned:
//     status    int         FALSE = success, TRUE = fail
//
//  Called:  Apcs, vvd
//
//  This revision:  2017 March 15
//
func TestApcs(t *testing.T) {
	const fname = "Apcs"

	var (
		date1, date2 float64
		pv, ebpv     [2][3]float64
		ehp          [3]float64
	)

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

	astrom := Apcs(date1, date2, pv, ebpv, ehp)

	vvd(t, astrom.pmt, 13.25248468622587269, 1e-11, fname, "pmt")
	vvd(t, astrom.eb[0], -0.9741827110629881886, 1e-12, fname, "eb(1)")
	vvd(t, astrom.eb[1], -0.2115130190136415986, 1e-12, fname, "eb(2)")
	vvd(t, astrom.eb[2], -0.09179840186954412099, 1e-12, fname, "eb(3)")
	vvd(t, astrom.eh[0], -0.9736425571689454706, 1e-12, fname, "eh(1)")
	vvd(t, astrom.eh[1], -0.2092452125850435930, 1e-12, fname, "eh(2)")
	vvd(t, astrom.eh[2], -0.09075578152248299218, 1e-12, fname, "eh(3)")
	vvd(t, astrom.em, 0.9998233241709796859, 1e-12, fname, "em")
	vvd(t, astrom.v[0], 0.2078704993282685510e-4, 1e-16, fname, "v(1)")
	vvd(t, astrom.v[1], -0.8955360106989405683e-4, 1e-16, fname, "v(2)")
	vvd(t, astrom.v[2], -0.3863338994289409097e-4, 1e-16, fname, "v(3)")
	vvd(t, astrom.bm1, 0.9999999950277561237, 1e-12, fname, "bm1")
	vvd(t, astrom.bpn[0][0], 1, 0, fname, "bpn(1,1)")
	vvd(t, astrom.bpn[1][0], 0, 0, fname, "bpn(2,1)")
	vvd(t, astrom.bpn[2][0], 0, 0, fname, "bpn(3,1)")
	vvd(t, astrom.bpn[0][1], 0, 0, fname, "bpn(1,2)")
	vvd(t, astrom.bpn[1][1], 1, 0, fname, "bpn(2,2)")
	vvd(t, astrom.bpn[2][1], 0, 0, fname, "bpn(3,2)")
	vvd(t, astrom.bpn[0][2], 0, 0, fname, "bpn(1,3)")
	vvd(t, astrom.bpn[1][2], 0, 0, fname, "bpn(2,3)")
	vvd(t, astrom.bpn[2][2], 1, 0, fname, "bpn(3,3)")
}
