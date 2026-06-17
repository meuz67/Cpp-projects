#include <iostream>
#include <string>
#include <Windows.h>
using namespace std;
class FootballPlayer {
public:
    class Stats {
    public:
        int goals;
        int matches;
        Stats(int g = 0, int m = 0) : goals(g), matches(m) {}
    };
    string name;
    string club;
    Stats stats;
    FootballPlayer(string playerName, string playerClub, int goals, int matches) : name(playerName), club(playerClub), stats(goals, matches) {}
};
int main() {
    SetConsoleCP(1251);
    SetConsoleOutputCP(1251);
    FootballPlayer player("Андрій Шевченко", "Динамо Київ", 124, 249);
    return 0;
}