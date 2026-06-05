#include <iostream>
#include <string>
#include <windows.h>
template<typename T>
T average(T arr[], int size) {
    T sum = 0;
    for (int i = 0; i < size; i++) {
        sum += arr[i];
    }
    return sum / size;
}
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    int arr[] = {67, 42, 1, 15};
    std::cout << average(arr, 4) << std::endl;
    return 0;
}