#include "stdc++.h"

int foo();
int add(int index) {
    index = index + 1;
    return index;
}

int main() {
    std::vector<int> res = {};
    for (auto x:res) {
        std::cout << x << std::endl;
    }
    std::cout << "hello world" << std::endl;
    printf("%d\n", foo());
    printf("%d\n", add(1));
    return 0;
}

int foo() {
    return 20;
}

