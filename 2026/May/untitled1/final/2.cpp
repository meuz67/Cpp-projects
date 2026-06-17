#include <iostream>
#include <string>
#include <Windows.h>
using namespace std;
class Ronaldo {
public:
    class Achievement {
    public:
        string title;
        int year;
        Achievement(string t = "", int y = 0) : title(t), year(y) {}
    };
    string name;
    string country;
    Achievement achievement;
    Ronaldo(string n, string c, string t, int y) : name(n), country(c), achievement(t, y) {}
};
int main() {
    SetConsoleCP(1251);
    SetConsoleOutputCP(1251);
    Ronaldo ronaldo("Кріштіану Роналду", "Португалія", "Золотий м'яч", 2008);
    return 0;
}