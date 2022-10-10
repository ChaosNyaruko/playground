#include "debug.hpp"
using namespace std;
class Solution {
public:
  int busyStudent(vector<int> &startTime, vector<int> &endTime, int queryTime) {
    int res = 0;
    for (int i = 0; i < startTime.size(); i++) {
      int s = startTime[i], e = endTime[i];
      if (s <= queryTime and e >= queryTime) {
        res++;
      }
    }
    return res;
  }
};

int main() {
  vector<int> start = {4};
  vector<int> end = {4};
  int queryTime = 4;
  /* vector<int> start = {9, 8, 7, 6, 5, 4, 3, 2, 1}; */
  /* vector<int> end = {10, 10, 10, 10, 10, 10, 10, 10, 10}; */
  /* int queryTime = 5; */
  Solution sl;
  cout << sl.busyStudent(start, end, queryTime);
}
