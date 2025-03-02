#include <iostream>
#include <thread>
#include <chrono>
int main() {
    std::cout << "Hello, "<<std::flush;  // 手动刷新，立即输出 "Hello, "
    std::this_thread::sleep_for(std::chrono::seconds(2));  // 等待2秒
    std::cout << "world!" << std::endl;  // 输出 "world!" 并自动换行和刷新
    return 0;
}