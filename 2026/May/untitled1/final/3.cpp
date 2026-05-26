#include <iostream>
#include <windows.h>
class Fridge {
private:
    int maxProducts;
    int currentProducts;
public:
    Fridge(int maxProd, int currentProd) : maxProducts(maxProd), currentProducts(currentProd) {
        if (this->currentProducts > this->maxProducts) {
            this->currentProducts = this->maxProducts;
        }
    }
    void add(int n) {
        if (this->currentProducts + n <= this->maxProducts) {
            this->currentProducts += n;
        } else {
            this->currentProducts = this->maxProducts;
        }
    }
    void remove(int n) {
        if (this->currentProducts - n >= 0) {
            this->currentProducts -= n;
        } else {
            this->currentProducts = 0;
        }
    }
    void status() const {
        std::cout << "Холодильник: заповнено на " << this->currentProducts
                  << " з " << this->maxProducts << " продуктів." << std::endl;
    }
};
int main() {
    SetConsoleOutputCP(CP_UTF8);
    SetConsoleCP(CP_UTF8);
    Fridge myFridge(50, 20);
    myFridge.status();
    myFridge.add(15);
    myFridge.status();
    myFridge.remove(5);
    myFridge.status();
    const Fridge ecoFridge(30, 25);
    ecoFridge.status();
    return 0;
}