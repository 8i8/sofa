package sofa

import "testing"

//
//  - - - - - - - -
//   T e s t L d n
//  - - - - - - - -
//
//  Test Ldn function.
//
//  Called:  Ldn, vvd
//
//  This revision:  2013 October 2
//
func TestLdn(t *testing.T) {
	const fname = "Ldn"
	var n int
	var ob, sc [3]float64
	b := make([]LDBODY, 3)

	n = 3
	b[0].bm = 0.00028574
	b[0].dl = 3e-10
	b[0].pv[0][0] = -7.81014427
	b[0].pv[0][1] = -5.60956681
	b[0].pv[0][2] = -1.98079819
	b[0].pv[1][0] = 0.0030723249
	b[0].pv[1][1] = -0.00406995477
	b[0].pv[1][2] = -0.00181335842
	b[1].bm = 0.00095435
	b[1].dl = 3e-9
	b[1].pv[0][0] = 0.738098796
	b[1].pv[0][1] = 4.63658692
	b[1].pv[0][2] = 1.9693136
	b[1].pv[1][0] = -0.00755816922
	b[1].pv[1][1] = 0.00126913722
	b[1].pv[1][2] = 0.000727999001
	b[2].bm = 1.0
	b[2].dl = 6e-6
	b[2].pv[0][0] = -0.000712174377
	b[2].pv[0][1] = -0.00230478303
	b[2].pv[0][2] = -0.00105865966
	b[2].pv[1][0] = 6.29235213e-6
	b[2].pv[1][1] = -3.30888387e-7
	b[2].pv[1][2] = -2.96486623e-7
	ob[0] = -0.974170437
	ob[1] = -0.2115201
	ob[2] = -0.0917583114
	sc[0] = -0.763276255
	sc[1] = -0.608633767
	sc[2] = -0.216735543

	tests := []struct {
		ref string
		fn  func(int, []LDBODY, [3]float64, [3]float64) [3]float64
	}{
		{"cgo", CgoLdn},
		{"go", GoLdn},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		sn := test.fn(n, b, ob, sc)

		vvd(t, sn[0], -0.7632762579693333866, 1e-12,
			tname, "1")
		vvd(t, sn[1], -0.6086337636093002660, 1e-12,
			tname, "2")
		vvd(t, sn[2], -0.2167355420646328159, 1e-12,
			tname, "3")
	}
}

func Benchmark(b *testing.B) {
	var n int
	var ob, sc [3]float64
	body := make([]LDBODY, 3)

	n = 3
	body[0].bm = 0.00028574
	body[0].dl = 3e-10
	body[0].pv[0][0] = -7.81014427
	body[0].pv[0][1] = -5.60956681
	body[0].pv[0][2] = -1.98079819
	body[0].pv[1][0] = 0.0030723249
	body[0].pv[1][1] = -0.00406995477
	body[0].pv[1][2] = -0.00181335842
	body[1].bm = 0.00095435
	body[1].dl = 3e-9
	body[1].pv[0][0] = 0.738098796
	body[1].pv[0][1] = 4.63658692
	body[1].pv[0][2] = 1.9693136
	body[1].pv[1][0] = -0.00755816922
	body[1].pv[1][1] = 0.00126913722
	body[1].pv[1][2] = 0.000727999001
	body[2].bm = 1.0
	body[2].dl = 6e-6
	body[2].pv[0][0] = -0.000712174377
	body[2].pv[0][1] = -0.00230478303
	body[2].pv[0][2] = -0.00105865966
	body[2].pv[1][0] = 6.29235213e-6
	body[2].pv[1][1] = -3.30888387e-7
	body[2].pv[1][2] = -2.96486623e-7
	ob[0] = -0.974170437
	ob[1] = -0.2115201
	ob[2] = -0.0917583114
	sc[0] = -0.763276255
	sc[1] = -0.608633767
	sc[2] = -0.216735543

	tests := []struct {
		ref string
		fn  func(int, []LDBODY, [3]float64, [3]float64) [3]float64
	}{
		{"cgo", CgoLdn},
		{"go", GoLdn},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(n, body, ob, sc)
			}
		})
	}
}
