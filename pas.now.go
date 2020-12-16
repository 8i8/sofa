package sofa

// double iauPas(double al, double ap, double bl, double bp)
/*
**  - - - - - - -
**   i a u P a s
**  - - - - - - -
**
**  Position-angle from spherical coordinates.
**
**  This function is part of the International Astronomical Union's
**  SOFA (Standards Of Fundamental Astronomy) software collection.
**
**  Status:  vector/matrix support function.
**
**  Given:
**     al     double     longitude of point A (e.g. RA) in radians
**     ap     double     latitude of point A (e.g. Dec) in radians
**     bl     double     longitude of point B
**     bp     double     latitude of point B
**
**  Returned (function value):
**            double     position angle of B with respect to A
**
**  Notes:
**
**  1) The result is the bearing (position angle), in radians, of point
**     B with respect to point A.  It is in the range -pi to +pi.  The
**     sense is such that if B is a small distance "east" of point A,
**     the bearing is approximately +pi/2.
**
**  2) Zero is returned if the two points are coincident.
**
**  This revision:  2013 June 18
**
**  SOFA release 2020-07-21
**
**  Copyright (C) 2020 IAU SOFA Board.  See notes at end.
*/
{
   double dl, x, y, pa;


   dl = bl - al;
   y = sin(dl) * cos(bp);
   x = sin(bp) * cos(ap) - cos(bp) * sin(ap) * cos(dl);
   pa = ((x != 0.0) || (y != 0.0)) ? atan2(y, x) : 0.0;

   return pa;

