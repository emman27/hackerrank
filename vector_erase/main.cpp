#include <algorithm>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <vector>
using namespace std;

int main() {
  /* Enter your code here. Read input from STDIN. Print output to STDOUT */
  int vector_size;
  vector<int> v;
  cin >> vector_size;
  for (auto i = 0; i < vector_size; i++) {
    int n;
    cin >> n;
    v.push_back(n);
  }
  int erased;
  cin >> erased;
  v.erase(v.begin() + erased - 1);
  int start, end;
  cin >> start >> end;
  v.erase(v.begin() + start - 1, v.begin() + end - 1);
  cout << v.size() << endl;
  for (auto i = 0; i < v.size() - 1; i++) {
    cout << v.at(i) << ' ';
  }
  cout << v.at(v.size() - 1) << endl;
  return 0;
}
