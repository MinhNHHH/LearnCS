(define (abs x) (if (< x 0)
(- x) x))

(define (improve guess x)
  (/ (+ (/ x (square guess)) (* 2 guess)) 3))

(define (good-enough? guess x)
	(let ((new-guess (improve guess x)))
    (< (abs (- new-guess guess)) (* 0.00001 guess))))

(define (cube-root guess x)
	(if (good-enough? guess x) 
		guess 
		(cube-root (improve guess x) x
)))


(define (cube x) (cube-root 1.0 x))

(cube 27)
(cube 64)
(cube 0.00064)
