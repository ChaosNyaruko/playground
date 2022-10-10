#include "debug.hpp"
using namespace std;
class Solution {
  int getMove(const int value, const int count, const int n) {
    printf("getMove: 0x%x, %d, %d \n", value, count, n);
    int ones = __builtin_popcount(value);
    if (n & 1) {
      if (abs(2 * count - n) > 1) {
        return -1;
      }
      if (abs(2 * ones - n) > 1) {
        return -1;
      }
      // 111110000 -> 101010101
      // 101010000
      //
      if (ones != (n + 1) / 2) { // 0 more than 1
        int delta = (value & 0xAAAAAAAA) ^ (value);
        printf("1 more than 0: 0x%x\n", delta);
        return __builtin_popcount(delta);
      } else {
        int delta = (value & 0x55555555) ^ (value);
        printf("0 more than 1: 0x%x\n", delta);
        return __builtin_popcount(delta);
      }
    } else {
      if (n / 2 != count) {
        return -1;
      }
      if (ones != n / 2) {
        return -1;
      }
      int a = (value & 0xAAAAAAAA) ^ (value);
      int b = (value & 0x55555555) ^ (value);
      printf("a:%x, b:%x\n", a, b);
      return min(__builtin_popcount(a), __builtin_popcount(b));
    }
  }

public:
  int movesToChessboard(vector<vector<int>> &board) {
    int n = board.size();
    int row = 0, col = 0;
    int rowCnt = 0, colCnt = 0;
    for (int i = 0; i < n; i++) {
      row |= (board[0][i] << i);
      col |= (board[i][0] << i);
    }
    int rRow = (row ^ ((1 << n) - 1));
    int rCol = (col ^ ((1 << n) - 1));
    printf("%x %x, %x %x\n", row, col, rRow, rCol);
    for (int i = 0; i < n; i++) {
      int curRow = 0;
      int curCol = 0;
      for (int j = 0; j < n; j++) {
        curRow |= (board[i][j] << j);
        curCol |= (board[j][i] << j);
      }
      printf("curRow:%x\n", curRow);
      // row detect
      if (curRow == row) {
        rowCnt++;
      } else if (curRow == rRow) {

      } else {
        return -1;
      }
      // col detect
      if (curCol == col) {
        colCnt++;
      } else if (curCol == rCol) {

      } else {
        return -1;
      }
    }
    int r = getMove(row, rowCnt, n);
    int c = getMove(col, colCnt, n);
    return r == -1 || c == -1 ? -1 : r + c;
  }
};

int main() {
  // 0
  vector<vector<int>> input{{1, 0}, {0, 1}};
  // -1
  /* vector<vector<int>> input{{1, 0}, {1, 0}}; */
  // 2
  /* vector<vector<int>> input{ */
  /*     {0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}}; // result: 2
   */
  /* vector<vector<int>> input{{1, 1, 0}, {0, 0, 1}, {0, 0, 1}}; // result: 2 */
  Solution sl;
  cout << sl.movesToChessboard(input) << endl;
}
