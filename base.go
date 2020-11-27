package sofa

// #include "sofa.h"
import "C"

// v3sC2GO translates a 3d vector from cgo into go.
func v3sC2Go(in [3]C.double) (out [3]float64) {
	out[0] = float64(in[0])
	out[1] = float64(in[1])
	out[2] = float64(in[2])
	return
}

// v3sGo2C translates a 3d vector from go into cgo.
func v3sGo2C(in [3]float64) (out [3]C.double) {
	out[0] = C.double(in[0])
	out[1] = C.double(in[1])
	out[2] = C.double(in[2])
	return
}

// v3dC2Go translates a 3d vector pair from cgo into go.
func v3dC2Go(in [2][3]C.double) (out [2][3]float64) {
	out[0][0] = float64(in[0][0])
	out[0][1] = float64(in[0][1])
	out[0][2] = float64(in[0][2])
	out[1][0] = float64(in[1][0])
	out[1][1] = float64(in[1][1])
	out[1][2] = float64(in[1][2])
	return
}

// v3dGo2C translates a 3d vector pair from go into cgo.
func v3dGo2C(in [2][3]float64) (out [2][3]C.double) {
	out[0][0] = C.double(in[0][0])
	out[0][1] = C.double(in[0][1])
	out[0][2] = C.double(in[0][2])
	out[1][0] = C.double(in[1][0])
	out[1][1] = C.double(in[1][1])
	out[1][2] = C.double(in[1][2])
	return
}

// v3tC2Go translates a 3d vector triple from cgo into go.
func v3tC2Go(in [3][3]C.double) (out [3][3]float64) {
	out[0][0] = float64(in[0][0])
	out[0][1] = float64(in[0][1])
	out[0][2] = float64(in[0][2])
	out[1][0] = float64(in[1][0])
	out[1][1] = float64(in[1][1])
	out[1][2] = float64(in[1][2])
	out[2][0] = float64(in[2][0])
	out[2][1] = float64(in[2][1])
	out[2][2] = float64(in[2][2])
	return
}

// v3tGo2C translates a 3d vector triple from go into cgo.
func v3tGo2C(in [3][3]float64) (out [3][3]C.double) {
	out[0][0] = C.double(in[0][0])
	out[0][1] = C.double(in[0][1])
	out[0][2] = C.double(in[0][2])
	out[1][0] = C.double(in[1][0])
	out[1][1] = C.double(in[1][1])
	out[1][2] = C.double(in[1][2])
	out[2][0] = C.double(in[2][0])
	out[2][1] = C.double(in[2][1])
	out[2][2] = C.double(in[2][2])
	return
}

// aGo2C translates an ASTROM from go into cgo.
func aGo2C(in ASTROM) (out C.iauASTROM) {

	out.pmt = C.double(in.pmt)
	out.eb = v3sGo2C(in.eb)
	out.eh = v3sGo2C(in.eh)
	out.em = C.double(in.em)
	out.v = v3sGo2C(in.v)
	out.bm1 = C.double(in.bm1)
	out.bpn = v3tGo2C(in.bpn)
	out.along = C.double(in.along)
	out.xpl = C.double(in.xpl)
	out.ypl = C.double(in.ypl)
	out.sphi = C.double(in.sphi)
	out.cphi = C.double(in.cphi)
	out.diurab = C.double(in.diurab)
	out.eral = C.double(in.eral)
	out.refa = C.double(in.refa)
	out.refb = C.double(in.refb)

	return
}

// aC2Go translates an iauASTROM from cgo into go.
func aC2Go(in C.iauASTROM) (out ASTROM) {

	out.pmt = float64(in.pmt)
	out.eb = v3sC2Go(in.eb)
	out.eh = v3sC2Go(in.eh)
	out.em = float64(in.em)
	out.v = v3sC2Go(in.v)
	out.bm1 = float64(in.bm1)
	out.bpn = v3tC2Go(in.bpn)
	out.along = float64(in.along)
	out.xpl = float64(in.xpl)
	out.ypl = float64(in.ypl)
	out.sphi = float64(in.sphi)
	out.cphi = float64(in.cphi)
	out.diurab = float64(in.diurab)
	out.eral = float64(in.eral)
	out.refa = float64(in.refa)
	out.refb = float64(in.refb)

	return
}
