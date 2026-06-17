#include <iostream>
#include <string>
#include <Windows.h>
using namespace std;
class Club {
public:
    class Player {
    public:
        string name;
        int number;
        string position;
        Player(string n = "", int num = 0, string p = "") : name(n), number(num), position(p) {}
        void display() {
            cout << name << " " << number << " " << position << endl;
        }
    };
    string clubName;
    string country;
    Player player;
    Club(string cn, string c, string pn, int num, string p) : clubName(cn), country(c), player(pn, num, p) {}
};
int main() {
    SetConsoleCP(1251);
    SetConsoleOutputCP(1251);
    Club club("Реал Мадрид", "Іспанія", "Лука Модрич", 10, "Півзахисник");
    club.player.display();
    return 0;
}