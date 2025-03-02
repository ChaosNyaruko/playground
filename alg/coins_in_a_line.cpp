#include "debug.hpp"
using namespace std;

// ref implement.
const int MAX_N = 100;
void printMoves(int P[][MAX_N], int A[], int N) {
  int sum1 = 0, sum2 = 0;
  int m = 0, n = N - 1;
  bool myTurn = true;
  while (m <= n) {
    int P1 = P[m + 1][n]; // If take A[m], opponent can get...
    int P2 = P[m][n - 1]; // If take A[n]
    cout << (myTurn ? "I" : "You") << " take coin no. ";
    if (P1 <= P2) {
      cout << m + 1 << " (" << A[m] << ")";
      m++;
    } else {
      cout << n + 1 << " (" << A[n] << ")";
      n--;
    }
    cout << (myTurn ? ", " : ".\n");
    myTurn = !myTurn;
  }
  cout << "\nThe total amount of money (maximum) I get is " << P[0][N - 1]
       << ".\n";
}

int maxMoney(int A[], int N) {
  int P[MAX_N][MAX_N] = {0};
  int a, b, c;
  for (int i = 0; i < N; i++) {
    for (int m = 0, n = i; n < N; m++, n++) {
      assert(m < N);
      assert(n < N);
      a = ((m + 2 <= N - 1) ? P[m + 2][n] : 0);
      b = ((m + 1 <= N - 1 && n - 1 >= 0) ? P[m + 1][n - 1] : 0);
      c = ((n - 2 >= 0) ? P[m][n - 2] : 0);
      P[m][n] = max(A[m] + min(a, b), A[n] + min(b, c));
    }
  }
  printMoves(P, A, N);
  return P[0][N - 1];
}

void print_moves(int *coins, int n, vector<vector<int>> &memo) {
  int l = 0, r = n - 1;
  bool my_turn = true;
  while (l <= r) {
    // I take l, oppo get ..., I'm try to minimise it.
    int p1 = (l + 1 <= r) ? memo[l + 1][r] : 0;
    /* int p1 = memo[l+1][r]; */
    /* int p2 = memo[l][r - 1]; */
    int p2 = (l <= r - 1) ? memo[l][r - 1] : 0;
    int taken = l;
    if (p1 <= p2) {
      taken = l;
      l++;
    } else {
      taken = r;
      r--;
    }
    cout << (my_turn ? "I" : "You") << " took " << taken + 1 << endl;
    my_turn = !my_turn;
  }
}

void print_coins(int *a, int n) {
  for (int i = 0; i < n; i++) {
    std::cout << a[i] << " ";
  }
  std::cout << std::endl;
}

int max_money(int **S, int P[], const int N) {
  vector<vector<int>> memo(N, vector<int>(N, -1));
  for (int len = 1; len <= N; len++) {
    for (int i = 0; i <= N - len; i++) {
      int j = i + len - 1;
      if (len == 1) {
        memo[i][j] = P[i];
      } else {
        assert(memo[i + 1][j] != -1 && "i+1, j should not be -1");
        assert(memo[i][j - 1] != -1 && "i, j-1  should not be -1");
        memo[i][j] = S[i][j] - min(memo[i + 1][j], memo[i][j - 1]);
      }
    }
  }
  print_moves(P, N, memo);
  return memo[0][N - 1];
}

int main(void) {
  while (1) {
    int n;
    std::cin >> n;
    int *coins = new int[n];
    for (int i = 0; i < n; i++) {
      std::cin >> coins[i];
    }
    print_coins(coins, n);
    int **S = new int *[n];
    for (int i = 0; i < n; i++) {
      S[i] = new int[n];
    }
    int *prefix = new int[n + 1];
    prefix[0] = 0;
    for (int i = 0; i < n; i++) {
      prefix[i + 1] = prefix[i] + coins[i];
    }
    for (int i = 0; i < n; i++) {
      for (int j = i; j < n; j++) {
        S[i][j] = prefix[j + 1] - prefix[i];
      }
    }
    cout << "my   result: " << max_money(S, coins, n) << endl;
    cout << "refs result: " << maxMoney(coins, n) << endl;
  }
  return 0;
}
