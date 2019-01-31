#include <gtest/gtest.h>
#include <lib.cpp>

TEST(SumTest, Basic) {
  ASSERT_EQ(6, sum(2, 2, 2));
  ASSERT_EQ(7, sum(2, 2, 3));
}

int main(int argc, char **argv) {
  testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
