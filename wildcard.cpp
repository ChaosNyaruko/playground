#include "stdc++.h"
using namespace std;
bool isMatch(string s, string p) {

  int i = 0, j = 0, iStar = -1, jStar = -1;

  while (i < s.length()) {

    if (s[i] == p[j] || p[j] == '?')

      i++, j++;

    else if (p[j] == '*') {

      iStar = i;

      jStar = j++;

    } else if (iStar >= 0) {

      i = ++iStar;

      j = jStar + 1;

    } else {

      return false;
    }
    printf("i:%d, j:%d, iStar:%d, jStar:%d\n", i, j, iStar, jStar);
  }

  while (j < p.length() && p[j] == '*')
    j++;

  return j == p.length();
}
int main() { cout << isMatch("aa", "*") << endl; }
