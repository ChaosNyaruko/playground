#include "debug.hpp"
using namespace std;
class Solution {
public:
  int numComponents(ListNode *head, vector<int> &nums) {
    ListNode *cur = head;
    unordered_set<int> s(nums.begin(), nums.end());
    int res = 0;
    int p = 0;
    bool started = false;
    while (cur) {
      if (s.count(cur->val)) {
        if (not started) {
          started = true;
          p++;
        }
      } else {
        started = false;
      }
      cur = cur->next;
    }
    return p;
  }
};

int main() {
  Solution sl;
  // head [0, 1, 2, 3] nums [0, 1, 3] res 2
  // [0,1] [3] max 2
  // head [0, 1, 2, 3, 4] nums [0, 3, 1, 4] res 2
  // [0,1] [3,4] max 2
  vector<int> head = {0, 1, 2, 3}, nums = {0, 1, 3};
  cout << sl.numComponents(build_list(head), nums) << endl;
  head = {0, 1, 2, 3, 4}, nums = {0, 3, 1, 4};
  cout << sl.numComponents(build_list(head), nums) << endl;
  head = {0, 1, 2}, nums = {0, 2};
  cout << sl.numComponents(build_list(head), nums) << endl;
  head = {0, 1, 2}, nums = {1, 0};
  cout << sl.numComponents(build_list(head), nums) << endl;
}
