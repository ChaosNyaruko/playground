#include "debug.hpp"
using namespace std;
class Solution {
public:
  int sumSubarrayMins(vector<int> &arr) {
    int n = arr.size();
    int mod = 1e9 + 7;
    stack<int> s;
    vector<int> left(n, 0);
    vector<int> right(n, 0);
    for (int i = 0; i < n; i++) {
      while (!s.empty() and arr[i] < arr[s.top()]) {
        s.pop();
      }
      left[i] = s.empty() ? i + 1 : i - s.top();
      s.push(i);
    }
    s = stack<int>{};
    for (int i = n - 1; i >= 0; i--) {
      while (!s.empty() and arr[i] <= arr[s.top()]) {
        s.pop();
      }
      right[i] = s.empty() ? n - i : s.top() - i;
      s.push(i);
    }
    /* cout << ">>> " << endl; */
    /* print(left); */
    /* cout << "--- " << endl; */
    /* print(right); */
    /* cout << "<<< " << endl; */
    long res = 0;
    for (int i = 0; i < n; i++) {
      res = (res + arr[i] * (long)left[i] * right[i]) % mod;
    }
    return res;
  }
};

// [3,1,2,4] 17
// [11,81,94,43,3] 444
// [117,1315,1336,4213,5634,6288,7640,8533,9688,10186,10593,11896,13673,14707,15484,17429,19639,20416,21375,23601,25800,26485,27893,28026,28695,29121,28642,28023,27642,26324,23844,22069,21124,20181,18957,15736,15364,13749,13612,11062,10319,9755,9367,7977,6463,6049,4886,3071,1331,865]
// 12363569
int main() {
  Solution sl;
  vector<vector<int>> inputs = {
      {3, 1, 2, 4},
      {11, 81, 94, 43, 3},
      {117,   1315,  1336,  4213,  5634,  6288,  7640,  8533,  9688,  10186,
       10593, 11896, 13673, 14707, 15484, 17429, 19639, 20416, 21375, 23601,
       25800, 26485, 27893, 28026, 28695, 29121, 28642, 28023, 27642, 26324,
       23844, 22069, 21124, 20181, 18957, 15736, 15364, 13749, 13612, 11062,
       10319, 9755,  9367,  7977,  6463,  6049,  4886,  3071,  1331,  865}};
  for (auto &input : inputs) {
    cout << sl.sumSubarrayMins(input) << endl;
  }
}
