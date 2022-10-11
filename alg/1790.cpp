#include "stdc++.h"
using namespace std;
class Solution {
public:
  bool areAlmostEqual(string s1, string s2) {
    if (s1.length() != s2.length()) {
      return false;
    }
    int differ = 0;
    unordered_map<char, char> m;
    for (int i = 0; i < s1.length(); i++) {
      if (s1[i] != s2[i]) {
        differ++;
        if (differ > 2)
          return false;
        m[s1[i]] = s2[i];
      }
    }
    if (m.size() == 2) {
      for (auto &[k, v] : m) {
        cout << k << " " << v << endl;
        if (k != m[v])
          return false;
      }
      return true;
    }
    return m.size() == 0;
  }
};

int main() {
  Solution sl;
  vector<vector<string>> inputs = {{"alg", "lag"},   {"abc", "def"},
                                   {"bank", "kanb"}, {"attack", "defend"},
                                   {"kelb", "kelb"}, {"abcd", "dcba"},
                                   {"acd", "bcd"}};
  for (auto &s : inputs) {
    cout << sl.areAlmostEqual(s[0], s[1]) << endl;
  }
}

class RefSolution {
public:
  bool areAlmostEqual(string s1, string s2) {
    int n = s1.size();
    vector<int> diff;
    for (int i = 0; i < n; ++i) {
      if (s1[i] != s2[i]) {
        if (diff.size() >= 2) {
          return false;
        }
        diff.emplace_back(i);
      }
    }
    if (diff.size() == 0) {
      return true;
    }
    if (diff.size() != 2) {
      return false;
    }
    return s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]];
  }
};
