#include <iostream>
#include <windows.h>
#include <ctime>
#include <cstdlib>
#include <string>
#include <fstream>
class Wallet {
private:
    std::string currency;
    double balance;
public:
    Wallet(std::string currency, double balance) {
        this->currency = currency;
        this->balance = balance;
    }
    friend std::ostream& operator<<(std::ostream& out, const Wallet& w);
};
std::ostream& operator<<(std::ostream& out, const Wallet& w) {
    out << w.currency << " " << w.balance;
    return out;
}
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Wallet myWallet("грн", 1550.50);
    std::cout << myWallet << std::endl;
    Wallet dollarWallet("USD", 250);
    std::cout << dollarWallet << std::endl;
    return 0;
}