#include <iostream>
#include <windows.h>
using namespace std;
template <typename T>
class ProductPrice {
private:
    T price;
public:
    ProductPrice(T pr) : price(pr) {}
    void SetPrice(T pr) {
        price = pr;
    }
    T GetPrice() const {
        return price;
    }
    void ShowPrice() const {
        cout << "Ціна товару: " << price << " грн" << endl;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    ProductPrice<int> bread(35);
    ProductPrice<double> cheese(125.75);
    bread.ShowPrice();
    cheese.ShowPrice();
    return 0;
}