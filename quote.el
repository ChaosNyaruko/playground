(defun print_first_fun (a b)  (c))

(print_first_fun (+ 1 2) (message  "%s" (+ 3 4)))

(defmacro print_first_macro (a b c) (message "%s" a))

(print_first_macro (+ 1 2) a b)
