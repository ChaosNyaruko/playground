#include "stdc++.h"

inline void print(std ::vector<int> &&vec) {
  for (auto &v : vec) {
    printf("%d\n", v);
  }
}

inline void print(std ::vector<std::string> &&vec) {
  for (auto &v : vec) {
    printf("%s\n", v.c_str());
  }
}

struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *build_list(const std::vector<int> &vec) {
  ListNode *dummy = new ListNode(0, nullptr);
  ListNode *cur = dummy;
  for (auto i : vec) {
    cur->next = new ListNode(i, nullptr);
    cur = cur->next;
  }
  return dummy->next;
}
