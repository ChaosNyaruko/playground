#include "debug.hpp"
using namespace std;
class Solution {
public:
  int shortestSubarray(vector<int> &nums, int k) {
    /* printf("starting\n"); */
    int res = INT_MAX;
    deque<int> q;
    int n = nums.size();
    vector<int> preSum(n + 1, 0);
    int curSum = 0;
    q.push_back(0);
    for (int j = 1; j < n + 1; j++) {
      curSum += nums[j - 1];
      preSum[j] = curSum;
      // curSum = s[i+1]
      /* printf("j = %d, curSum = %d\n", j, curSum); */
      while (!q.empty() and curSum - preSum[q.front()] >= k) {
        /* printf("pop front -- j=%d, i=%d\n", j, q.front()); */
        res = min(res, j - q.front());
        q.pop_front();
      }
      while (!q.empty() and curSum <= preSum[q.back()]) {
        /* printf("pop back -- j=%d, old=%d\n", j, q.back()); */
        q.pop_back();
      }
      q.push_back(j);
    }
    /* print(preSum); */
    return res == INT_MAX ? -1 : res;
  }
};

int main() {
  Solution sl;
  vector<pair<vector<int>, int>> inputs = {make_pair(vector<int>{1}, 1),
                                           make_pair(vector<int>{1, 2}, 4),
                                           make_pair(vector<int>{2, -1, 2}, 3)};
  for (auto &i : inputs) {
    cout << sl.shortestSubarray(i.first, i.second) << endl;
  }
  return 0;
}
