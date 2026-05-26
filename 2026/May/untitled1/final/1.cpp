#include <iostream>
#include <windows.h>
class Bottle {
private:
    double capacity;
    double currentVolume;
public:
    Bottle(double cap, double current) : capacity(cap), currentVolume(current) {
        if (currentVolume > capacity) {
            currentVolume = capacity;
        }
    }
    void fill(double liters) {
        if (this->currentVolume + liters <= this->capacity) {
            this->currentVolume += liters;
        } else {
            this->currentVolume = this->capacity;
        }
    }
    void getStatus() const {
        std::cout << "Пляшка: об'єм = " << this->capacity
                  << " л, поточна кількість води = " << this->currentVolume
                  << " л." << std::endl;
    }
};
int main() {
    SetConsoleOutputCP(CP_UTF8);
    SetConsoleCP(CP_UTF8);
    Bottle smallBottle(0.5, 0.2);
    Bottle largeBottle(2.0, 1.5);
    const Bottle ecoBottle(1.0, 0.7);
    smallBottle.fill(0.2);
    largeBottle.fill(1.0);
    smallBottle.getStatus();
    largeBottle.getStatus();
    ecoBottle.getStatus();
    return 0;
}