"""
  Alternative Strategy: Relative Change
  + We check if the absolute difference between the new guess and the current guess is less than a small fraction (eg: 0.00001) 
  + This method adapts to the scale of the numbers being worked with:
    + For small numbers, the guess quickly converges as the relative change becomes significant even with small differences.
    + For large numbers, the method avoids the pitfalls of limited precision by focusing on the relative change rather than absolute differences.
"""

(define (sqrt-iter guess x)
	(if (good-enough? guess x) 
		guess 
		(sqrt-iter (improve guess x) x
)))

(define (improve guess x)
	(average guess (/ x guess)))

(define (average x y)
	(/ (+ x y) 2))

(define (good-enough? guess x)
	(let ((new-guess (improve guess x)))
    (< (abs (- new-guess guess)) (* 0.00001 guess))))

(define (abs x) (if (< x 0)
(- x) x))

(define (sqrt x) (sqrt-iter 1.0 x))

(sqrt 0.0004)
(sqrt 0.0001)
(sqrt (+ 100 37)) 
