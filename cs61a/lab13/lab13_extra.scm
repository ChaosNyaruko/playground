; Q4
(define (rle s)
  (define (helper prev cnt s) 
    (cond 
           ((null? s) (cons-stream (list prev cnt) nil))
           ((eq? prev (car s)) (helper prev (+ cnt 1) (cdr-stream s)))
           (else (cons-stream (list prev cnt) (helper (car s) 1 (cdr-stream s))))
           )
    )
  (if (null? s) nil (helper (car s) 1 (cdr-stream s)))
)

; Q4 testing functions
(define (list-to-stream lst)
    (if (null? lst) nil
                    (cons-stream (car lst) (list-to-stream (cdr lst))))
)

(define (stream-to-list s)
    (if (null? s) nil
                 (cons (car s) (stream-to-list (cdr-stream s))))
)

; Q5
; (define (insert n s)
;   (define (helper n prev after)
;     (cond ((null? after) (append prev (list n)))
;         ((> n (car after)) (helper n (append prev (list (car after))) (cdr after)))
;         ; insert before (after)
;         (else (append prev (list n) after))
;         )
;    )
;   (helper n nil s)
; )

; ref, still not constant space
(define (insert n s)
    (define (helper n prev suf)
      (if (null? suf) (append prev (list n))
        (if (< n (car suf))
            (append (if (null? prev) (list n) (append prev (list n))) suf)
            (helper n (append prev (list (car suf))) (cdr suf))
        )
      )
    )
    (helper n nil s)
)


; Q6
(define (deep-map fn s)
  (cond ((null? s) nil)
        ((list? (car s)) (cons (deep-map fn (car s)) (deep-map fn (cdr s))))
        ( else (cons (fn (car s)) (deep-map fn (cdr s))))
  )
)

; Q7
; Feel free to use these helper procedures in your solution
(define (map fn s)
  (if (null? s) nil
      (cons (fn (car s))
            (map fn (cdr s)))))

(define (filter fn s)
  (cond ((null? s) nil)
        ((fn (car s)) (cons (car s)
                            (filter fn (cdr s))))
        (else (filter fn (cdr s)))))

; Implementing and using these helper procedures is optional. You are allowed
; to delete them.
(define (unique s)
  (cond ((null? s) nil)
        (else (cons (car s) (filter (lambda (x) (not (eq? (car s) x))) (unique (cdr s)))))
        )
)

(define (count name s)
  (cond ((null? s) 0)
        ((eq? (car s) name) (+ 1 (count name (cdr s))))
        (else (count name (cdr s)))
        )
)

(define (tally names)
  (map (lambda (x) (cons x (count x names)))  (unique names))
  ; expected ((james . 1)) not ((james 1))
  ; (map (lambda (x) (list x (count x names)))  (unique names))
)
