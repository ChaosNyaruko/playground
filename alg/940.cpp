#include "debug.hpp"
using namespace std;
class Solution {
public:
  int distinctSubseqII(string s) {
    // dp[i][j] 0..i-1, end with a-z disSII?
    // how to ensure the distinctness
    // f[i] counts for sequences ending with s[i]
    // f[i] >= f[j] if s[i] == s[j] and i > j
    // f[i] = 1 + sum_(k = a-z)_f[last[k]]
    //
    int n = s.size();
    int mod = 1e9 + 7;
    int res = 0;
    // solution 1: time complexity O(n * 26)
    vector<int> f(n, 1);
    vector<int> last(26, -1);
    for (int i = 0; i < n; i++) {
      for (int k = 0; k < 26; k++) {
        if (last[k] != -1) {
          f[i] = (f[i] + f[last[k]]) % mod;
        }
      }
      last[s[i] - 'a'] = i;
    }
    for (int i = 0; i < 26; i++) {
      if (last[i] != -1) {
        res = (res + f[last[i]]) % mod;
      }
    }
    return res;

    vector<int> g(26, 0);
    long long total = accumulate(g.begin(), g.end(), 0);
    for (int i = 0; i < n; i++) {
      int g_next = (1 + total) % mod;
      total = (total + (g_next - g[s[i] - 'a'])) % mod;
      g[s[i] - 'a'] = g_next;
    }

    for (int i = 0; i < 26; i++) {
      res = (res + g[i]) % mod;
    }
    return res;
  }
};

int main() {
  Solution sl;
  vector<string> inputs = {
      "abc", "aba", "aaa",
      "zchmliaqdgvwncfatcfivphddpzjkgyygueikthqzyeeiebczqbqhdytkoawkehkbizdmcni"
      "lcjjlpoeoqqoqpswtqdpvszfaksn",
      "knqmywztzgalovcyitifjmllyltjjnwbehsqaofidwzygekdylwmwxtsnhowpyuwkxomdqsl"
      "dbcuseojgyimebpvqyzmvubgwhku"}; // expected 7 6 3 97915677
  for (auto s : inputs) {
    cout << sl.distinctSubseqII(s) << endl;
  }
  return 0;
}
