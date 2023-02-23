#include "debug.hpp"
using namespace std;

class Solution {
public:
  vector<int> circularPermutation(int n, int start) {
    if (n == 1) {
      return {start, start ^ 0x1};
    }
    vector<int> &&child = circularPermutation(n - 1, start >> 1);
    vector<int> res;
    for (int i = 0; i < child.size(); i++) {
      res.emplace_back(child[i] * 2 + (start & 0x1));
    }
    for (int i = child.size(); i > 0; i--) {
      res.emplace_back(child[i - 1] * 2 + ((start & 0x1) ^ 0x1));
    }
    /* printf("(%d, %d)", n, start); */
    /* print(res); */
    return res;
  }
};

int main() {
  Solution sl;
  /* print(sl.circularPermutation(1, 1)); */
  print(sl.circularPermutation(3, 1));
}
