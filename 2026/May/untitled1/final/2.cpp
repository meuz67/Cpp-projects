#include <iostream>
#include <string>
#include <windows.h>
class Room {
private:
    std::string name;
    double temperature;
public:
    Room(std::string roomName, double initialTemp) : name(roomName), temperature(initialTemp) {}

    void heat(double t) {
        this->temperature += t;
    }
    void cool(double t) {
        this->temperature -= t;
    }
    void show() const {
        std::cout << "Кімната: " << this->name
                  << ", поточна температура: " << this->temperature
                  << "°C" << std::endl;
    }
};
int main() {
    SetConsoleOutputCP(CP_UTF8);
    SetConsoleCP(CP_UTF8);
    Room livingRoom("Вітальня", 20.5);
    Room bedroom("Спальня", 18.0);
    const Room kitchen("Кухня", 22.0);
    livingRoom.heat(2.5);
    bedroom.cool(1.5);
    livingRoom.show();
    bedroom.show();
    kitchen.show();
    return 0;
}