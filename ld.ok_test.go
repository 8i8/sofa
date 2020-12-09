package sofa

import "testing"

//
//  - - - - - - -
//   T e s t L d
//  - - - - - - -
//
//  Test Ld function.
//
//  Called:  Ld, vvd
//
//  This revision:  2013 October 2
//
func TestLd(t *testing.T) {
	const fname = "Ld"
	var bm, em, dlim float64
	var p, q, e [3]float64

	bm = 0.00028574
	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	q[0] = -0.763276255
	q[1] = -0.608633767
	q[2] = -0.216735543
	e[0] = 0.76700421
	e[1] = 0.605629598
	e[2] = 0.211937094
	em = 8.91276983
	dlim = 3e-10

	tests := []struct {
		ref string
		fn  func(a1 float64, a2, a3, a4 [3]float64,
			a5, a6 float64) [3]float64
	}{
		{"cgo", CgoLd},
		{"go", GoLd},
	}
	for _, test := range tests {
		tname := fname + " " + test.ref
		p1 := test.fn(bm, p, q, e, em, dlim)

		vvd(t, p1[0], -0.7632762548968159627, 1e-12,
			tname, "1")
		vvd(t, p1[1], -0.6086337670823762701, 1e-12,
			tname, "2")
		vvd(t, p1[2], -0.2167355431320546947, 1e-12,
			tname, "3")
	}
}

func BenchmarkLd(b *testing.B) {
	var bm, em, dlim float64
	var p, q, e [3]float64

	bm = 0.00028574
	p[0] = -0.763276255
	p[1] = -0.608633767
	p[2] = -0.216735543
	q[0] = -0.763276255
	q[1] = -0.608633767
	q[2] = -0.216735543
	e[0] = 0.76700421
	e[1] = 0.605629598
	e[2] = 0.211937094
	em = 8.91276983
	dlim = 3e-10

	tests := []struct {
		ref string
		fn  func(a1 float64, a2, a3, a4 [3]float64,
			a5, a6 float64) [3]float64
	}{
		{"cgo", CgoLd},
		{"go", GoLd},
	}
	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(bm, p, q, e, em, dlim)
			}
		})
	}
}
