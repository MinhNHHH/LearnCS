(define (even? n)
	(= (remainder n 2) 0))

(define (fast-multiply a b)
	(define (iter a b product)
		(cond ((= 0 b) product)
			((even? b) (iter (* 2 a) (/ b 2) product))
			(else  (iter a (- b 1) (+ a product)))))
	(iter a b 0))
		
(fast-multiply 4 2)
(fast-multiply 3 2)