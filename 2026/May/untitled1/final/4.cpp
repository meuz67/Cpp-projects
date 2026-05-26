#include <iostream>
#include <windows.h>
class Lamp {
private:
    bool isOn;
    int brightness;
public:
    Lamp() : isOn(false), brightness(0) {}
    Lamp(bool state, int b) : isOn(state), brightness(b) {
        if (this->brightness < 0) this->brightness = 0;
        if (this->brightness > 100) this->brightness = 100;
        if (!this->isOn) this->brightness = 0;
    }
    void turnOn() {
        this->isOn = true;
        if (this->brightness == 0) {
            this->brightness = 50;
        }
    }
    void turnOff() {
        this->isOn = false;
        this->brightness = 0;
    }
    void setBrightness(int b) {
        if (this->isOn) {
            this->brightness = b;
            if (this->brightness < 0) this->brightness = 0;
            if (this->brightness > 100) this->brightness = 100;
        }
    }
    void show() const {
        std::cout << "Лампа: " << (this->isOn ? "Ввімкнена" : "Вимкнена")
                  << ", Яскравість: " << this->brightness << "%" << std::endl;
    }
};
int main() {
    SetConsoleOutputCP(CP_UTF8);
    SetConsoleCP(CP_UTF8);
    Lamp deskLamp;
    deskLamp.show();
    deskLamp.turnOn();
    deskLamp.setBrightness(85);
    deskLamp.show();
    deskLamp.turnOff();
    deskLamp.show();
    const Lamp nightLamp(true, 20);
    nightLamp.show();
    return 0;
}