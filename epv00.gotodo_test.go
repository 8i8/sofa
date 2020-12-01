package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t E p v 0 0
//  - - - - - - - - - -
//
//  Test Epv00 function.
//
//  Called: Epv00, vvd
//
//  This revision:  2013 August 7
//
func TestEpv00(t *testing.T) {
	const fname = "Epv00"
	tests := []struct {
		ref string
		fn  func(a, b float64) (c, d [2][3]float64, f error)
	}{
		{"cgo", CgoEpv00},
		//{"go"}, GoEpv00},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref

		pvh, pvb, err := test.fn(2400000.5, 53411.52501161)
		if err != nil {
			t.Errorf("%s error: %s", tname, err)
		}

		vvd(t, pvh[0][0], -0.7757238809297706813, 1e-14, tname, "ph(x)")
		vvd(t, pvh[0][1], 0.5598052241363340596, 1e-14, tname, "ph(y)")
		vvd(t, pvh[0][2], 0.2426998466481686993, 1e-14, tname, "ph(z)")

		vvd(t, pvh[1][0], -0.1091891824147313846e-1, 1e-15, tname, "vh(x)")
		vvd(t, pvh[1][1], -0.1247187268440845008e-1, 1e-15, tname, "vh(y)")
		vvd(t, pvh[1][2], -0.5407569418065039061e-2, 1e-15, tname, "vh(z)")

		vvd(t, pvb[0][0], -0.7714104440491111971, 1e-14, tname, "pb(x)")
		vvd(t, pvb[0][1], 0.5598412061824171323, 1e-14, tname, "pb(y)")
		vvd(t, pvb[0][2], 0.2425996277722452400, 1e-14, tname, "pb(z)")

		vvd(t, pvb[1][0], -0.1091874268116823295e-1, 1e-15, tname, "vb(x)")
		vvd(t, pvb[1][1], -0.1246525461732861538e-1, 1e-15, tname, "vb(y)")
		vvd(t, pvb[1][2], -0.5404773180966231279e-2, 1e-15, tname, "vb(z)")
	}
}

func BenchmarkEpv00(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a, b float64) (c, d [2][3]float64, f error)
	}{
		{"cgo", CgoEpv00},
		//{"go"}, GoEpv00},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = test.fn(2400000.5, 53411.52501161)
			}
		})
	}
}
