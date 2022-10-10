; Q1
(define (compose-all funcs)
  (define (dfs funcs x) 
      (cond ((null? funcs) x)
            (else (dfs (cdr funcs) ((car funcs) x)))
        )
    )
  ; (define (inner x) (dfs funcs x))
  (lambda (x) (dfs funcs x))
)

; Q2
(define (tail-replicate x n)
  (define (dfs x n ret) 
      (if (= n 0) ret
            (dfs x (- n 1) (cons x ret))
            )
      )
  (dfs x n nil)
)
