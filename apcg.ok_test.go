package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t A p c g
//  - - - - - - - - -
//
//  Test Apcg function.
//
//  Called:  Apcg, vvd
//
//  This revision:  2017 March 15
//
func TestApcg(t *testing.T) {
	const fname = "Apcg"

	var date1, date2 float64
	var ebpv [2][3]float64
	var ehp [3]float64
	var astr ASTROM

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

	tests := []struct {
		ref string
		fn  func(a, b float64, c [2][3]float64,
			d [3]float64, e ASTROM) ASTROM
	}{
		{"cgo", CgoApcg},
		{"go", GoApcg},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom := test.fn(date1, date2, ebpv, ehp, astr)

		vvd(t, astrom.pmt, 12.65133794027378508, 1e-11,
			tname, "pmt")

		vvd(t, astrom.eb[0], 0.901310875, 1e-12,
			tname, "eb(1)")
		vvd(t, astrom.eb[1], -0.417402664, 1e-12,
			tname, "eb(2)")
		vvd(t, astrom.eb[2], -0.180982288, 1e-12,
			tname, "eb(3)")

		vvd(t, astrom.eh[0], 0.8940025429324143045, 1e-12,
			tname, "eh(1)")
		vvd(t, astrom.eh[1], -0.4110930268679817955, 1e-12,
			tname, "eh(2)")
		vvd(t, astrom.eh[2], -0.1782189004872870264, 1e-12,
			tname, "eh(3)")

		vvd(t, astrom.em, 1.010465295811013146, 1e-12,
			tname, "em")

		vvd(t, astrom.v[0], 0.4289638913597693554e-4, 1e-16,
			tname, "v(1)")
		vvd(t, astrom.v[1], 0.8115034051581320575e-4, 1e-16,
			tname, "v(2)")
		vvd(t, astrom.v[2], 0.3517555136380563427e-4, 1e-16,
			tname, "v(3)")

		vvd(t, astrom.bm1, 0.9999999951686012981, 1e-12,
			tname, "bm1")

		vvd(t, astrom.bpn[0][0], 1.0, 0.0, tname, "bpn(1,1)")
		vvd(t, astrom.bpn[1][0], 0.0, 0.0, tname, "bpn(2,1)")
		vvd(t, astrom.bpn[2][0], 0.0, 0.0, tname, "bpn(3,1)")

		vvd(t, astrom.bpn[0][1], 0.0, 0.0, tname, "bpn(1,2)")
		vvd(t, astrom.bpn[1][1], 1.0, 0.0, tname, "bpn(2,2)")
		vvd(t, astrom.bpn[2][1], 0.0, 0.0, tname, "bpn(3,2)")

		vvd(t, astrom.bpn[0][2], 0.0, 0.0, tname, "bpn(1,3)")
		vvd(t, astrom.bpn[1][2], 0.0, 0.0, tname, "bpn(2,3)")
		vvd(t, astrom.bpn[2][2], 1.0, 0.0, tname, "bpn(3,3)")
	}
}

func BenchmarkApcg(b *testing.B) {
	var date1, date2 float64
	var ebpv [2][3]float64
	var ehp [3]float64
	var astr ASTROM

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

	tests := []struct {
		ref string
		fn  func(a, b float64, c [2][3]float64,
			d [3]float64, e ASTROM) ASTROM
	}{
		{"cgo", CgoApcg},
		{"go", GoApcg},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(date1, date2, ebpv, ehp, astr)
			}
		})
	}
}
