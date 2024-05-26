; Matrix Representation of Fibonacci Sequence


; Matrix Exponentiation by Squaring

; To compute Tn−1Tn−1 efficiently, we use the method of matrix exponentiation by squaring, similar to the fast exponentiation algorithm.

; Base Case: If n=1n=1, then T1=TT1=T.
; Recursive Case: To compute TnTn:
;  If nn is even, we use the property Tn=(Tn/2)2Tn=(Tn/2)2.
;  If nn is odd, we use Tn=T⋅Tn−1Tn=T⋅Tn−1.

; Transformations TpqTpq​

; Given a transformation TpqTpq​ defined by:

; a′=bq+aq+apa′=bq+aq+ap
; b′=bp+aqb′=bp+aq

; To compute the effect of applying this transformation twice, we need to multiply the transformation matrices.
; When we apply TpqTpq​ twice, we get:
; Tpq2 ​= Tpq​×Tpq
; p′= p2 + q2
; q′= pq + 2q2

(define (sum-two-square a b)
	(+ (square a) (square b)))

(define (fib n) (fib-iter 1 0 0 1 n))
	(define (fib-iter a b p q count) 
		(cond ((= count 0) b)
					((even? count) (fib-iter a b (sum-two-square p q) (+ (* 2 (* p q)) (square q)) (/ count 2)))
					(else (fib-iter (+ (* b q) (* a q) (* a p)) (+ (* b p) (* a q)) p q (- count 1)))))

(fib 5)