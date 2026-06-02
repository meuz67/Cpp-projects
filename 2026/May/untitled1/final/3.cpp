#include <iostream>
#include <string>
#include <windows.h>
class BankAccount {
private:
    std::string accountNumber;
    double balance;
    std::string ownerName;
public:
    BankAccount(std::string number, std::string name, double initialBalance = 0.0)
        : accountNumber(number), ownerName(name), balance(initialBalance) {}
    std::string getAccountNumber() const { return accountNumber; }
    std::string getOwnerName() const { return ownerName; }
    double getBalance() const { return balance; }
    bool operator==(const BankAccount& other) const {
        return this->balance == other.balance;
    }
    bool operator>(const BankAccount& other) const {
        return this->balance > other.balance;
    }
    BankAccount& operator+=(double amount) {
        if (amount > 0) {
            this->balance += amount;
        }
        return *this;
    }
    BankAccount& operator-=(double amount) {
        if (amount > 0 && this->balance >= amount) {
            this->balance -= amount;
        }
        return *this;
    }
    BankAccount& operator++() {
        this->balance += 100.0;
        return *this;
    }
    BankAccount& operator--() {
        this->balance -= 50.0;
        return *this;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    BankAccount acc1("UA12345", "Олександр", 1000.0);
    BankAccount acc2("UA67890", "Максим", 1000.0);
    BankAccount acc3("UA55555", "Дмитро", 2500.0);
    std::cout << "Рахунок 1 (" << acc1.getOwnerName() << "): " << acc1.getBalance() << " грн\n";
    std::cout << "Рахунок 2 (" << acc2.getOwnerName() << "): " << acc2.getBalance() << " грн\n";
    std::cout << "Рахунок 3 (" << acc3.getOwnerName() << "): " << acc3.getBalance() << " грн\n\n";
    if (acc1 == acc2) {
        std::cout << "Баланси рахунків 1 та 2 рівні\n";
    }
    if (acc3 > acc1) {
        std::cout << "На рахунку 3 більше коштів, ніж на рахунку 1\n\n";
    }
    acc1 += 500.0;
    std::cout << "Баланс рахунку 1 після += 500: " << acc1.getBalance() << " грн\n";
    acc1 -= 300.0;
    std::cout << "Баланс рахунку 1 після -= 300: " << acc1.getBalance() << " грн\n\n";
    ++acc1;
    std::cout << "Баланс рахунку 1 після ++ (бонус 100 грн): " << acc1.getBalance() << " грн\n";
    --acc1;
    std::cout << "Баланс рахунку 1 після -- (комісія 50 грн): " << acc1.getBalance() << " грн\n";
    return 0;
}