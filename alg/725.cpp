#include "debug.hpp"
using namespace std;
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    vector<ListNode*> splitListToParts(ListNode* head, int k) {
        int l = 0;
        ListNode* h = head;
        while (h) {
            l++;
            h = h->next;
        }
        int n = l / k;
        int remainder = l - k * n;
        ListNode* prev =nullptr;
        //printf("len: %d, n: %d, remainder: %d\n", l, n, remainder);
        vector<ListNode*> res(k, nullptr);
        for (int i = 0; i < k; i++) {
            //printf("constructing: %d, %d\n", i, head->val);
            for (int j =0 ; j < n; j++) {
                if (res[i] == nullptr) res[i] = head;
                prev = head;
                head = head->next;
            }
            for (auto x = res[i]; x; x=x->next) {
                //printf("\t%d", x->val);
            }
            if (remainder>0) {
                if (res[i] == nullptr) res[i] = head;
                remainder--;
                //printf("pushing remainder: %d\n", head->val);
                prev = head;
                head = head->next;
            } 
            if (prev)
                prev->next = nullptr;
            if (!head) break;
        }
        assert(head == nullptr);
        return res;
    }
};
