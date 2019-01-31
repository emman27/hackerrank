cmake_minimum_required(VERSION 2.6)

find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIRS})
add_library(lib lib.cpp)
include_directories("${PROJECT_SOURCE_DIR}")

add_executable(runTests main_test.cpp)
target_link_libraries(runTests ${GTEST_LIBRARIES} pthread lib)
