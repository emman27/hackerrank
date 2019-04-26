#include <algorithm>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <set>
#include <vector>
using namespace std;

int main() {
  /* Enter your code here. Read input from STDIN. Print output to STDOUT */
  int queries;
  cin >> queries;
  set<int> s;
  for (auto i = 0; i < queries; i++) {
    int queryType, val;
    cin >> queryType >> val;
    if (queryType == 1) {
      s.insert(val);
    } else if (queryType == 2) {
      s.erase(val);
    } else if (s.find(val) != s.end()) {
      cout << "Yes" << endl;
    } else {
      cout << "No" << endl;
    }
  }
  return 0;
}
