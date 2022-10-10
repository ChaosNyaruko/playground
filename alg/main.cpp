#include "stdc++.h"

using namespace std;
const int MOUSEWIN = 1;
const int CATWIN = 2; 
const int DRAW = 0;
class Solution {
    int dp[51][51][51];
    int n;
    vector<vector<int> > g;
    int getResult(int mouse, int cat, int turn) {
        if (turn >= 2*n) {
            return DRAW;
        }
        if (dp[mouse][cat][turn] >= 0) return dp[mouse][cat][turn];
        if (mouse == 0) {
            dp[mouse][cat][turn] = MOUSEWIN;
        } else if (cat == mouse) {
            dp[mouse][cat][turn] = CATWIN;
        } else{
            getNextResult(mouse, cat, turn);
        }

        return dp[mouse][cat][turn];
    }

    void getNextResult(int mouse, int cat, int turn) {
        int curMove = turn % 2 == 0? mouse: cat;
        int defaultResult = curMove == mouse? CATWIN: MOUSEWIN;
        int res = defaultResult;
        for (auto next : g[curMove]) {
            if (curMove == cat && next == 0) {
                continue;
            }
            int nextMouse = curMove == mouse? next: mouse;
            int nextCat = curMove == cat?next : cat;
            int nextResult = getResult(nextMouse, nextCat, turn + 1); 
            if (nextResult != defaultResult) {
                res = nextResult;
                if (nextResult != DRAW) {
                    break;
                }
            }
        }
        dp[mouse][cat][turn] = res;
    }

public:
    int catMouseGame(vector<vector<int>>& graph) {
        return getResult(1, 2, 0);
    }
};

int main() {
    int arr[2][3][4];
    cout << sizeof(arr) << " " << sizeof(int) << endl;
    cout<<"test"<<endl;
    Solution sl = Solution();
    vector<vector<int> > a = {{2,5}, {0,4,5}, {1,4,5}, {2,3}, {0, 2, 3}};
    printf("%d", sl.catMouseGame(a));
    return 0;
}
