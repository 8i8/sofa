package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t P n 0 6 a
//  - - - - - - - - - -
//
//  Test Pn06a function.
//
//  Called:  Pn06a, vvd
//
//  This revision:  2013 August 7
//
func TestPn06a(t *testing.T) {
	const fname = "Pn06a"
	var dpsi, deps, epsa float64
	var rb, rp, rbp, rn, rbpn [3][3]float64

	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2, c3 float64,
			c4, c5, c6, c7, c8 [3][3]float64)
	}{
		{"cgo", CgoPn06a},
		{"go", GoPn06a},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dpsi, deps, epsa, rb, rp, rbp, rn, rbpn = test.fn(
			2400000.5, 53736.0)

		vvd(t, dpsi, -0.9630912025820308797e-5, 1e-12,
			tname, "dpsi")
		vvd(t, deps, 0.4063238496887249798e-4, 1e-12,
			tname, "deps")
		vvd(t, epsa, 0.4090789763356509926, 1e-12,
			tname, "epsa")

		vvd(t, rb[0][0], 0.9999999999999942497, 1e-12,
			tname, "rb11")
		vvd(t, rb[0][1], -0.7078368960971557145e-7, 1e-14,
			tname, "rb12")
		vvd(t, rb[0][2], 0.8056213977613185606e-7, 1e-14,
			tname, "rb13")

		vvd(t, rb[1][0], 0.7078368694637674333e-7, 1e-14,
			tname, "rb21")
		vvd(t, rb[1][1], 0.9999999999999969484, 1e-12,
			tname, "rb22")
		vvd(t, rb[1][2], 0.3305943742989134124e-7, 1e-14,
			tname, "rb23")

		vvd(t, rb[2][0], -0.8056214211620056792e-7, 1e-14,
			tname, "rb31")
		vvd(t, rb[2][1], -0.3305943172740586950e-7, 1e-14,
			tname, "rb32")
		vvd(t, rb[2][2], 0.9999999999999962084, 1e-12,
			tname, "rb33")

		vvd(t, rp[0][0], 0.9999989300536854831, 1e-12,
			tname, "rp11")
		vvd(t, rp[0][1], -0.1341646886204443795e-2, 1e-14,
			tname, "rp12")
		vvd(t, rp[0][2], -0.5829880933488627759e-3, 1e-14,
			tname, "rp13")

		vvd(t, rp[1][0], 0.1341646890569782183e-2, 1e-14,
			tname, "rp21")
		vvd(t, rp[1][1], 0.9999990999913319321, 1e-12,
			tname, "rp22")
		vvd(t, rp[1][2], -0.3835944216374477457e-6, 1e-14,
			tname, "rp23")

		vvd(t, rp[2][0], 0.5829880833027867368e-3, 1e-14,
			tname, "rp31")
		vvd(t, rp[2][1], -0.3985701514686976112e-6, 1e-14,
			tname, "rp32")
		vvd(t, rp[2][2], 0.9999998300623534950, 1e-12,
			tname, "rp33")

		vvd(t, rbp[0][0], 0.9999989300056797893, 1e-12,
			tname, "rbp11")
		vvd(t, rbp[0][1], -0.1341717650545059598e-2, 1e-14,
			tname, "rbp12")
		vvd(t, rbp[0][2], -0.5829075756493728856e-3, 1e-14,
			tname, "rbp13")

		vvd(t, rbp[1][0], 0.1341717674223918101e-2, 1e-14,
			tname, "rbp21")
		vvd(t, rbp[1][1], 0.9999990998963748448, 1e-12,
			tname, "rbp22")
		vvd(t, rbp[1][2], -0.3504269280170069029e-6, 1e-14,
			tname, "rbp23")

		vvd(t, rbp[2][0], 0.5829075211461454599e-3, 1e-14,
			tname, "rbp31")
		vvd(t, rbp[2][1], -0.4316708436255949093e-6, 1e-14,
			tname, "rbp32")
		vvd(t, rbp[2][2], 0.9999998301093032943, 1e-12,
			tname, "rbp33")

		vvd(t, rn[0][0], 0.9999999999536227668, 1e-12,
			tname, "rn11")
		vvd(t, rn[0][1], 0.8836241998111535233e-5, 1e-14,
			tname, "rn12")
		vvd(t, rn[0][2], 0.3830834608415287707e-5, 1e-14,
			tname, "rn13")

		vvd(t, rn[1][0], -0.8836086334870740138e-5, 1e-14,
			tname, "rn21")
		vvd(t, rn[1][1], 0.9999999991354657474, 1e-12,
			tname, "rn22")
		vvd(t, rn[1][2], -0.4063240188248455065e-4, 1e-14,
			tname, "rn23")

		vvd(t, rn[2][0], -0.3831193642839398128e-5, 1e-14,
			tname, "rn31")
		vvd(t, rn[2][1], 0.4063236803101479770e-4, 1e-14,
			tname, "rn32")
		vvd(t, rn[2][2], 0.9999999991671663114, 1e-12,
			tname, "rn33")

		vvd(t, rbpn[0][0], 0.9999989440480669738, 1e-12,
			tname, "rbpn11")
		vvd(t, rbpn[0][1], -0.1332881418091915973e-2, 1e-14,
			tname, "rbpn12")
		vvd(t, rbpn[0][2], -0.5790767447612042565e-3, 1e-14,
			tname, "rbpn13")

		vvd(t, rbpn[1][0], 0.1332857911250989133e-2, 1e-14,
			tname, "rbpn21")
		vvd(t, rbpn[1][1], 0.9999991109049141908, 1e-12,
			tname, "rbpn22")
		vvd(t, rbpn[1][2], -0.4097767128546784878e-4, 1e-14,
			tname, "rbpn23")

		vvd(t, rbpn[2][0], 0.5791308482835292617e-3, 1e-14,
			tname, "rbpn31")
		vvd(t, rbpn[2][1], 0.4020580099454020310e-4, 1e-14,
			tname, "rbpn32")
		vvd(t, rbpn[2][2], 0.9999998314954628695, 1e-12,
			tname, "rbpn33")
	}
}

func BenchmarkPn06a(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64) (c1, c2, c3 float64,
			c4, c5, c6, c7, c8 [3][3]float64)
	}{
		{"cgo", CgoPn06a},
		{"go", GoPn06a},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2400000.5, 53736.0)
			}
		})
	}
}
