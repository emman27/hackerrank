#include <bits/stdc++.h>

using namespace std;

class Box {
 private:
  long long l, b, h;

 public:
  Box() : l(0), b(0), h(0) {}

  Box(int l, int b, int h) : l(l), b(b), h(h) {}

  Box(const Box& box)
      : l(box.getLength()), b(box.getBreadth()), h(box.getHeight()) {}

  int getLength() const { return l; }

  int getBreadth() const { return b; }

  int getHeight() const { return h; }

  long long CalculateVolume() { return l * b * h; }

  bool operator<(Box& box) {
    return l < box.getLength() ||
           (l == box.getLength() && b < box.getBreadth()) ||
           (l == box.getLength() && b == box.getBreadth() &&
            h < box.getHeight());
  }

  friend ostream& operator<<(ostream& out, const Box& box) {
    out << box.getLength() << ' ' << box.getBreadth() << ' ' << box.getHeight();
    return out;
  }
};
// Implement the class Box
// l,b,h are integers representing the dimensions of the box

// The class should have the following functions :

// Constructors:
// Box();
// Box(int,int,int);
// Box(Box);

// int getLength(); // Return box's length
// int getBreadth (); // Return box's breadth
// int getHeight ();  //Return box's height
// long long CalculateVolume(); // Return the volume of the box

// Overload operator < as specified
// bool operator<(Box& b)

// Overload operator << as specified
// ostream& operator<<(ostream& out, Box& B)
