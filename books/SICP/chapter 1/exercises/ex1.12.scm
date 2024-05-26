(define (pascals r c)
	(cond ((or (= c 0) (= c r)) 1)
			(else (+ (pascals (- r 1) (- c 1))
								(pascals (- r 1) c)))))

(pascals 3 1)