// s string even(n) a-z
// a ""
// op1:   a= s[0] + a s = s[1:]
// op2    a= s[-1] + a s.pop()
// b ""
// cbzzzzzz
//
// P(i, j) -> 0/1/2 [i, j] // draw 0
// s[i:j+1] 
// s[i] -> max_string = TODO(P(i+1, j)) + s[i]
// TODO=
// P(i, j) = P(i+1, j) == 2 
// j - i >= 1 j - i must be odd
// s = abcd
// P(1, 2) = win
// P(0, 3) = 
//    1. try s[0] = a ->  opponent b: P(i+2, j) d:P(i+1, j-1)
//    2. try s[3] = d -> P(1,2) win but d > a -> p(0, 3) == win(2)   
//    P(0, 3) = max(p1, p2) = win
// P(i, j) = s[
