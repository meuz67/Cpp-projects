#include <iostream>
#include <windows.h>
using namespace std;
template <typename T>
class ShoppingBag {
private:
    T totalWeight;
public:
    ShoppingBag(T initialWeight) : totalWeight(initialWeight) {}
    void AddItem(T itemWeight) {
        totalWeight += itemWeight;
    }
    void ShowWeight() const {
        cout << "Поточна вага покупок: " << totalWeight << " кг" << endl;
    }
    bool IsOverloaded() const {
        return totalWeight > 10.0;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    ShoppingBag<double> bag(8.5);
    bag.AddItem(2.3);
    bag.ShowWeight();
    if (bag.IsOverloaded()) {
        cout << "Пакет занадто важкий" << endl;
    }
    return 0;
}