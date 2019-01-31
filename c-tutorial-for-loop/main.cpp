#include <cstdio>
#include <iostream>
#include <map>
using namespace std;

int main() {
  int a, b;
  scanf("%d\n%d", &a, &b);
  map<int, string> digits = {{1, "one"},   {2, "two"},   {3, "three"},
                             {4, "four"},  {5, "five"},  {6, "six"},
                             {7, "seven"}, {8, "eight"}, {9, "nine"}};
  for (int i = a; i <= b; i++) {
    if (i < 10) {
      cout << digits.at(i) << endl;
    } else if (i % 2 == 0) {
      cout << "even" << endl;
    } else {
      cout << "odd" << endl;
    }
  }
  return 0;
}
