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
