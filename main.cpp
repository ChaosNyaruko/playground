#include "stdc++.h"

int foo();
int main() {
    std::vector<int> res = {};
    for (auto x:res) {
        std::cout << x << std::endl;
    }
    std::cout << "hello world" << std::endl;
    printf("%d", foo());
    return 0;
}

int foo() {
    return 20;
}

