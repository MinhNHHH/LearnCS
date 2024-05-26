(define (* a b) 
	(if (= b 0)
      0
      (+ a (* a (- b 1)))))

(define (even? n)
	(= (remainder n 2) 0))

(define (fast-multiply a b)
  (cond ((= b 0) 0)
        ((even? b) (fast-multiply (* 2 a) (/ b 2)))
        (else (+ a (fast-multiply a (- b 1))))
))

(fast-multiply 3 4)
