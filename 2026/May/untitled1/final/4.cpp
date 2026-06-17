#include <iostream>
#include <string>
#include <Windows.h>
using namespace std;
class Singer {
public:
    class Song {
    public:
        string title;
        string genre;
        double duration;
        Song(string t = "", string g = "", double d = 0.0) : title(t), genre(g), duration(d) {}
        void display() {
            cout << title << " " << genre << " " << duration << endl;
        }
    };
    string name;
    string country;
    Song song;
    Singer(string n, string c, string t, string g, double d) : name(n), country(c), song(t, g, d) {}
};
int main() {
    SetConsoleCP(1251);
    SetConsoleOutputCP(1251);
    Singer singer("Океан Ельзи", "Україна", "Обійми", "Рок", 3.45);
    singer.song.display();
    return 0;
}