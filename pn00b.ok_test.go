package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P n 0 0 b
//  - - - - - - - - - -
//
//  Test Pn00b function.
//
//  Called:  Pn00b, vvd
//
//  This revision:  2013 August 7
//
func TestPn00b(t *testing.T) {
	const fname = "Pn00b"
	var dpsi, deps, epsa float64
	var rb, rp, rbp, rn, rbpn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (
			c1, c2, c3 float64,
			c4, c5, c6, c7, c8 [3][3]float64)
	}{
		{"cgo", CgoPn00b},
		{"go", GoPn00b},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dpsi, deps, epsa, rb, rp, rbp, rn, rbpn = test.fn(
			2400000.5, 53736.0)

		vvd(t, dpsi, -0.9632552291148362783e-5, 1e-12,
			tname, "dpsi")
		vvd(t, deps, 0.4063197106621159367e-4, 1e-12,
			tname, "deps")
		vvd(t, epsa, 0.4090791789404229916, 1e-12,
			tname, "epsa")

		vvd(t, rb[0][0], 0.9999999999999942498, 1e-12,
			tname, "rb11")
		vvd(t, rb[0][1], -0.7078279744199196626e-7, 1e-16,
			tname, "rb12")
		vvd(t, rb[0][2], 0.8056217146976134152e-7, 1e-16,
			tname, "rb13")

		vvd(t, rb[1][0], 0.7078279477857337206e-7, 1e-16,
			tname, "rb21")
		vvd(t, rb[1][1], 0.9999999999999969484, 1e-12,
			tname, "rb22")
		vvd(t, rb[1][2], 0.3306041454222136517e-7, 1e-16,
			tname, "rb23")

		vvd(t, rb[2][0], -0.8056217380986972157e-7, 1e-16,
			tname, "rb31")
		vvd(t, rb[2][1], -0.3306040883980552500e-7, 1e-16,
			tname, "rb32")
		vvd(t, rb[2][2], 0.9999999999999962084, 1e-12,
			tname, "rb33")

		vvd(t, rp[0][0], 0.9999989300532289018, 1e-12,
			tname, "rp11")
		vvd(t, rp[0][1], -0.1341647226791824349e-2, 1e-14,
			tname, "rp12")
		vvd(t, rp[0][2], -0.5829880927190296547e-3, 1e-14,
			tname, "rp13")

		vvd(t, rp[1][0], 0.1341647231069759008e-2, 1e-14,
			tname, "rp21")
		vvd(t, rp[1][1], 0.9999990999908750433, 1e-12,
			tname, "rp22")
		vvd(t, rp[1][2], -0.3837444441583715468e-6, 1e-14,
			tname, "rp23")

		vvd(t, rp[2][0], 0.5829880828740957684e-3, 1e-14,
			tname, "rp31")
		vvd(t, rp[2][1], -0.3984203267708834759e-6, 1e-14,
			tname, "rp32")
		vvd(t, rp[2][2], 0.9999998300623538046, 1e-12,
			tname, "rp33")

		vvd(t, rbp[0][0], 0.9999989300052243993, 1e-12,
			tname, "rbp11")
		vvd(t, rbp[0][1], -0.1341717990239703727e-2, 1e-14,
			tname, "rbp12")
		vvd(t, rbp[0][2], -0.5829075749891684053e-3, 1e-14,
			tname, "rbp13")

		vvd(t, rbp[1][0], 0.1341718013831739992e-2, 1e-14,
			tname, "rbp21")
		vvd(t, rbp[1][1], 0.9999990998959191343, 1e-12,
			tname, "rbp22")
		vvd(t, rbp[1][2], -0.3505759733565421170e-6, 1e-14,
			tname, "rbp23")

		vvd(t, rbp[2][0], 0.5829075206857717883e-3, 1e-14,
			tname, "rbp31")
		vvd(t, rbp[2][1], -0.4315219955198608970e-6, 1e-14,
			tname, "rbp32")
		vvd(t, rbp[2][2], 0.9999998301093036269, 1e-12,
			tname, "rbp33")

		vvd(t, rn[0][0], 0.9999999999536069682, 1e-12,
			tname, "rn11")
		vvd(t, rn[0][1], 0.8837746144871248011e-5, 1e-14,
			tname, "rn12")
		vvd(t, rn[0][2], 0.3831488838252202945e-5, 1e-14,
			tname, "rn13")

		vvd(t, rn[1][0], -0.8837590456632304720e-5, 1e-14,
			tname, "rn21")
		vvd(t, rn[1][1], 0.9999999991354692733, 1e-12,
			tname, "rn22")
		vvd(t, rn[1][2], -0.4063198798559591654e-4, 1e-14,
			tname, "rn23")

		vvd(t, rn[2][0], -0.3831847930134941271e-5, 1e-14,
			tname, "rn31")
		vvd(t, rn[2][1], 0.4063195412258168380e-4, 1e-14,
			tname, "rn32")
		vvd(t, rn[2][2], 0.9999999991671806225, 1e-12,
			tname, "rn33")

		vvd(t, rbpn[0][0], 0.9999989440499982806, 1e-12,
			tname, "rbpn11")
		vvd(t, rbpn[0][1], -0.1332880253640849194e-2, 1e-14,
			tname, "rbpn12")
		vvd(t, rbpn[0][2], -0.5790760898731091166e-3, 1e-14,
			tname, "rbpn13")

		vvd(t, rbpn[1][0], 0.1332856746979949638e-2, 1e-14,
			tname, "rbpn21")
		vvd(t, rbpn[1][1], 0.9999991109064768883, 1e-12,
			tname, "rbpn22")
		vvd(t, rbpn[1][2], -0.4097740555723081811e-4, 1e-14,
			tname, "rbpn23")

		vvd(t, rbpn[2][0], 0.5791301929950208873e-3, 1e-14,
			tname, "rbpn31")
		vvd(t, rbpn[2][1], 0.4020553681373720832e-4, 1e-14,
			tname, "rbpn32")
		vvd(t, rbpn[2][2], 0.9999998314958529887, 1e-12,
			tname, "rbpn33")
	}
}

func BenchmarkPn00b(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (
			c1, c2, c3 float64,
			c4, c5, c6, c7, c8 [3][3]float64)
	}{
		{"cgo", CgoPn00b},
		{"go", GoPn00b},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2400000.5, 53736.0)
			}
		})
	}
}
