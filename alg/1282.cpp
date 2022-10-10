#include "stdc++.h"
using namespace std;
class Solution {
public:
    vector<vector<int>> groupThePeople(vector<int>& groupSizes) {
        unordered_map<int, vector<int>> g;
        int n= groupSizes.size();
        for (int i = 0; i < n; i++) {
            int x = groupSizes[i];
            g[x].push_back(i);
        }
        /* for (auto& [x,v]:g) { */
        /*     cout << x << endl; */
        /*     for (auto b : v) { */
        /*         cout << "\t" << b << " "; */
        /*     } */
        /*     cout << endl; */
        /* } */
        /* cout << "----"<<endl; */
        // len(g[x]) / x 
        vector<vector<int>> res;
        for (auto& [x, v]: g) {
            int ns = v.size() / x;
            for (int i = 0; i < ns; i++) {
                res.push_back(vector<int>{});
                for (int j = i * x; j < (i+1) * x; j++) {
                    res.back().push_back(v[j]);
                }
            }
            
        }
        /* cout << "end" << endl; */
        return res;
    }
};

int main() {
    Solution sl;

    cout << "running..."<<endl;
    /* vector<int> input{3,3,3,3,3,1,3}; */
    vector<int> input  {2,1,3,3,3,2};
    vector<vector<int>> res = sl.groupThePeople(input);
    for (auto a : res) {
        for (auto b : a)
            cout << b << " ";
        cout << endl;
    }
    return 0;
}
