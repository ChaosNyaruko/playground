#include "debug.hpp"
using namespace std;
class Solution {
public:
  int movesToMakeZigzag(vector<int> &nums) {
    auto helper = [&](int pos) -> int {
      int res = 0;
      for (int i = pos; i < nums.size(); i += 2) {
        int s = 0;
        if (i - 1 >= 0 && nums[i] >= nums[i - 1]) {
          s = nums[i] - nums[i - 1] + 1;
        }
        if (i + 1 < nums.size() && nums[i] >= nums[i + 1]) {
          s = max(nums[i] - nums[i + 1] + 1, s);
        }
        printf("nums[%d] = %d, s= %d\n", i, nums[i], s);
        res += s;
      }
      return res;
    };
    return min(helper(0), helper(1));
  }
};

int main() {
  Solution sl;
  vector<vector<int>> inputs = {{2, 7, 10, 9, 8, 9}};
  for (auto &i : inputs) {
    cout << sl.movesToMakeZigzag(i) << endl;
  }
}
