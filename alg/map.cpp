#include "debug.hpp"
using namespace std;
class Solution {
    // 数位dp/模板
    unordered_map<pair<int, int>, int> dp;
    int dfs(const string& s, int i, int cnt, bool limit) {
        if (i == s.size()) {
            return cnt;
        }
        int m = limit ? s[i] - '0' : 9;
        int res = 0;
        for (int j = 0; j <= m; j++) {
            res += dfs(s, i + 1, cnt + (j == 1 ? 1:0), limit && (j == s[i] - '0'));
        }
        return res;
    }
public:
    int countDigitOne(int n) {
        return dfs(to_string(n), 0, 0, true);
    }
};
int main() {
  Solution sl;
  cout << sl.countDigitOne(13) << endl;
  return 0;
}
