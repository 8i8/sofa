package sofa

import (
	"testing"
)

//
//  - - - - - - - - - - -
//   T e s t A p c g 1 3
//  - - - - - - - - - - -
//
//  Test Apcg13 function.
//
//  Called:  Apcg13, vvd
//
//  This revision:  2017 March 15
//
func TestApcg13(t *testing.T) {
	const fname = "Apcg13"
	date1 := 2456165.5
	date2 := 0.401182685

	tests := []struct {
		ref string
		fn  func(a, b float64) ASTROM
	}{
		{"cgo", Apcg13},
		{"go", goApcg13},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		astrom := test.fn(date1, date2)
		vvd(t, astrom.pmt, 12.65133794027378508, 1e-11, tname, "pmt")
		vvd(t, astrom.eb[0], 0.9013108747340644755, 1e-12, tname, "eb(1)")
		vvd(t, astrom.eb[1], -0.4174026640406119957, 1e-12, tname, "eb(2)")
		vvd(t, astrom.eb[2], -0.1809822877867817771, 1e-12, tname, "eb(3)")
		vvd(t, astrom.eh[0], 0.8940025429255499549, 1e-12, tname, "eh(1)")
		vvd(t, astrom.eh[1], -0.4110930268331896318, 1e-12, tname, "eh(2)")
		vvd(t, astrom.eh[2], -0.1782189006019749850, 1e-12, tname, "eh(3)")
		vvd(t, astrom.em, 1.010465295964664178, 1e-12, tname, "em")
		vvd(t, astrom.v[0], 0.4289638912941341125e-4, 1e-16, tname, "v(1)")
		vvd(t, astrom.v[1], 0.8115034032405042132e-4, 1e-16, tname, "v(2)")
		vvd(t, astrom.v[2], 0.3517555135536470279e-4, 1e-16, tname, "v(3)")
		vvd(t, astrom.bm1, 0.9999999951686013142, 1e-12, tname, "bm1")
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

func BenchmarkApcg13(b *testing.B) {
	date1 := 2456165.5
	date2 := 0.401182685

	tests := []struct {
		ref string
		fn  func(a, b float64) ASTROM
	}{
		{"cgo", Apcg13},
		{"go", goApcg13},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(date1, date2)
			}
		})
	}
}
