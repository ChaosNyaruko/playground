
#include <iostream>
#include <stack>
#include <vector>
using namespace std;

// [0,1,0,2,1,0,1,3,2,1,2,1] --> 6
int Solve(vector<int> &input) {
  input.push_back(0);
  stack<int> stk;
  int res = 0;
  for (int i = 0; i < input.size(); i++) {
    while (!stk.empty() && input[i] >= input[stk.top()]) {
      int top = stk.top();
      stk.pop();
      if (!stk.empty()) {
        int start = stk.top();
        int height = min(input[i], input[start]) - input[top];
        res += height * (i - start - 1);
        printf("i: %d, start: %d, top: %d, height: %d\n", i, start, top, height);
      }
    }

    stk.push(i);
  }
  return res;
}

int main() {
  vector<int> v = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
  vector<int> v1 = {2,0,3};
  vector<int> v2 = {2,0,1};
  vector<int> v3 = {1,1,1};
  cout << Solve(v) << endl;
  cout << Solve(v1) << endl;
  cout << Solve(v2) << endl;
  cout << Solve(v3) << endl;
}
