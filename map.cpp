#include "stdc++.h"

using namespace std;

int main() {
    unordered_map<char, vector<int>> c;
    for (auto&& [_, arr] : c) {
        cout << arr.size() << endl;
    }
    return 0;
}
