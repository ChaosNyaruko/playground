#include "debug.hpp"
using namespace std;
class Solution {
public:
    int kthGrammar(int n, int k) {
        if (n == 1 or k == 1) {
            return 0;
        }
        if (k > (1 << (n-2))) {
            return 1 ^ kthGrammar(n-1, k - (1 << (n-2)));
        }
        return kthGrammar(n-1,k);
    }
};

int main() {
    Solution sl;
    vector<vector<int>> inputs = {{1,1},{2,1},{2,2}};
    for (auto& i : inputs) {
        cout << sl.kthGrammar(i[0], i[1]) << endl;
    }
}

    //  0
    // 0 1
   // 01 10
// 0110 1001
// 0110 1001 1001 0110
// 0110 1001 1001 0110 1001 0110 0110 1001
// the kth element in line n
// 2**(n-1) == 1 << (n-1)
