package sofa

import "testing"

/*
**  - - - - - - -
**   t _ a p c g
**  - - - - - - -
**
**  Test Apcg function.
**
**  Returned:
**     status    int         FALSE = success, TRUE = fail
**
**  Called:  iauApcg, vvd
**
**  This revision:  2017 March 15
 */
func TestApcg(t *testing.T) {
	const fname = "Apcg"

	var (
		date1, date2 float64
		ebpv         [2][3]float64
		ehp          [3]float64
	)

	date1 = 2456165.5
	date2 = 0.401182685
	ebpv[0][0] = 0.901310875
	ebpv[0][1] = -0.417402664
	ebpv[0][2] = -0.180982288
	ebpv[1][0] = 0.00742727954
	ebpv[1][1] = 0.0140507459
	ebpv[1][2] = 0.00609045792
	ehp[0] = 0.903358544
	ehp[1] = -0.415395237
	ehp[2] = -0.180084014

	astrom := Apcg(date1, date2, ebpv, ehp)

	vvd(t, astrom.pmt, 12.65133794027378508, 1e-11, fname, "pmt")
	vvd(t, astrom.eb[0], 0.901310875, 1e-12, fname, "eb(1)")
	vvd(t, astrom.eb[1], -0.417402664, 1e-12, fname, "eb(2)")
	vvd(t, astrom.eb[2], -0.180982288, 1e-12, fname, "eb(3)")
	vvd(t, astrom.eh[0], 0.8940025429324143045, 1e-12, fname, "eh(1)")
	vvd(t, astrom.eh[1], -0.4110930268679817955, 1e-12, fname, "eh(2)")
	vvd(t, astrom.eh[2], -0.1782189004872870264, 1e-12, fname, "eh(3)")
	vvd(t, astrom.em, 1.010465295811013146, 1e-12, fname, "em")
	vvd(t, astrom.v[0], 0.4289638913597693554e-4, 1e-16, fname, "v(1)")
	vvd(t, astrom.v[1], 0.8115034051581320575e-4, 1e-16, fname, "v(2)")
	vvd(t, astrom.v[2], 0.3517555136380563427e-4, 1e-16, fname, "v(3)")
	vvd(t, astrom.bm1, 0.9999999951686012981, 1e-12, fname, "bm1")
	vvd(t, astrom.bpn[0][0], 1.0, 0.0, fname, "bpn(1,1)")
	vvd(t, astrom.bpn[1][0], 0.0, 0.0, fname, "bpn(2,1)")
	vvd(t, astrom.bpn[2][0], 0.0, 0.0, fname, "bpn(3,1)")
	vvd(t, astrom.bpn[0][1], 0.0, 0.0, fname, "bpn(1,2)")
	vvd(t, astrom.bpn[1][1], 1.0, 0.0, fname, "bpn(2,2)")
	vvd(t, astrom.bpn[2][1], 0.0, 0.0, fname, "bpn(3,2)")
	vvd(t, astrom.bpn[0][2], 0.0, 0.0, fname, "bpn(1,3)")
	vvd(t, astrom.bpn[1][2], 0.0, 0.0, fname, "bpn(2,3)")
	vvd(t, astrom.bpn[2][2], 1.0, 0.0, fname, "bpn(3,3)")
}
