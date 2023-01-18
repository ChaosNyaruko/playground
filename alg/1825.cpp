#include "debug.hpp"
using namespace std;
/* 给你两个整数 m 和 k ，以及数据流形式的若干整数。你需要实现一个数据结构，计算这个数据流的
 * MK 平均值 。 */

/* MK 平均值 按照如下步骤计算： */

/* 如果数据流中的整数少于 m 个，MK 平均值 为 -1 ，否则将数据流中最后
 * m 个元素拷贝到一个独立的容器中。 */
/* 从这个容器中删除最小的 k 个数和最大的 k 个数。 */
/* 计算剩余元素的平均值，并 向下取整到最近的整数 。 */
/* 请你实现 MKAverage 类： */

/* MKAverage(int m, int k) 用一个空的数据流和两个整数 m 和
 * k 初始化 MKAverage 对象。 */
/* void addElement(int num) 往数据流中插入一个新的元素 num 。 */
/* int calculateMKAverage() 对当前的数据流计算并返回 MK 平均数 ，结果需
 * 向下取整到最近的整数 。 */
class MKAverage {
  deque<int> ori;
  int sum;
  int ksum;
  multiset<int> q;
  multiset<int> preK;
  multiset<int> postK;
  int m;
  int k;

public:
  MKAverage(int m, int k) {
    ori.clear();
    q.clear();
    sum = 0;
    ksum = 0;
    this->m = m;
    this->k = k;
  }

  void print(int a) {
    printf("---after add %d---\nksum:%d, sum:%d\n", a, ksum, sum);
    printf("ori:\n");
    for (auto v : ori) {
      printf(" %d", v);
    }
    printf("\n");
    printf("preK:\n");
    for (auto v : preK) {
      printf(" %d", v);
    }
    printf("\n");
    printf("q:\n");
    for (auto v : q) {
      printf(" %d", v);
    }
    printf("\n");
    printf("postK:\n");
    for (auto v : postK) {
      printf(" %d", v);
    }
    printf("\n");
  }
  void addElement(int num) {
    /* cout << "adding:  " << num << endl; */
    if (ori.size() < m) {
      ori.push_back(num);
      sum += num;
      q.insert(num);
      if (ori.size() == m) {
        int i = 0;
        for (auto iter = q.begin(); i < k; iter = q.erase(iter), i++) {
          ksum += *iter;
          preK.insert(*iter);
        }
        i = 0;
        for (auto iter = prev(q.end()); i < k;
             iter = prev(q.erase(iter)), i++) {
          ksum += *iter;
          postK.insert(*iter);
        }
      }
      print(num);
      return;
    }
    int v = ori.front();
    ori.pop_front();
    ori.push_back(num);
    sum += num - v;

    int part = 0;
    multiset<int>::iterator iter = preK.find(v);
    if (iter == preK.end()) {
      iter = postK.find(v);
      if (iter == postK.end()) {
        iter = q.find(v); // it won't be q.end()
        part = 1;         // it beblongs to q
      } else {
        part = 2;
      }
    } else {
      part = 0;
    }
    if (part == 0) {
      preK.erase(iter);
      preK.insert(*q.begin());
      ksum = ksum - v + *q.begin();
      q.erase(q.begin());
    } else if (part == 1) {
      q.erase(iter);
    } else {
      postK.erase(iter);
      postK.insert(*q.rbegin());
      ksum = ksum - v + *q.rbegin();
      q.erase(prev(q.end()));
    }

    if (num >= *preK.rbegin() and num <= *postK.begin()) {
      q.insert(num);
    } else if (num < *preK.rbegin()) {
      ksum += num - *preK.rbegin();
      q.insert(*preK.rbegin());
      preK.erase(prev(preK.end()));
      preK.insert(num);
    } else {
      ksum += num - *postK.begin();
      q.insert(*postK.begin());
      postK.erase(postK.begin());
      postK.insert(num);
    }
    print(num);
  }

  int calculateMKAverage() {
    if (ori.size() < m) {
      return -1;
    }
    assert(m == ori.size());
    return (sum - ksum) / (m - 2 * k);
  }
};
/**
 * Your MKAverage object will be instantiated and called as such:
 * MKAverage* obj = new MKAverage(m, k);
 * obj->addElement(num);
 * int param_2 = obj->calculateMKAverage();
 */

int main() {
  MKAverage *obj = new MKAverage(3, 1);
  obj->addElement(3);
  obj->addElement(1);
  cout << obj->calculateMKAverage() << endl;
  obj->addElement(5);
  obj->addElement(5);
  obj->addElement(5);
  cout << obj->calculateMKAverage() << endl;
  /* vector<int> v{1, 2, 3, 4, 5}; */
  /* cout << v.end() - v.begin() << endl; */
  /* int a = 1; */
  /* a = a; */
  /* ["MKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement"]
   */
  /* [[3,1],[17612],[74607],[],[8272],[33433],[],[15456],[64938],[],[99741]] */
}
