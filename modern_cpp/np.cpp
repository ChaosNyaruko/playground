#include<iostream>

namespace std{
    int out() {
        return 0;
    }
    int endl = 0;
}

int main(){
    std::cout << std::out() << std::endl;
}
