#include "debug.hpp"
using namespace std;
/* [2,7,9,4,4] */
/* M = 1 */

/* dfs(i, m) */
/* if i == n */
/* dfs() */

class Solution {
  int n;
  vector<vector<int>> memo;
  vector<int> prefix;
  int dfs(int i, int m) {
    /* printf("enter i:%d, n%d\n", i, n); */
    if (i >= n)
      return 0;
    if (memo[i][m] != -1)
      return memo[i][m];
    int right = 2 * m;
    int res = INT_MIN;
    for (int j = 1; j <= right; j++) {
      if (i + j > n)
        break;
      int get = prefix[n] - prefix[i];
      int opponent = dfs(i + j, max(j, m));
      res = max(res, get - opponent);
      printf("dfs(%d, %d), j = %d, get=%d, opponent(%d, %d)=%d, res=%d\n", i, m,
             j, get, i + j, max(m, j), opponent, res);
    }
    memo[i][m] = res;
    printf("get a res: dfs(%d, %d)=%d\n", i, m, res);
    return res;
  }

public:
  int stoneGameII(vector<int> &piles) {
    int n = piles.size();
    this->n = n;
    prefix = vector<int>(n + 1, 0);
    memo = vector<vector<int>>(n + 1, vector<int>(2 * n + 1, -1));
    for (int i = 0; i < n; i++) {
      prefix[i + 1] = prefix[i] + piles[i];
    }
    return dfs(0, 1);
  }
};

int main() {
  vector<int> v1 = {2, 7, 9, 4, 4};
  Solution sl;
  /* cout << sl.stoneGameII(v1) << endl; */
  v1 = {1, 2, 3, 4, 5, 100};
  cout << sl.stoneGameII(v1) << endl;
}
