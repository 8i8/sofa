package sofa

import "testing"

//
//  - - - - - - - -
//   t _ e p v 0 0
//  - - - - - - - -
//
//  Test Epv00 function.
//
//  Returned:
//     status    int         FALSE = success, TRUE = fail
//
//  Called: iauEpv00, vvd
//
//  This revision:  2013 August 7
//
func TestEpv00(t *testing.T) {
	const fname = "Epv00"
	pvh, pvb, err := Epv00(2400000.5, 53411.52501161)
	if err != nil {
		t.Errorf("%s error: %s", fname, err)
	}
	vvd(t, pvh[0][0], -0.7757238809297706813, 1e-14, fname, "ph(x)")
	vvd(t, pvh[0][1], 0.5598052241363340596, 1e-14, fname, "ph(y)")
	vvd(t, pvh[0][2], 0.2426998466481686993, 1e-14, fname, "ph(z)")

	vvd(t, pvh[1][0], -0.1091891824147313846e-1, 1e-15, fname, "vh(x)")
	vvd(t, pvh[1][1], -0.1247187268440845008e-1, 1e-15, fname, "vh(y)")
	vvd(t, pvh[1][2], -0.5407569418065039061e-2, 1e-15, fname, "vh(z)")

	vvd(t, pvb[0][0], -0.7714104440491111971, 1e-14, fname, "pb(x)")
	vvd(t, pvb[0][1], 0.5598412061824171323, 1e-14, fname, "pb(y)")
	vvd(t, pvb[0][2], 0.2425996277722452400, 1e-14, fname, "pb(z)")

	vvd(t, pvb[1][0], -0.1091874268116823295e-1, 1e-15, fname, "vb(x)")
	vvd(t, pvb[1][1], -0.1246525461732861538e-1, 1e-15, fname, "vb(y)")
	vvd(t, pvb[1][2], -0.5404773180966231279e-2, 1e-15, fname, "vb(z)")
}

// static void t_epv00(int *status)
// /*
// **  - - - - - - - -
// **   t _ e p v 0 0
// **  - - - - - - - -
// **
// **  Test iauEpv00 function.
// **
// **  Returned:
// **     status    int         FALSE = success, TRUE = fail
// **
// **  Called: iauEpv00, vvd, viv
// **
// **  This revision:  2013 August 7
// */
// {
//    double pvh[2][3], pvb[2][3];
//    int j;

//    j = iauEpv00(2400000.5, 53411.52501161, pvh, pvb);

//    vvd(pvh[0][0], -0.7757238809297706813, 1e-14,
//        "iauEpv00", "ph(x)", status);
//    vvd(pvh[0][1], 0.5598052241363340596, 1e-14,
//        "iauEpv00", "ph(y)", status);
//    vvd(pvh[0][2], 0.2426998466481686993, 1e-14,
//        "iauEpv00", "ph(z)", status);

//    vvd(pvh[1][0], -0.1091891824147313846e-1, 1e-15,
//        "iauEpv00", "vh(x)", status);
//    vvd(pvh[1][1], -0.1247187268440845008e-1, 1e-15,
//        "iauEpv00", "vh(y)", status);
//    vvd(pvh[1][2], -0.5407569418065039061e-2, 1e-15,
//        "iauEpv00", "vh(z)", status);

//    vvd(pvb[0][0], -0.7714104440491111971, 1e-14,
//        "iauEpv00", "pb(x)", status);
//    vvd(pvb[0][1], 0.5598412061824171323, 1e-14,
//        "iauEpv00", "pb(y)", status);
//    vvd(pvb[0][2], 0.2425996277722452400, 1e-14,
//        "iauEpv00", "pb(z)", status);

//    vvd(pvb[1][0], -0.1091874268116823295e-1, 1e-15,
//        "iauEpv00", "vb(x)", status);
//    vvd(pvb[1][1], -0.1246525461732861538e-1, 1e-15,
//        "iauEpv00", "vb(y)", status);
//    vvd(pvb[1][2], -0.5404773180966231279e-2, 1e-15,
//        "iauEpv00", "vb(z)", status);

//    viv(j, 0, "iauEpv00", "j", status);

// }
