#include "debug.hpp"
using namespace std;
class Solution {
public:
  int partitionDisjoint(vector<int> &nums) {
    int res = 0;
    int allmax = INT_MIN;
    int leftmax = INT_MAX;
    for (int i = 0; i < nums.size(); i++) {
      allmax = max(allmax, nums[i]);
      if (nums[i] < leftmax) {
        res = i + 1;
        leftmax = allmax;
      }
      /* printf("leftmax:%d, allmax:%d, res:%d\n", leftmax, allmax, res); */
    }
    return res;
  }
};

int main() {
  Solution sl;
  vector<vector<int>> inputs = {{5, 0, 3, 8, 6}, {1, 1, 1, 0, 6, 12}};
  for (auto &v : inputs) {
    cout << sl.partitionDisjoint(v) << endl;
  }
  return 0;
}
