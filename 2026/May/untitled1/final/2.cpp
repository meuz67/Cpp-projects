#include <iostream>
#include <string>
#include <windows.h>
template <typename T>
class Smartphone {
private:
    std::string model;
    std::string manufacturer;
    int memoryGB;
    T price;
    double rating;
public:
    Smartphone(std::string m, std::string mf, int mem, T p, double r) : model(m), manufacturer(mf), memoryGB(mem), price(p), rating(r) {}
    void ShowInfo() {
        std::cout << "Виробник: " << manufacturer << ", Модель: " << model << ", Пам'ять: " << memoryGB << " GB, Ціна: " << price << ", Рейтинг: " << rating << "\n";
    }
    void IsRecommended() {
        if (rating > 4.5) {
            std::cout << "Смартфон рекомендовано до покупки!\n";
        } else {
            std::cout << "Смартфон має звичайний рейтинг.\n";
        }
    }
    double PricePerGB() {
        return static_cast<double>(price) / memoryGB;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Smartphone<int> phone("Galaxy S24", "Samsung", 256, 38000, 4.8);
    phone.ShowInfo();
    phone.IsRecommended();
    std::cout << "Ціна за 1 GB: " << phone.PricePerGB() << " грн.\n";
    return 0;
}