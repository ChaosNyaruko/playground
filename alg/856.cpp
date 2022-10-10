#include "stdc++.h"
using namespace std;
class Solution {
public:
  int i = 0;
  int scoreOfParentheses(string s) {
    /* int n = s.size(); */
    /* int res = 0; */
    /* while (i < n and s[i] == '(') { */
    /*   i++; */
    /*   if (s[i] == ')') { */
    /*     res += 1; */
    /*   } else { */
    /*     res += 2 * scoreOfParentheses(s); */
    /*   } */
    /*   i++; */
    /* } */
    /* return res; */
    stack<int> stk;
    stk.push(0);
    for (auto c : s) {
      if (c == '(') {
        stk.push(0);
      } else {
        int v = stk.top();
        stk.pop();
        if (v == 0) {
          v = 1;
        } else {
          v = 2 * v;
        }
        stk.top() += v;
      }
    }
    return stk.top();
  }
};

int main() {
  Solution sl;
  vector<string> inputs = {"()", "(())", "()()", "(()(()))", "(())()"};
  for (auto i : inputs) {
    sl.i = 0;
    cout << "working on: " << i << endl;
    cout << sl.scoreOfParentheses(i) << endl;
  }
  return 0;
  stack<int> s;
  s.push(1);
  s.push(2);
  s.push(3);
}
