#include <iostream>
#include <string>
#include <fstream>
#include <windows.h>
#include <vector>
using namespace std;
class CPU {
private:
    string model;
public:
    CPU(string cpuModel) {
        model = cpuModel;
        cout << "  [Конструктор CPU]: " << model << " создан." << endl;
    }
    ~CPU() {
        cout << "  [Деструктор CPU]: " << model << " уничтожен!" << endl;
    }
    void showInfo() {
        cout << "Процессор: " << model << endl;
    }
};
class RAM {
private:
    int sizeGb;
public:
    RAM(int size) {
        sizeGb = size;
        cout << "  [Конструктор RAM]: " << sizeGb << " ГБ создано." << endl;
    }
    ~RAM() {
        cout << "  [Деструктор RAM]: " << sizeGb << " ГБ уничтожено!" << endl;
    }
    void showInfo() {
        cout << "Оперативная память: " << sizeGb << " ГБ" << endl;
    }
};
class User {
private:
    string name;
public:
    User(string userName) {
        name = userName;
        cout << "[Конструктор User]: Пользователь " << name << " создан." << endl;
    }
    ~User() {
        cout << "[Деструктор User]: Пользователь " << name << " ушел (уничтожен)." << endl;
    }
    string getName() {
        return name;
    }
};
class Computer {
private:
    std::string brand;
    CPU cpu;
    RAM ram;
    User* currentUser;
public:
    Computer(string pcBrand, string cpuModel, int ramSize)
        : cpu(cpuModel), ram(ramSize) {
        brand = pcBrand;
        currentUser = nullptr;
        cout << " [Конструктор Computer]: ПК " << brand << " собран." << endl;
    }
    ~Computer() {
        cout << " [Деструктор Computer]: ПК " << brand << " утилизируется..." << endl;
    }
    void loginUser(User* user) {
        currentUser = user;
    }
    void showSystemInfo() {
        cout << "\n--- Системная информация ПК " << brand << " ---" << endl;
        cpu.showInfo();
        ram.showInfo();
        if (currentUser != nullptr) {
            cout << "Текущий пользователь: " << currentUser->getName() << endl;
        } else {
            cout << "Нет активного пользователя." << endl;
        }
        cout << "---------------------------------------\n" << endl;
    }
};
int main()
{
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);

    return 0;
}