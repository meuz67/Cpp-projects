#include <iostream>
#include <windows.h>
using namespace std;
template <typename T>
class FuelTank {
private:
    T fuelAmount;
public:
    FuelTank(T initialFuel) : fuelAmount(initialFuel) {}
    void AddFuel(T amount) {
        fuelAmount += amount;
    }
    void ShowFuel() const {
        cout << "У баку: " << fuelAmount << " літрів" << endl;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    FuelTank<int> car(40);
    car.AddFuel(10);
    car.ShowFuel();
    return 0;
}