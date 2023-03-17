#include "debug.hpp"
using namespace std;
class Solution {
public:
  vector<int> answerQueries(vector<int> &nums, vector<int> &queries) {
    sort(nums.begin(), nums.end());
    vector<int> res;
    vector<int> sum(nums.size() + 1, 0);
    for (int i = 0; i < nums.size(); i++) {
      sum[i + 1] = sum[i] + nums[i];
    }
    res.reserve(queries.size());
    for (auto q : queries) {
      // find the biggest i that makes sum[0..i] <= q
      int l = 0, r = nums.size();
      while (l < r) {
        int m = l + (r - l + 1) / 2;
        if (sum[m] > q) {
          r = m - 1;
        } else {
          l = m;
        }
      }
      res.push_back(l);
    }
    return std::move(res);
  }
};

/* [4,5,2,1] */
/* [3,10,21] */
/* [2,3,4,5] */
/* [1] */
int main() {
  Solution sl;

  vector<tuple<vector<int>, vector<int>>> inputs = {
      make_tuple(vector<int>{4, 5, 2, 1}, vector<int>{3, 10, 21}),
      make_tuple(vector<int>{2, 3, 4, 5}, vector<int>{1})};
  for (auto &&input : inputs) {
    print(sl.answerQueries(get<0>(input), get<1>(input)));
  }
}
