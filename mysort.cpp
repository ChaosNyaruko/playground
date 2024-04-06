// input string sort 
// aab aba

#include <iostream>
#include <queue>
#include <vector>
#include <string>
using namespace std;
string mySort(string s) {
  if (s.length() < 2) {
    return s;
  }

  int len = s.length();
  vector<int> cnt(26,0);
  int cntMax = 0;
  for (int i = 0 ; i < s.length(); i++) {
    cnt[s[i] - 'a']++;
    cntMax = max(cntMax, cnt[s[i] - 'a']);
  }
  // not work for "abaa"
  if (cntMax > (len + 1) / 2) {
    return "";
  }

  auto cmp = [&](char a1, char a2) -> bool {
    return cnt[a1 - 'a'] < cnt[a2 - 'a'];
  };
  priority_queue<char, vector<char>, decltype(cmp)> q(cmp);
  for (char c = 'a'; c <= 'z'; c++) {
    if (cnt[c - 'a'] > 0) {
      q.push(c);
    }
  }
  string res;
  while (q.size() > 1) {
    // fetch two letters?
    char first = q.top();
    q.pop();
    char second = q.top();
    q.pop();
    cnt[first - 'a']--;
    cnt[second - 'a']--;
    res += first;
    res += second;
    if (cnt[first - 'a'] > 0) {
      q.push(first);
    }
    if (cnt[second - 'a'] > 0) {
      q.push(second);
    }
  }
  if (!q.empty()) {
    // assert q.size() == 1
    res += q.top();
    q.pop();
  }
  return res;
}

int main() {
  cout << mySort("aba")<<endl;
  cout << mySort("aab")<<endl;
  cout << mySort("abaa")<<endl;
  cout << mySort("ab")<<endl;
  cout << mySort("a")<<endl;
  cout << mySort("aa")<<endl;
  cout << mySort("abcefaaaa")<<endl;
}
