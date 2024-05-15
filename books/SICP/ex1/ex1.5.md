# Problem

```bash
(define (p) (p)) 
(define (test x y)
	(if (= x 0) 0 y))
```

# Evaluate expression
```bash
(test 0 (p))
```

# Answer
1. Applicative-order evaluation:
	- The interpreter first evaluate the operator and the operands then apply function to evaluate result of args.
	- In this case: when evaluate (p) so we receive results in an infinite recursive loop. Thus the expression never terminate.

2. Normal-order evaluation:
	- The interpreter is not evaluate args until they are actually used within the function.
	- The interpreter subsitute the expresion into the function withou evaluate them
	- In this case: The expression will not evaluate 0 because condition (if (= x 0)) is work.