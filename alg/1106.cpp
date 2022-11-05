#include "debug.hpp"
using namespace std;
class Solution {
    int i = 0;
    int dfs(string expression, int level) {
        printf("level=%d, i=%d/%d, expression=%s\n", level, i, expression.size(), expression.substr(i).c_str());
        if (expression.size() <= i) {
            return 1;
        }
        if (expression[i] == '!') {
            i += 2;
            int ex = dfs(expression, level + 1);
            i += 1;
            return !ex;
        } else if (expression[i] == '&') {
            i += 2;
            int ex = 1;
            while (i < expression.size() && expression[i] != ')') {
                if (expression[i] == ',') i++;
                else {
                    int e = dfs(expression, level + 1);
                    //printf("i=%d, e = %d\n", i, e);
                    if (e == 0) {
                        ex = 0;
                    }
                }
            }
            i += 1;
            //printf("& expression:%d\n", ex);
            return ex;
        } else if (expression[i] == '|') {
            i += 2;
            int ex = 0;
            while (i < expression.size() && expression[i] != ')') {
                if (expression[i] == ',') {
                    i++;
                } else {
                    int e = dfs(expression, level + 1);
                    printf("\te=%d\n", e);
                    if (e > 0) {
                        ex = 1;
                    }
                }
            }
            i += 1; //eat )
            printf("| expression:%d, %s, %d\n", i, expression.substr(i).c_str(), ex);
            return ex;
        } else {
            assert(expression[i] == 't' or expression[i] == 'f');
            if (expression[i] == 't') {
                i += 1;
                return 1;
            } else {
                i += 1;
                return 0;
            }
        }
        printf("error returns\n");
        return 1;
    }
public:
    bool parseBoolExpr(string expression) {
        i = 0;
        return dfs(expression, 0);
    }
};
int main() {
    Solution sl;
    cout << sl.parseBoolExpr("!(&(&(!(&(f)),&(t),|(f,f,t)),&(t),&(t,t,f)))") << endl;
    /* cout << sl.parseBoolExpr("!(|(f,f,t))") << endl; */
    return 0;
}
