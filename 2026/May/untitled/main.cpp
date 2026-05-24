#include <iostream>
#include <cstring>
#include <windows.h>
class Dish {
private:
    std::string name;
    int time;
    public:
    Dish(std::string name, int time) {
        this->name = name;
        this->time = time;
    }
    void print() {
        std::cout << name << std::endl;
        std::cout << time << std::endl;
    }
};
int main() {
    SetConsoleOutputCP(CP_UTF8);
    SetConsoleCP(CP_UTF8);
    Dish* dish = new Dish("dish", 10);
    dish->print();
    delete dish;
    return 0;
}