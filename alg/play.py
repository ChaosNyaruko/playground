a = [2,3,1,5]
first = a[0]
second = a[1]
for i in range(2, len(a)):
    n = a[i]
    if n > second:
        break
    elif n > first:
        second = n
    else:
        first = n
    print(first, second)
print("out", first, second)

