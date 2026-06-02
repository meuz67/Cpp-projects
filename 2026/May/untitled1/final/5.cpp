#include <iostream>
#include <string>
#include <windows.h>
class WarehouseItem {
private:
    std::string name;
    int quantity;
    double pricePerUnit;
public:
    WarehouseItem(std::string itemName, int q, double price)
        : name(itemName), quantity(q), pricePerUnit(price) {
        if (quantity < 0) quantity = 0;
        if (pricePerUnit < 0.0) pricePerUnit = 0.0;
    }
    std::string getName() const { return name; }
    int getQuantity() const { return quantity; }
    double getPricePerUnit() const { return pricePerUnit; }
    double getTotalValue() const { return quantity * pricePerUnit; }
    bool operator==(const WarehouseItem& other) const {
        return this->pricePerUnit == other.pricePerUnit;
    }
    bool operator>(const WarehouseItem& other) const {
        return this->getTotalValue() > other.getTotalValue();
    }
    WarehouseItem& operator+=(int count) {
        if (count > 0) {
            this->quantity += count;
        }
        return *this;
    }
    WarehouseItem& operator-=(int count) {
        if (count > 0) {
            if (this->quantity >= count) {
                this->quantity -= count;
            } else {
                this->quantity = 0;
            }
        }
        return *this;
    }
    WarehouseItem& operator++() {
        this->quantity += 1;
        return *this;
    }
    WarehouseItem& operator--() {
        if (this->quantity > 0) {
            this->quantity -= 1;
        }
        return *this;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    WarehouseItem item1("Монітор", 10, 4500.0);
    WarehouseItem item2("Клавіатура", 50, 1200.0);
    WarehouseItem item3("Мишка", 30, 1200.0);
    std::cout << "Товар 1: " << item1.getName() << " | К-сть: " << item1.getQuantity() << " | Ціна: " << item1.getPricePerUnit() << " грн | Загальна вартість: " << item1.getTotalValue() << " грн\n";
    std::cout << "Товар 2: " << item2.getName() << " | К-сть: " << item2.getQuantity() << " | Ціна: " << item2.getPricePerUnit() << " грн | Загальна вартість: " << item2.getTotalValue() << " грн\n";
    std::cout << "Товар 3: " << item3.getName() << " | К-сть: " << item3.getQuantity() << " | Ціна: " << item3.getPricePerUnit() << " грн | Загальна вартість: " << item3.getTotalValue() << " грн\n\n";
    if (item2 == item3) {
        std::cout << "Ціна за одиницю товару \"" << item2.getName() << "\" та \"" << item3.getName() << "\" однакова\n";
    }
    if (item2 > item1) {
        std::cout << "Загальна вартість запасів \"" << item2.getName() << "\" більша, ніж \"" << item1.getName() << "\"\n\n";
    }
    item1 += 5;
    std::cout << "Кількість \"" << item1.getName() << "\" після += 5 (надходження): " << item1.getQuantity() << "\n";
    item1 -= 3;
    std::cout << "Кількість \"" << item1.getName() << "\" після -= 3 (відвантаження): " << item1.getQuantity() << "\n\n";
    ++item1;
    std::cout << "Кількість \"" << item1.getName() << "\" після ++ (автододавання 1 шт): " << item1.getQuantity() << "\n";
    --item1;
    std::cout << "Кількість \"" << item1.getName() << "\" після -- (списання пошкодженої 1 шт): " << item1.getQuantity() << "\n";
    return 0;
}