#include <iostream>
#include <string>
#include <windows.h>

class Airplane {
private:
    std::string type;
    int currentPassengers;
    int maxPassengers;

public:
    Airplane(std::string t, int max_p, int current_p = 0)
        : type(t), maxPassengers(max_p), currentPassengers(current_p) {
        if (currentPassengers > maxPassengers) {
            currentPassengers = maxPassengers;
        }
    }
    std::string getType() const { return type; }
    int getCurrentPassengers() const { return currentPassengers; }
    int getMaxPassengers() const { return maxPassengers; }
    bool operator==(const Airplane& other) const {
        return this->type == other.type;
    }
    Airplane& operator++() {
        if (this->currentPassengers < this->maxPassengers) {
            this->currentPassengers++;
        }
        return *this;
    }
    Airplane& operator--() {
        if (this->currentPassengers > 0) {
            this->currentPassengers--;
        }
        return *this;
    }
    bool operator>(const Airplane& other) const {
        return this->maxPassengers > other.maxPassengers;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Airplane plane1("Boeing 737", 180, 178);
    Airplane plane2("Airbus A320", 150, 100);
    Airplane plane3("Boeing 737", 200, 50);
    if (plane1 == plane3) {
        std::cout << "Самолет 1 и Самолет 3 одного типа (" << plane1.getType() << ")\n";
    }
    if (plane3 > plane2) {
        std::cout << plane3.getType() << " больше по вместимости, чем " << plane2.getType() << "\n";
    }
    std::cout << "Пассажиры самолета 1: " << plane1.getCurrentPassengers() << "\n";
    ++plane1;
    std::cout << "После ++plane1: " << plane1.getCurrentPassengers() << "\n";
    ++plane1;
    std::cout << "После еще одного ++plane1: " << plane1.getCurrentPassengers() << "\n";
    --plane1;
    std::cout << "После --plane1: " << plane1.getCurrentPassengers() << "\n";
    return 0;
}