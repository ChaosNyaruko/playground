#include "debug.hpp"
using namespace std;
class Solution {
  int dir[4][2] = {{-1, 0}, {1, 0}, {0, 1}, {0, -1}};

public:
  int shortestBridge(vector<vector<int>> &grid) {
    int n = grid.size();
    queue<int> q;
    queue<int> start;
    bool two = false;
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < n; j++) {
        if (grid[i][j] == 1) {
          q.push(i * n + j);
          grid[i][j] = 2;
          start.push(i * n + j);
          while (!q.empty()) {
            int x = q.front() / n;
            int y = q.front() % n;
            q.pop();
            for (int k = 0; k < 4; k++) {
              int nx = x + dir[k][0];
              int ny = y + dir[k][1];
              if (nx >= 0 and nx < n and ny >= 0 and ny < n and
                  grid[nx][ny] == 1) {
                grid[nx][ny] = 2;
                q.push(nx * n + ny);
                start.push(nx * n + ny);
              }
            }
          }
          two = true;
          break;
        }
      }
      if (two) {
        break;
      }
    }
    int step = 0;
    while (!start.empty()) {
      for (int _ = start.size(); _ > 0; _--) {
        int x = start.front() / n;
        int y = start.front() % n;
        /* cout << "--: " << start.front() << " " << step << endl; */
        start.pop();
        for (int k = 0; k < 4; k++) {
          int nx = x + dir[k][0];
          int ny = y + dir[k][1];
          if (nx >= 0 and nx < n and ny >= 0 and ny < n) {
            if (grid[nx][ny] == 0) {
              grid[nx][ny] = 3;
              start.push(nx * n + ny);
            } else if (grid[nx][ny] == 1) {
              return step;
            }
          }
        }
      }
      step++;
    }
    return 999;
  }
};

// 01 10 1
// 010 000 001 2
// 11111  10001 10101 10001 11111 1
int main() {
  vector<vector<vector<int>>> inputs = {{{0, 1}, {1, 0}},
                                        {{0, 1, 0}, {0, 0, 0}, {0, 0, 1}},
                                        {{1, 1, 1, 1, 1},
                                         {1, 0, 0, 0, 1},
                                         {1, 0, 1, 0, 1},
                                         {1, 0, 0, 0, 1},
                                         {1, 1, 1, 1, 1}}};
  Solution sl;
  for (auto &input : inputs) {
    cout << sl.shortestBridge(input) << endl;
  }
  return 0;
}
