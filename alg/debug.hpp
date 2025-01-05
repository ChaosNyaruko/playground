#include "stdc++.h"

inline void print(const std ::vector<int> &vec) {
  for (auto &v : vec) {
    printf("%d ", v);
  }
  printf("\n");
}

inline void print(std ::vector<int> &&vec) {
  for (auto &v : vec) {
    printf("%d ", v);
  }
  printf("\n");
}

inline void print(std ::vector<std::string> &&vec) {
  for (auto &v : vec) {
    printf("%s", v.c_str());
  }
  printf("\n");
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

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

void print(TreeNode* node) {
  if (node == nullptr) {
    printf("null,");
    return;
  }
  int val = node->val;
  printf("%d,", val);
  print(node->left);
  print(node->right);
  return;
}


/* struct key_hash : public std::unary_function<key_t, std::size_t> */
/* { */
/*    std::size_t operator()(const key_t& k) const */
/*    { */
/*       return std::get<0>(k)[0] ^ std::get<1>(k) ^ std::get<2>(k); */
/*    } */
/* }; */

/* struct key_equal : public std::binary_function<key_t, key_t, bool> */
/* { */
/*    bool operator()(const key_t& v0, const key_t& v1) const */
/*    { */
/*       return ( */
/*                std::get<0>(v0) == std::get<0>(v1) && */
/*                std::get<1>(v0) == std::get<1>(v1) && */
/*                std::get<2>(v0) == std::get<2>(v1) */
/*              ); */
/*    } */
/* }; */

/* struct data */
/* { */
/*    std::string x; */
/* }; */

/* typedef std::unordered_map<key_t,data,key_hash,key_equal> map_t; */

