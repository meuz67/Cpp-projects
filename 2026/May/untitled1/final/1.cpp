#include <iostream>
#include <string>
#include <windows.h>
template <typename T>
class Car {
private:
    std::string brand;
    std::string color;
    int year;
    T price;
public:
    Car(std::string b, std::string c, int y, T p) : brand(b), color(c), year(y), price(p) {}
    void ShowInfo() {
        std::cout << "Марка: " << brand << ", Колір: " << color << ", Рік: " << year << ", Ціна: " << price << "\n";
    }
    void IsExpensive() {
        if (price > 1000000) {
            std::cout << "Автомобіль є дорогим (ціна більше 1 000 000 грн).\n";
        } else {
            std::cout << "Автомобіль має помірну ціну.\n";
        }
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Car<double> myCar("Audi", "Чорний", 2022, 1250000.50);
    myCar.ShowInfo();
    myCar.IsExpensive();
    return 0;
}