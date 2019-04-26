#include <algorithm>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <map>
#include <vector>
using namespace std;

int main() {
  /* Enter your code here. Read input from STDIN. Print output to STDOUT */
  int queries;
  cin >> queries;

  map<string, int> m;
  for (int i = 0; i < queries; i++) {
    int queryType, val;
    string name;
    cin >> queryType;
    if (queryType == 1) {
      cin >> name >> val;
    } else {
      cin >> name;
    }
    if (queryType == 1) {
      if (m.find(name) != m.end()) {
        m[name] = val + m[name];
      } else {
        m[name] = val;
      }
    } else if (queryType == 2) {
      m.erase(name);
    } else {
      if (m.find(name) != m.end()) {
        cout << m[name] << endl;
      } else {
        cout << 0 << endl;
      }
    }
  }

  return 0;
}
