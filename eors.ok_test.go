package sofa

import "testing"

//
//  - - - - - - - - -
//   T e s t E o r s
//  - - - - - - - - -
//
//  Test Eors function.
//
//  Called:  Eors, vvd
//
//  This revision:  2013 August 7
//
func TestEors(t *testing.T) {
	const fname = "Eors"
	var rnpb [3][3]float64
	var s float64

	rnpb[0][0] = 0.9999989440476103608
	rnpb[0][1] = -0.1332881761240011518e-2
	rnpb[0][2] = -0.5790767434730085097e-3

	rnpb[1][0] = 0.1332858254308954453e-2
	rnpb[1][1] = 0.9999991109044505944
	rnpb[1][2] = -0.4097782710401555759e-4

	rnpb[2][0] = 0.5791308472168153320e-3
	rnpb[2][1] = 0.4020595661593994396e-4
	rnpb[2][2] = 0.9999998314954572365

	s = -0.1220040848472271978e-7

	tests := []struct {
		ref string
		fn  func([3][3]float64, float64) float64
	}{
		{"cgo", CgoEors},
		{"go", GoEors},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref
		eo := test.fn(rnpb, s)
		vvd(t, eo, -0.1332882715130744606e-2, 1e-14, tname, "")
	}
}

func BenchmarkEors(b *testing.B) {
	var rnpb [3][3]float64
	var s float64

	rnpb[0][0] = 0.9999989440476103608
	rnpb[0][1] = -0.1332881761240011518e-2
	rnpb[0][2] = -0.5790767434730085097e-3

	rnpb[1][0] = 0.1332858254308954453e-2
	rnpb[1][1] = 0.9999991109044505944
	rnpb[1][2] = -0.4097782710401555759e-4

	rnpb[2][0] = 0.5791308472168153320e-3
	rnpb[2][1] = 0.4020595661593994396e-4
	rnpb[2][2] = 0.9999998314954572365

	s = -0.1220040848472271978e-7

	tests := []struct {
		ref string
		fn  func([3][3]float64, float64) float64
	}{
		{"cgo", CgoEors},
		{"go", GoEors},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(rnpb, s)
			}
		})
	}
}
