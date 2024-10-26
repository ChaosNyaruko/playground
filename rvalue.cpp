#include "debug.hpp"
using namespace std;
class Solution {
 std::vector<std::string> split(const std::string& oris, const std::string& delimiter) {
        std::vector<std::string> tokens;
        size_t pos = 0;
        std::string token;
        string s = oris;
        while ((pos = s.find(delimiter)) != std::string::npos) {
            token = s.substr(0, pos);
            tokens.push_back(token);
            s.erase(0, pos + delimiter.length());
        }
        tokens.push_back(s);
    
        return tokens;
    }
    bool helper(const vector<string>& s1, const vector<string>& s2) {
        int l = -1, r = s1.size();
        int rc = s2.size();
        while (l+1 < s2.size() && s1[l+1] == s2[l+1]) {
            l++;
        }
        if (l == s2.size()) {
            return true;
        }
        while (rc - 1 > l && s1[r - 1] == s2[rc - 1]) {
            r--;
            rc--;
        }
        return l + rc == s2.size() - 1;
    }
public:
    bool areSentencesSimilar(string sentence1, string sentence2) {
        vector<string>&& s1 = split(sentence1, " ");        
        vector<string>&& s2 = split(sentence2, " ");        
        if (s1.size() < s2.size()) {
            return helper(s2, s1);
        }
        return helper(s1,s2);
    }
};
