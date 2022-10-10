#include <iostream>
#include <vector>
using namespace std;

bool g(vector<int>& a, vector<int>& v, int m, int n, int b) {
    if (m == 0 && n == 0) return true;
    if (n == 0) return g(a, v, m - 1, b, b);
    for (int i = 0; i < a.size(); i++) {
        if (v[i] || a[i] > n) continue;
        v[i]++;
        if (g(a, v, m, n - a[i], b)) return true;
        v[i]--;
    }
    return false;
}

bool f(vector<int> a) {
    int sum = 0;
    for (int n: a) sum += n;
    if (sum % 4 != 0) return false;
    int b = sum / 4;
    vector<int> v(a.size());
    return g(a, v, 4, 0, b);
}

int main() {
    //int a;
    //cin >> a;
    cout << "Hello World!" << endl;
    vector<int> a{1, 1, 2, 2, 2};
    cout << f(a) << endl;
    vector<int> b{3, 3, 3, 3, 4};
    cout << f(b) << endl;
    vector<int> c{1569462,2402351,9513693,2220521,7730020,7930469,1040519,5767807,876240,350944,4674663,4809943,8379742,3517287,8034755};
    vector<int> d{5,5,5,5,16,4,4,4,4,4,3,3,3,3,4};
    cout << f(c) << endl;
    cout << f(d) << endl;
}
