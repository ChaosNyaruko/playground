#include "stdc++.h"

int* foo(int x) {
  int a = x;
  return &a;
}

int main() {
  auto p1 = foo(1);
  auto p2 = foo(2);
  printf("%d\n", *p1);
  printf("%d\n", *p2);
  return 0;
}
