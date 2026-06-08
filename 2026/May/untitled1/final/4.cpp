#include <iostream>
#include <string>
#include <windows.h>
template <typename T>
class OnlineShop {
private:
    std::string productName;
    std::string category;
    int stockCount;
    T pricePerUnit;
    int soldCount;
public:
    OnlineShop(std::string name, std::string cat, int stock, T price, int sold) : productName(name), category(cat), stockCount(stock), pricePerUnit(price), soldCount(sold) {}
    void ShowInfo() {
        std::cout << "Товар: " << productName << ", Категорія: " << category << ", На складі: " << stockCount << ", Ціна: " << pricePerUnit << ", Продано: " << soldCount << "\n";
    }
    T ShowIncome() {
        return pricePerUnit * soldCount;
    }
    void NeedRestock() {
        if (stockCount < 10) {
            std::cout << "Увага! На складі менше 10 одиниць товару. Потрібно поповнити запаси!\n";
        } else {
            std::cout << "Товару на складі достатньо.\n";
        }
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    OnlineShop<double> shop("Ноутбук", "Електроніка", 5, 24500.0, 15);
    shop.ShowInfo();
    std::cout << "Дохід від продажів: " << shop.ShowIncome() << " грн.\n";
    shop.NeedRestock();
    return 0;
}