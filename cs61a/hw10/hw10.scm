(define (accumulate combiner start n term)
  (if (= n 0)
      start
      (combiner (term n) (accumulate combiner start (- n 1) term))
  )
)

(define (accumulate-tail combiner start n term)
  (define (helper combiner start n term)
    (if (= n 0)
      start
      (helper combiner (combiner start (term n)) (- n 1) term)
      )
    )
  (helper combiner start n term)
)

; (define-macro (list-of expr for var in seq if filter-fn)
;               (list 'map (list 'lambda (list var) expr) (list 'filter (list 'lambda (list var) filter-fn) seq))
; )

; using quasi-quote
(define-macro (list-of expr for var in seq if filter-fn)
              `(map (lambda (,var) ,expr) (filter (lambda (,var) ,filter-fn) ,seq))
)
