#include "stdc++.h"

using namespace std;

int main() {
  cout << "hello" << endl;
  unordered_map<char, vector<int>> c;
  c['a'] = {1, 2, 3, 5};
  auto tmp = c['a'];
  for (auto &&[_, arr] : c) {
    cout << arr.size() << endl;
  }
  for (auto &&[_, arr] : c) {
    cout << arr[0] << endl;
  }
  c['a'][0] = 2;

  for (auto &&[_, arr] : c) {
    cout << arr[0] << endl;
  }
  cout << tmp[0] << endl;
  return 0;
}
