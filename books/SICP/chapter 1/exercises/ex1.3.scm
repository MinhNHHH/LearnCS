(define a 1)
(define b 2)
(define c 3)
(define (square a) (* a a))

(define (isSmallest x y z) 
  (and (< x y) (< x z)))

(define (sum-square x y) (+ (square x) (square y)))

(define (sum-square-of-two-larger-nums x y z) 
  (cond ((isSmallest x y z) (sum-square y z))
      ((isSmallest y x z) (sum-square x z))
  (else (sum-square x y)))
)

(sum-square-of-two-larger-nums a b c)
(sum-square-of-two-larger-nums 2 4 1)