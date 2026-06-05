#include <iostream>
#include <windows.h>
using namespace std;
template <typename T>
class BankAccount {
private:
    T balance;
public:
    BankAccount(T initialBalance) : balance(initialBalance) {}
    void Deposit(T amount) {
        balance += amount;
    }
    void Withdraw(T amount) {
        if (balance >= amount) {
            balance -= amount;
        } else {
            cout << "Помилка: Недостатньо коштів на рахунку!" << endl;
        }
    }
    void ShowBalance() const {
        cout << "Баланс: " << balance << " грн" << endl;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    BankAccount<double> account(1000.50);
    account.Deposit(500);
    account.Withdraw(200);
    account.ShowBalance();
    return 0;
}