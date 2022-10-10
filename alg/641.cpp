#include "stdc++.h"
using namespace std;
class MyCircularDeque {
  int cap;
  int front, rear;
  vector<int> q;

public:
  MyCircularDeque(int k) {
    cap = k + 1;
    front = rear = 0;
    q = vector<int>(k + 1, 0);
  }

  bool insertFront(int value) {
    if (isFull())
      return false;
    front = (front - 1 + cap) % cap;
    q[front] = value;
    return true;
  }

  bool insertLast(int value) {
    if (isFull()) {
      return false;
    }
    q[rear] = value;
    rear = (rear + 1) % cap;
    return true;
  }

  bool deleteFront() {
    if (isEmpty()) {
      return false;
    }
    front = (front + 1) % cap;
    return true;
  }

  bool deleteLast() {
    if (isEmpty())
      return false;
    rear = (rear - 1 + cap) % cap;
    return true;
  }

  int getFront() {
    if (isEmpty())
      return -1;
    return q[front];
  }

  int getRear() {
    if (isEmpty())
      return -1;
    return q[(rear - 1) % cap];
  }
  bool isEmpty() { return rear == front; }

  bool isFull() { return (rear + 1) % cap == front; }
};

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * MyCircularDeque* obj = new MyCircularDeque(k);
 * bool param_1 = obj->insertFront(value);
 * bool param_2 = obj->insertLast(value);
 * bool param_3 = obj->deleteFront();
 * bool param_4 = obj->deleteLast();
 * int param_5 = obj->getFront();
 * int param_6 = obj->getRear();
 * bool param_7 = obj->isEmpty();
 * bool param_8 = obj->isFull();
 * ["MyCircularDeque","insertFront","getRear","insertFront","getRear","insertLast","getFront","getRear","getFront","insertLast","deleteLast","getFront"]
    [[3],[9],[],[9],[],[5],[],[],[],[8],[],[]]
 */
/* ["MyCircularDeque","insertFront","getRear","insertFront","getRear","insertLast","getFront","getRear","getFront","insertLast","deleteLast","getFront"]
 */
/* [[3],[9],[],[9],[],[5],[],[],[],[8],[],[]] */
int main() {
  MyCircularDeque *obj = new MyCircularDeque(3);
  cout << obj->insertFront(9);
  cout << obj->getRear();
  cout << obj->insertFront(9);
  cout << obj->getRear();
  cout << obj->insertLast(5);
  cout << obj->getFront();
  cout << obj->getRear();
  cout << obj->getFront();
  cout << obj->insertLast(8);
  cout << obj->deleteLast();
  cout << obj->getFront();
  cout << endl;

  return 0;
}
