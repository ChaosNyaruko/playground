#include "debug.hpp"
using namespace std;
class Solution {
  vector<string> putDot(const string &a) {
    // cout << "putDot" << a << endl;
    vector<string> res;
    if (a[0] != '0' or a == "0") {
      res.push_back(a);
    }
    if (a.back() == '0') {
      return res;
    }
    for (int i = 1; i < a.size(); i++) {
      if (a[0] == '0' and i > 1)
        continue;
      if (a.back() == '0')
        continue;
      res.push_back(a.substr(0, i) + "." + a.substr(i));
    }
    return res;
  }

public:
  vector<string> ambiguousCoordinates(string s) {
    int n = s.size();
    vector<string> res;
    for (int i = 1; i < s.size() - 2; i++) {
      vector<string> &&a = putDot(s.substr(1, i));
      vector<string> &&b = putDot(s.substr(i + 1, s.size() - i - 2));
      for (auto &i : a) {
        for (auto &j : b) {
          res.push_back("(" + i + ", " + j + ")");
        }
      }
    }
    return res;
  }
};

int main() {
  Solution sl;
  vector<string> inputs = {"(123)", "(00011)", "(0123)", "(100)"};
  for (auto input : inputs) {
    vector<string> &&res = sl.ambiguousCoordinates(input);
    for (auto &&r : res) {
      cout << r << endl;
    }
    cout << endl;
  }
  return 0;
}
