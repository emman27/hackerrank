#include <iostream>
#include <string>
using namespace std;

int main() {
  // Complete the program
  string a, b;
  cin >> a >> b;
  cout << a.length() << ' ' << b.length() << endl;
  cout << a << b << endl;
  string tmp = string(1, a[0]);
  string ap = string(1, b[0]) + a.substr(1, a.length() - 1);
  string bp = tmp + b.substr(1, b.length() - 1);
  cout << ap << ' ' << bp << endl;
  return 0;
}
