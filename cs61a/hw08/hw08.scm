(define (cddr s)
  (cdr (cdr s)))

(define (cadr s)
  (car (cdr s))
)

(define (caddr s)
  (car (cddr s))
)

(define (sign x)
  (cond ((> x 0) 1)
        ((< x 0) -1)
        (else 0)
  )
)

(define (square x) (* x x))

(define (pow b n)
  (cond
    ((= n 0) 1)
    ((= n 1) b)
    ((even? n) (square (pow b (quotient n 2))))
    ((odd? n) (* b (square (pow b (quotient n 2)))))
    )
)

(define (ordered? s)
  (cond 
    ((null? s) #t)
     ((null? (cdr s)) #t)
     ((> (car s) (cadr s)) #f)
     (else (ordered? (cdr s)))
    )
)

; (1 (2.3))
(define (nodots s)
  (cond
    ((null? s) nil)
    ((not (pair? s)) (cons s nil))
    ; s is a pair
    (else (cons (cond 
                   ((pair? (car s)) (nodots (car s)))
                   (else (car s))
                )
                (nodots (cdr s))
          )
    )
  )
)

; Sets as sorted lists

(define (empty? s) (null? s))

(define (contains? s v)
    (cond ((empty? s) #f)
          ((= (car s) v) #t)
          ((> (car s) v) False)
          (else (contains? (cdr s) v)) ; replace this line
          ))

; Equivalent Python code, for your reference:
;
; def empty(s):
;     return s is Link.empty
;
; def contains(s, v):
;     if empty(s):
;         return False
;     elif s.first > v:
;         return False
;     elif s.first == v:
;         return True
;     else:
;         return contains(s.rest, v)

(define (add s v)
    (cond ((empty? s) (list v))
          ((null? s) (cons v nil))
          ((< v (car s)) (cons v s))
          ((> v (car s)) (cons (car s) (add (cdr s) v)))
          (else s) ; replace this line
          ))

(define (intersect s t)
    (cond ((or (empty? s) (empty? t)) nil)
          ((= (car s) (car t)) (cons (car s) (intersect (cdr s) (cdr t))))
          ((< (car s) (car t)) (intersect (cdr s) t))
          (else  (intersect s (cdr t))) ; replace this line
          ))

; Equivalent Python code, for your reference:
;
; def intersect(set1, set2):
;     if empty(set1) or empty(set2):
;         return Link.empty
;     else:
;         e1, e2 = set1.first, set2.first
;         if e1 == e2:
;             return Link(e1, intersect(set1.rest, set2.rest))
;         elif e1 < e2:
;             return intersect(set1.rest, set2)
;         elif e2 < e1:
;             return intersect(set1, set2.rest)

(define (union s t)
    (cond ((empty? s) t)
          ((empty? t) s)
          ; 'YOUR-CODE-HERE
          ((> (car s) (car t)) (cons (car t) (union s (cdr t))))
          ((< (car s) (car t)) (cons (car s) (union (cdr s) t)))
          (else (cons (car s) (union (cdr s) (cdr t)))) ; replace this line
          ))
