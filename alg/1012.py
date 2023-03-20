import functools
class Solution:
  def numDupDigitsAtMostN(self, n: int) -> int:
    s = str(n)
    @functools.cache
    def f(mask, i, same):
      if i == len(s):
        return 1;
      t = 9 if not same else ord(s[i]) - ord('0')
      res = 0
      for k in range(0, t+1):
        if (mask & (1 << k)):
          continue
        if mask == 0 and k == 0:
          newMask = mask
        else:
          newMask = mask | (1 << k)
        res += f(newMask, i + 1, same and k == t)
      return res
    return n + 1 - f(0, 0, True)

if __name__ == '__main__':
  for n in range(0, 100):
    print(f"n={n} answer={Solution().numDupDigitsAtMostN(n)}")

