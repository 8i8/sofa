package sofa

// #include "sofa.h"
import "C"
import "math"

type ASTROM struct {
	pmt    float64       // PM time interval (SSB, Julian years) 
	eb     [3]float64    // SSB to observer (vector, au) 
	eh     [3]float64    // Sun to observer (unit vector) 
	em     float64       // distance from Sun to observer (au) 
	v      [3]float64    // barycentric observer velocity (vector, c) 
	bm1    float64       // sqrt(1-|v|^2): reciprocal of Lorenz factor 
	bpn    [3][3]float64 // bias-precession-nutation matrix 
	along  float64       // longitude + s' + dERA(DUT) (radians) 
	phi    float64       // geodetic latitude (radians) 
	xpl    float64       // polar motion xp wrt local meridian (radians) 
	ypl    float64       // polar motion yp wrt local meridian (radians) 
	sphi   float64       // sine of geodetic latitude 
	cphi   float64       // cosine of geodetic latitude 
	diurab float64       // magnitude of diurnal aberration vector 
	eral   float64       // "local" Earth rotation angle (radians) 
	refa   float64       // refraction constant A (radians) 
	refb   float64       // refraction constant B (radians) 
}

// Body parameters for light deflection
type LDBODY struct {
	bm float64       // mass of the body (solar masses)
	dl float64       // deflection limiter (radians^2/2)
	pv [2][3]float64 // barycentric PV of the body (au, au/day)
}

// v2sC2Go translates a vector from cgo into go.
func v2sC2Go(in [2]C.double) (out [2]float64) {
	out[0] = float64(in[0])
	out[1] = float64(in[1])
	return
}


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

// v4sIntC2Go translates a 1d 3d vector from cgo into go.
func v4sIntC2Go(in [4]C.int) (out [4]int) {
	out[0] = int(in[0])
	out[1] = int(in[1])
	out[2] = int(in[2])
	out[3] = int(in[3])
	return
}

// v4sIntGo2C translates a 1d 3d vector from go into cgo.
func v4sIntGo2C(in [4]int) (out [4]C.int) {
	out[0] = C.int(in[0])
	out[0] = C.int(in[0])
	out[0] = C.int(in[0])
	out[0] = C.int(in[0])
	return
}

// astrGo2C translates an ASTROM from go into cgo.
func astrGo2C(in ASTROM) (out C.iauASTROM) {

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

// astrC2Go translates an iauASTROM from cgo into go.
func astrC2Go(in C.iauASTROM) (out ASTROM) {

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

// ldbodyGo2C translates an LDBODY from go into cgo.
func ldbodyGo2C(in LDBODY) (out C.iauLDBODY) {
	out.bm = C.double(in.bm)
	out.dl = C.double(in.dl)
	out.pv = v3dGo2C(in.pv)
	return
}

// ldbodyC2Go( translates an iauLDBODY from cgo into go.
func ldbodyC2Go(in C.iauLDBODY) (out LDBODY) {
	out.bm = float64(in.bm)
	out.dl = float64(in.dl)
	out.pv = v3dC2Go(in.pv)
	return
}

// dsign gives the magnitude of 'a' with sign of 'b' (double).
func dsign(a, b float64) float64 {
	if b < 0.0 {
		return -math.Abs(a)
	}
	return math.Abs(a)
}

// dnint round to nearest whole number (double).
func dnint(a float64) (res float64) {
	if math.Abs(a) < 0.5 {
		res = 0.0
	} else if a < 0.0 {
		res = math.Ceil(a - 0.5)
	} else {
		res = math.Floor(a + 0.5)
	}
	return
}

/* dint(A) - truncate to nearest whole number towards zero (double) */
func dint(a float64) (res float64) {
	if a < 0.0 {
		res = math.Ceil(a)
	} else {
		res = math.Floor(a)
	}
	return
}

// fmax returns the greatest of two floats.
func fmax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// fmin returns the lesast of two floats.
func fmin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// pow is an integer power calculation.
func pow(base, exp int) int {
	var res = 1
	for exp > 0 {
		if exp&1 > 0 {
			res *= base
		}
		exp >>= 1
		base *= base
	}
	return res
}
