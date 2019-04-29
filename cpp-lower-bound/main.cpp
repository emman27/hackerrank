#include <algorithm>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <vector>
using namespace std;

int main() {
  /* Enter your code here. Read input from STDIN. Print output to STDOUT */
  int len, queries;
  cin >> len;
  vector<int> v;
  int x;
  for (auto i = 0; i < len; i++) {
    cin >> x;
    v.push_back(x);
  }
  cin >> queries;
  int q;
  for (auto i = 0; i < queries; i++) {
    cin >> q;
    auto it = lower_bound(v.begin(), v.end(), q);
    if (*it == q) {
      cout << "Yes ";
    } else {
      cout << "No ";
    }
    cout << it - v.begin() + 1 << endl;
  }
  return 0;
}
