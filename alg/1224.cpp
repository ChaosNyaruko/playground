#include "debug.hpp"
using namespace std;
class Solution {
public:
  int maxEqualFreq(vector<int> &nums) {
    print(move(nums));
    int res = 1;
    unordered_map<int, int> c, f;
    int maxF = 1;
    for (int i = 0; i < nums.size(); i++) {
      int num = nums[i];
      cout << "i:" << i << endl;
      bool ok = false;
      if (c[num] > 0) {
        f[c[num]]--;
      }
      c[num]++;
      int newf = c[num];
      if (newf > maxF) {
        maxF = newf;
      }
      f[c[num]]++;
      if (maxF == 1) {
        cout << "\t case1" << endl;
        ok = true;
      } else if (f[maxF] == 1 and
                 maxF * f[maxF] + (maxF - 1) * f[maxF - 1] == i + 1) {
        cout << "\t case2" << endl;
        ok = true;
      } else if (f[maxF] * maxF == i and f[1] == 1) {
        cout << "\t case3" << endl;
        ok = true;
      }
      if (ok) {
        res = max(res, i + 1);
      }
      for (auto &[k, v] : f) {
        cout << "\t" << k << " " << v << endl;
      }
      printf("\tnums[i]:%d, count:%d, freq:%d, maxF:%d, ok:%d\n", nums[i],
             c[nums[i]], f[c[nums[i]]], maxF, ok);
    }
    return res;
  }
};

int main() {
  Solution sl;
  /* vector<int> input = {2, 2, 1, 1, 5, 3, 3, 5}; */
  /* vector<int> input = {1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}; */
  /* vector<int> input = {1, 2}; */
  /* vector<int> input = {1, 1}; */
  vector<int> input = {
      1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
      2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2};
  cout << sl.maxEqualFreq(input) << endl;
  return 0;
}

/*
输入：nums = [2,2,1,1,5,3,3,5]
输出：7
解释：对于长度为 7 的子数组 [2,2,1,1,5,3,3]，如果我们从中删去 nums[4] =
5，就可以得到 [2,2,1,1,3,3]，里面每个数字都出现了两次。 示例 2：

输入：nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
输出：13
*/
