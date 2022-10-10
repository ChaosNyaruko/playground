#include "stdc++.h"
using namespace std;
class Solution {
public:
  int minSwap(vector<int> &nums1, vector<int> &nums2) {
    int n = nums1.size();
    vector<int> use(n, n);
    vector<int> not_use(n, n);
    use[0] = 1;
    not_use[0] = 0;
    for (int i = 1; i < n; i++) {
      if (nums2[i] > nums2[i - 1] and nums1[i] > nums1[i - 1]) {
        not_use[i] = min(not_use[i], not_use[i - 1]);
        use[i] = min(use[i], use[i - 1] + 1);
      }
      if (nums2[i] > nums1[i - 1] and nums1[i] > nums2[i - 1]) {
        use[i] = min(use[i], not_use[i - 1] + 1);
        not_use[i] = min(not_use[i], use[i - 1]);
      }
    }
    return min(use[n - 1], not_use[n - 1]);
  }
};

int main() {
  vector<vector<int>> i1 = {{1, 3, 5, 4}, {1, 2, 3, 7}};
  vector<vector<int>> i2 = {{0, 3, 5, 8, 9}, {2, 1, 4, 6, 9}};
  Solution sl;
  cout << sl.minSwap(i1[0], i1[1]) << endl;
  cout << sl.minSwap(i2[0], i2[1]) << endl;
  return 0;
}
