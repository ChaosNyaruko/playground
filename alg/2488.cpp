#include "debug.hpp"
using namespace std;
/* for (int i = 0; i < n; i++) { */
/*   sum += sign(nums[i] - k); */
/*   if (i < kIndex) { */
/*     counts[sum]++; */
/*   } else { */
/*     int prev0 = counts[sum]; */
/*     int prev1 = counts[sum - 1]; */
/*     ans += prev0 + prev1; */
/*   } */
/* } */

class Solution {
public:
  int countSubarrays(vector<int> &nums, int k) {
    int n = nums.size();
    int kIndex = n + 1;
    unordered_map<int, int> cnt;
    int sum = 0;
    cnt[0] = 1;
    int res = 0;

    for (int i = 0; i < n; i++) {
      sum += (nums[i] > k ? 1 : (nums[i] == k ? 0 : -1));
      if (nums[i] == k) {
        kIndex = i;
      }
      if (i < kIndex) { // i + 1 can be the left bound
        cnt[sum]++;
      } else { // i can be the right bound
        int prev0 = cnt[sum - 0];
        int prev1 = cnt[sum - 1];
        res += prev0 + prev1;
      }
    }
    return res;
  }
};

//[3,2,1,4,5]
// 4
/* [2,3,1] */
/* 3 */
int main() {
  Solution sl;
  vector<tuple<vector<int>, int>> inputs = {
      make_tuple(vector<int>{3, 2, 1, 4, 5}, 4),
      make_tuple(vector<int>{2, 3, 1}, 3)};
  for (auto input : inputs) {
    cout << sl.countSubarrays(get<0>(input), get<1>(input)) << endl;
  }
}
