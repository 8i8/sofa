package sofa

import (
	"log"
	"testing"
)

func TestApcg13(t *testing.T) {
	const fname = "Apcg13"
	date1 := 2456165.5
	date2 := 0.401182685

	astrom, err := Apcg13(date1, date2)
	if err != nil {
		t.Errorf("%s failed: error %s", fname, err)
	} else if verbose {
		log.Printf("%s passed: error %s", fname, err)
	}
	vvd(t, astrom.pmt, 12.65133794027378508, 1e-11, fname, "pmt")
	vvd(t, astrom.eb[0], 0.9013108747340644755, 1e-12, fname, "eb(1)")
	vvd(t, astrom.eb[1], -0.4174026640406119957, 1e-12, fname, "eb(2)")
	vvd(t, astrom.eb[2], -0.1809822877867817771, 1e-12, fname, "eb(3)")
	vvd(t, astrom.eh[0], 0.8940025429255499549, 1e-12, fname, "eh(1)")
	vvd(t, astrom.eh[1], -0.4110930268331896318, 1e-12, fname, "eh(2)")
	vvd(t, astrom.eh[2], -0.1782189006019749850, 1e-12, fname, "eh(3)")
	vvd(t, astrom.em, 1.010465295964664178, 1e-12, fname, "em")
	vvd(t, astrom.v[0], 0.4289638912941341125e-4, 1e-16, fname, "v(1)")
	vvd(t, astrom.v[1], 0.8115034032405042132e-4, 1e-16, fname, "v(2)")
	vvd(t, astrom.v[2], 0.3517555135536470279e-4, 1e-16, fname, "v(3)")
	vvd(t, astrom.bm1, 0.9999999951686013142, 1e-12, fname, "bm1")
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
