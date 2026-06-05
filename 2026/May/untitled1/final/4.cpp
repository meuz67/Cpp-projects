#include <iostream>
#include <windows.h>
using namespace std;
template <typename T>
class Cup {
private:
    T volume;
public:
    Cup(T vol) : volume(vol) {}
    void Show() const {
        cout << "Об'єм напою: " << volume << " мл" << endl;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Cup<int> tea(250);
    Cup<double> coffee(330.5);
    tea.Show();
    coffee.Show();
    return 0;
}