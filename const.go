package sofa

type ASTROM struct {
	pmt    float64       /* PM time interval (SSB, Julian years) */
	eb     [3]float64    /* SSB to observer (vector, au) */
	eh     [3]float64    /* Sun to observer (unit vector) */
	em     float64       /* distance from Sun to observer (au) */
	v      [3]float64    /* barycentric observer velocity (vector, c) */
	bm1    float64       /* sqrt(1-|v|^2): reciprocal of Lorenz factor */
	bpn    [3][3]float64 /* bias-precession-nutation matrix */
	along  float64       /* longitude + s' + dERA(DUT) (radians) */
	phi    float64       /* geodetic latitude (radians) */
	xpl    float64       /* polar motion xp wrt local meridian (radians) */
	ypl    float64       /* polar motion yp wrt local meridian (radians) */
	sphi   float64       /* sine of geodetic latitude */
	cphi   float64       /* cosine of geodetic latitude */
	diurab float64       /* magnitude of diurnal aberration vector */
	eral   float64       /* "local" Earth rotation angle (radians) */
	refa   float64       /* refraction constant A (radians) */
	refb   float64       /* refraction constant B (radians) */
}

const (
	// D2PI 2Pi
	D2PI = 6.283185307179586476925287

	// Schwarzschild radius of the Sun (au)
	// = 2 * 1.32712440041e20 / (2.99792458e8)^2 / 1.49597870700e11
	SRS = 1.97412574336e-8

	/* Arcseconds to radians */
	DAS2R = 4.848136811095359935899141e-6
)
