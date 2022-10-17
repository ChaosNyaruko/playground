#include "debug.hpp"
using namespace std;
class Solution {
public:
  int totalFruit(vector<int> &fruits) {
    int n = fruits.size();
    vector<int> cnt(n, 0);
    int diff = 0;
    int l = 0, r = 0;
    int res = 1;
    while (r < n) {
      if (cnt[fruits[r]]++ == 0) {
        diff++;
      }
      while (l < r and diff > 2) {
        if (--cnt[fruits[l++]] == 0) {
          diff--;
        }
      }
      /* cout << "header" << l << " " << r << " " << diff << endl; */
      /* print(cnt); */
      res = max(res, r - l + 1);
      r++;
    }
    return res;
  }
};

int main() {
  vector<vector<int>> inputs = {{1, 2, 1}, {0, 1, 2, 2}, {1, 2, 3, 2, 2}};

  Solution sl;
  for (auto &input : inputs) {
    cout << sl.totalFruit(input) << endl;
  }
}
