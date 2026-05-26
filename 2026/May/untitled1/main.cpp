#include <iostream>
#include <windows.h>
#include <ctime>
#include <cstdlib>
#include <string>
#include <fstream>
class Package {
private:
    std::string name;
    double weight;
    double shippingCost;
public:
    Package(std::string n, double w, double cost) : name(n), weight(w), shippingCost(cost) {}
    double getWeight() const {
        return this->weight;
    }
    std::string getName() const {
        return this->name;
    }
    friend std::ostream& operator<<(std::ostream& os, const Package& pkg);
};
bool operator>(const Package& p1, const Package& p2) {
    return p1.getWeight() > p2.getWeight();
}
std::ostream& operator<<(std::ostream& os, const Package& pkg) {
    os << "Посилка \"" << pkg.name << "\" (Вага: " << pkg.weight
       << " кг, Вартість доставки: " << pkg.shippingCost << " грн)";
    return os;
}
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Package pkg1("Ноутбук", 2.5, 120.0);
    Package pkg2("Крісло", 15.0, 450.0);
    std::cout << pkg1 << std::endl;
    std::cout << pkg2 << std::endl;
    std::cout << "\nРезультат порівняння:" << std::endl;
    if (pkg1 > pkg2) {
        std::cout << "Посилка \"" << pkg1.getName() << "\" важча за посилку \"" << pkg2.getName() << "\"." << std::endl;
    } else if (pkg2 > pkg1) {
        std::cout << "Посилка \"" << pkg2.getName() << "\" важча за посилку \"" << pkg1.getName() << "\"." << std::endl;
    } else {
        std::cout << "Вага посилок однакова." << std::endl;
    }
    return 0;
};