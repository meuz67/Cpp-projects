#include <iostream>
#include <string>
#include <Windows.h>
using namespace std;
class MusicBand {
public:
    class Member {
    public:
        string name;
        string instrument;
        int experience;
        Member(string n = "", string i = "", int e = 0) : name(n), instrument(i), experience(e) {}
        void display() {
            cout << name << " " << instrument << " " << experience << endl;
        }
    };
    string bandName;
    string style;
    Member member;
    MusicBand(string bn, string s, string mn, string i, int e) : bandName(bn), style(s), member(mn, i, e) {}
};
int main() {
    SetConsoleCP(1251);
    SetConsoleOutputCP(1251);
    MusicBand band("The Beatles", "Рок", "Джон Леннон", "Гітара", 10);
    band.member.display();
    return 0;
}