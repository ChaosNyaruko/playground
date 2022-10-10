import re
def clean(s):
    n = 1
    r = 1
    while n:
        print("round", r, s)
        s, n = re.subn(r'(.)\1{3,}', '', s)
        print(s, n)
        r += 1
    return s

if __name__ == '__main__':
    print(clean("WRRBBBRAAAAAAAB"))
