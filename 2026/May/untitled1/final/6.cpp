#include <iostream>
#include <windows.h>
using namespace std;
template <typename T>
class ProductWeight {
private:
    T weight;
public:
    ProductWeight(T w) : weight(w) {}
    void Show() const {
        cout << "Вага: " << weight << " кг" << endl;
    }
    bool IsHeavier(const ProductWeight<T>& other) const {
        return this->weight > other.weight;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    ProductWeight<double> apple(0.25);
    ProductWeight<double> melon(3.5);
    if (melon.IsHeavier(apple)) {
        cout << "Диня важча" << endl;
    }
    return 0;
}