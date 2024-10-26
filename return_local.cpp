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

  int* p = new int[3];
  p[0] = 1;
  printf("%d, %d, %d\n", p[0], p[1], p[2]);
  delete[] p;
  p = nullptr;
  delete[] p;


    // a lot of work
  return 0;
}
