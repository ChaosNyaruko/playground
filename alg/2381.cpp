#include "debug.hpp"
using namespace std;
class Solution {
    char shiftChar(char c, int x) {
        x = x % 26;
        if (x < 0) {
            x = 26 + x;
        }
        printf("shift %c by %d\n", c, x);
        return (c - 'a' + x) % 26 + 'a';
    }
public:
    string shiftingLetters(string s, vector<vector<int>>&& shifts) {
        int n = s.size();
        vector<int> diff(n);
        for (auto& s: shifts) {
            int b = s[0];
            int e = s[1];
            int d = s[2] == 1 ? 1 : -1;
            diff[b] += d;
            if (e + 1 < n) {
                diff[e + 1] -= d;
            }
        }
        print(diff);
        int x = 0;
        for (int i = 0; i < n; i++) {
            x += diff[i];
            s[i] = shiftChar(s[i], x);
        }
        return s;
    }
};

int main() {
  Solution sl;
  
  cout << sl.shiftingLetters("abc", {{0, 1, 0}, {1,2,1},{0,2,1}}) << endl;
  cout << sl.shiftingLetters("dztz", {{0, 0, 0}, {1,1,1}}) << endl;
}
