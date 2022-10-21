/* 例如，如果未来7天股票的价格是 [100, 80, 60, 70, 60, 75, 85]，那么股票跨度将是
 * [1, 1, 1, 2, 1, 4, 6]。 */
#include "debug.hpp"
using namespace std;

/* 例如，如果未来7天股票的价格是 [100, 80, 60, 70, 60, 75, 85]，那么股票跨度将是
 * [1, 1, 1, 2, 1, 4, 6]。 */

class StockSpanner {
  stack<int> s;
  vector<int> v;

public:
  StockSpanner() {}

  int next(int price) {
    int i = v.size();
    v.push_back(price);
    while (!s.empty() and price >= v[s.top()]) {
      s.pop();
    }
    int res = s.empty() ? i + 1 : i - s.top();
    s.push(i);
    return res;
  }
};

int main() {
  StockSpanner sl;
  cout << sl.next(100) << endl;
  cout << sl.next(80) << endl;
  cout << sl.next(60) << endl;
  cout << sl.next(70) << endl;
  cout << sl.next(60) << endl;
  cout << sl.next(75) << endl;
  cout << sl.next(85) << endl;
}
