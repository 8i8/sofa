package sofa

import (
	"testing"

	"github.com/8i8/sofa/en"
)

//
//  - - - - - - - - - - -
//   T e s t P l a n 9 4
//  - - - - - - - - - - -
//
//  Test Plan94 function.
//
//  Called:  Plan94, vvd, viv
//
//  This revision:  2013 October 2
//
func TestPlan94(t *testing.T) {
	const fname = "Plan94"
	var pv [2][3]float64
	var err en.ErrNum

	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 int) (
			[2][3]float64, en.ErrNum)
	}{
		{"cgo", CgoPlan94},
		{"go", GoPlan94},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		pv, err = test.fn(2400000.5, 1e6, 0)

		vvd(t, pv[0][0], 0.0, 0.0, tname, "x 1")
		vvd(t, pv[0][1], 0.0, 0.0, tname, "y 1")
		vvd(t, pv[0][2], 0.0, 0.0, tname, "z 1")

		vvd(t, pv[1][0], 0.0, 0.0, tname, "xd 1")
		vvd(t, pv[1][1], 0.0, 0.0, tname, "yd 1")
		vvd(t, pv[1][2], 0.0, 0.0, tname, "zd 1")

		errEN(t, -1, err, tname, "j 1")

		pv, err = test.fn(2400000.5, 1e6, 10)

		errEN(t, -1, err, tname, "j 2")

		pv, err = test.fn(2400000.5, -320000, 3)

		vvd(t, pv[0][0], 0.9308038666832975759, 1e-11,
			tname, "x 3")
		vvd(t, pv[0][1], 0.3258319040261346000, 1e-11,
			tname, "y 3")
		vvd(t, pv[0][2], 0.1422794544481140560, 1e-11,
			tname, "z 3")

		vvd(t, pv[1][0], -0.6429458958255170006e-2, 1e-11,
			tname, "xd 3")
		vvd(t, pv[1][1], 0.1468570657704237764e-1, 1e-11,
			tname, "yd 3")
		vvd(t, pv[1][2], 0.6406996426270981189e-2, 1e-11,
			tname, "zd 3")

		errEN(t, 1, err, tname, "j 3")

		pv, err = test.fn(2400000.5, 43999.9, 1)

		vvd(t, pv[0][0], 0.2945293959257430832, 1e-11,
			tname, "x 4")
		vvd(t, pv[0][1], -0.2452204176601049596, 1e-11,
			tname, "y 4")
		vvd(t, pv[0][2], -0.1615427700571978153, 1e-11,
			tname, "z 4")

		vvd(t, pv[1][0], 0.1413867871404614441e-1, 1e-11,
			tname, "xd 4")
		vvd(t, pv[1][1], 0.1946548301104706582e-1, 1e-11,
			tname, "yd 4")
		vvd(t, pv[1][2], 0.8929809783898904786e-2, 1e-11,
			tname, "zd 4")

		errT(t, nil, err, tname, "j 4")
	}
}

func BenchmarkPlan94(b *testing.B) {
	tests := []struct {
		ref string
		fn  func(a1, a2 float64, a3 int) (
			[2][3]float64, en.ErrNum)
	}{
		{"cgo", CgoPlan94},
		{"go", GoPlan94},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn(2400000.5, 1e6, 0)
				test.fn(2400000.5, 1e6, 10)
				test.fn(2400000.5, -320000, 3)
				test.fn(2400000.5, 43999.9, 1)
			}
		})
	}
}
