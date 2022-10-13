#include "debug.hpp"
using namespace std;

class Solution {
public:
  int maxChunksToSorted(vector<int> &arr) {
    int n = arr.size();
    int res = 0;
    int need = 0;
    int curMax = -1;
    int cnt = 0;
    for (int i = 0; i < n; i++) {
      int a = arr[i];
      if (a > curMax) {
        need = (a + 1) - 1 - cnt;
        curMax = max(curMax, a);
      } else {
        need--;
      }
      if (need == 0)
        res++;
      printf("a = %d, curMax = %d, need = %d\n", a, curMax, need);
      /* assert(need >= 0); */
      cnt++;
    }
    return res;
  }
};

int main() {
  Solution sl;
  vector<vector<int>> inputs = {{4, 3, 2, 1, 0}, {1, 0, 2, 3, 4}, {0}};
  for (auto &v : inputs) {
    cout << sl.maxChunksToSorted(v) << endl;
  }
}
