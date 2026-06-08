#include <iostream>
#include <string>
#include <windows.h>
template <typename T>
class TravelPackage {
private:
    std::string country;
    std::string hotel;
    int days;
    T packageCost;
    int touristsCount;
public:
    TravelPackage(std::string c, std::string h, int d, T cost, int tourists) : country(c), hotel(h), days(d), packageCost(cost), touristsCount(tourists) {}
    void ShowInfo() {
        std::cout << "Країна: " << country << ", Готель: " << hotel << ", Кількість днів: " << days << ", Вартість путівки: " << packageCost << ", Кількість туристів: " << touristsCount << "\n";
    }
    T ShowTotalCost() {
        return packageCost * touristsCount;
    }
    void IsVIP() {
        if (packageCost > 50000) {
            std::cout << "Це VIP путівка!\n";
        } else {
            std::cout << "Стандартна путівка.\n";
        }
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    TravelPackage<double> pack("Єгипет", "Rixos", 10, 65000.0, 2);
    pack.ShowInfo();
    std::cout << "Загальна вартість для всіх туристів: " << pack.ShowTotalCost() << " грн.\n";
    pack.IsVIP();
    return 0;
}