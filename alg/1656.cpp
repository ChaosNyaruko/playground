#include "debug.hpp"

using namespace std;
class OrderedStream {
  vector<string> stream;
  int ptr;

public:
  OrderedStream(int n) {
    ptr = 1;
    stream.resize(n + 1);
  }

  vector<string> insert(int idKey, string value) {
    stream[idKey] = value;
    vector<string> res;
    for (; ptr < stream.size() and stream[ptr] != ""; ptr++) {
      res.emplace_back(stream[ptr]);
    }
    return res;
  }
};

/**
 * Your OrderedStream object will be instantiated and called as such:
 * OrderedStream* obj = new OrderedStream(n);
 * vector<string> param_1 = obj->insert(idKey,value);
 *
 */

int main() {
  OrderedStream *obj = new OrderedStream(5);
  print(obj->insert(3, "ccccc"));
  print(obj->insert(1, "aaaaa"));
  print(obj->insert(2, "bbbbb"));
  print(obj->insert(5, "eeeee"));
  print(obj->insert(4, "ddddd"));
}

/* OrderedStream os= new OrderedStream(5); */
/* os.insert(3, "ccccc"); // 插入 (3, "ccccc")，返回 [] */
/* os.insert(1, "aaaaa"); // 插入 (1, "aaaaa")，返回 ["aaaaa"] */
/* os.insert(2, "bbbbb"); // 插入 (2, "bbbbb")，返回 ["bbbbb", "ccccc"] */
/* os.insert(5, "eeeee"); // 插入 (5, "eeeee")，返回 [] */
/* os.insert(4, "ddddd"); // 插入 (4, "ddddd")，返回 ["ddddd", "eeeee"] */
