#include <iostream>
#include <string>
#include <windows.h>
using namespace std;
class WarehouseItem {
private:
    string name;
    int quantity;
    double price;
public:
    WarehouseItem(string n, int q, double p) {
        name = n;
        quantity = q;
        price = p;
    }
    bool operator==(const WarehouseItem& other) {
        return price == other.price;
    }
    bool operator>(const WarehouseItem& other) {
        return quantity > other.quantity;
    }
    WarehouseItem& operator++() {
        quantity++;
        return *this;
    }
    WarehouseItem& operator--() {
        if (quantity > 0)
            quantity--;
        return *this;
    }
    WarehouseItem& operator+=(int amount) {
        quantity += amount;
        return *this;
    }
    WarehouseItem& operator-=(int amount) {
        quantity -= amount;
        if (quantity < 0)
            quantity = 0;
        return *this;
    }
    void print() {
        cout << name
             << " | Кількість: " << quantity
             << " | Ціна: " << price
             << endl;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    WarehouseItem w1("Ноутбук",10,25000);
    WarehouseItem w2("Миша",20,800);
    w1.print();
    w2.print();
    cout << (w1 == w2) << endl;
    cout << (w1 > w2) << endl;
    ++w1;
    --w1;
    w1 += 15;
    w1 -= 50;
    w1.print();
    return 0;
}