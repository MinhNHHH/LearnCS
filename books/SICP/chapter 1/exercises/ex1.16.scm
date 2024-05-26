(define (even? n)
	(= (remainder n 2) 0))

(define (fast-expt b n)
	(define (fast-expt-iter b n product)
	(cond ((= n 0) product)
		((even? n) (fast-expt-iter (square b) (/ n 2) product))
		(else (fast-expt-iter b (- n 1) (* b product)))))
	(fast-expt-iter b n 1))


(fast-expt 3 2)
(fast-expt 3 3)
(fast-expt 3 4)
