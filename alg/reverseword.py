def reverseWord(s):
    i, j = 0, 0
    n = len(s)
    while i < n:
        while i < j or i < n and s[i] == ' ': i += 1
        while j < i or j < n and s[j] != ' ': j += 1
        print(s[i:j])

if __name__ == '__main__':
    reverseWord("  hello word")


