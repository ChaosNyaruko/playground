#include "debug.hpp"
using namespace std;

//  f(mask,  i, same)
//  mask: which digit has been used for ith
//  i: working on ith digit(left 0 -> right n - 1)
//  same: i's range, i.e. 0..i-1 is the same as N,
//        if same if true, i can only have a range of 0~n[i],
//        otherwise 0~9 are all allowed.
class Solution {
public:
  int A(int x, int y) {
    int res = 1;
    for (int i = 0; i < x; i++) {
      res *= y--;
    }
    return res;
  }

  int f(int mask, const string &sn, int i, bool same) {
    if (i == sn.size()) {
      return 1;
    }
    int t = same ? sn[i] - '0' : 9, res = 0, c = __builtin_popcount(mask) + 1;
    for (int k = 0; k <= t; k++) {
      if (mask & (1 << k)) {
        continue;
      }
      if (same && k == t) {
        res += f(mask | (1 << t), sn, i + 1, true);
      } else if (mask == 0 && k == 0) {
        res += f(0, sn, i + 1, false);
      } else {
        res += A(sn.size() - 1 - i, 10 - c);
      }
    }
    return res;
  }

  int numDupDigitsAtMostN(int n) {
    string sn = to_string(n);
    return n + 1 - f(0, sn, 0, true);
  }
};

class Solution1 {
public:
  vector<vector<int>> dp;

  int f(int mask, const string &sn, int i, bool same) {
    if (i == sn.size()) {
      return 1;
    }
    if (!same && dp[i][mask] >= 0) {
      return dp[i][mask];
    }
    int res = 0, t = same ? (sn[i] - '0') : 9;
    for (int k = 0; k <= t; k++) {
      if (mask & (1 << k)) {
        continue;
      }
      res += f(mask == 0 && k == 0 ? mask : mask | (1 << k), sn, i + 1,
               same && k == t);
    }
    if (!same) {
      dp[i][mask] = res;
    }
    return res;
  }

  int numDupDigitsAtMostN(int n) {
    string sn = to_string(n);
    dp.resize(sn.size(), vector<int>(1 << 10, -1));
    return n + 1 - f(0, sn, 0, true);
  }
};
int main() {
  Solution1 sl;
  for (int i = 0; i < 100; i++) {
    cout << "n=" << i << " answer=" << sl.numDupDigitsAtMostN(i) << endl;
  }
}
