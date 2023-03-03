#include "debug.hpp"
using namespace std;

class S {
public:
  string a;
  /* S(string i) { */
  /*   cout << "int i used" << endl; */
  /*   a = i; */
  /* } */
  S(string &i) {
    cout << "int& i used" << endl;
    a = i;
  }
  S(string &&i) {
    cout << "int&& used" << endl;
    a = i;
  }
  S(const S &s) { cout << "copy constructor called!" << endl; }
  S(S &&s) { cout << "move constructor called!" << endl; }
  bool operator==(const S &other) const { return true; }
};

template <> struct hash<S> {
  std::size_t operator()(const S &k) const { return 1; }
};

S &&t() { return S("test"); }

string f() { return "()"; }
int g() { return 1; }

int main() {
  vector<string> res;
  res.push_back(f());
  unordered_map<string, int> index;
  index[f()] = 1;

  // refer to an object on stack, undefined behaviour, just as the warning says
  // different -Oxx result in different outputs.
  cout << "res[0]" << res[0] << endl;
  for (auto &[k, v] : index) {
    cout << "kv " << k << " " << v << endl;
  }
  return 0;
  cout << "====" << endl;
  vector<S> vs;
  unordered_map<S, int> sindex;
  vs.push_back(t());
  for (auto &&x : vs) {
    cout << "vs[0]=" << x.a << endl;
  }
  sindex[t()] = 1;
  /* for (auto [k, v] : sindex) { */
  /*   printf("%s:%d\n", k.a.c_str(), v); */
  /* } */
  cout << sindex[t()] << endl;
}
