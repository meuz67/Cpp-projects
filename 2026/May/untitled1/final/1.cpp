#include <iostream>
#include <string>
#include <windows.h>
class Circle {
private:
    double radius;
public:
    Circle(double r = 0.0) : radius(r) {};
    double getRadius()const {
        return radius;
    }
    bool operator==(const Circle& c) const {
        return radius == c.radius;
    }
    bool operator>(const Circle& c) const {
        return radius > c.radius;
    }
    Circle& operator+=(double value) {
        if (this->radius + value >= 0) {
            this->radius += value;
        } else {
            this->radius = 0;
        }
        return *this;
    }
    Circle& operator-=(double value) {
        if (this->radius - value >= 0) {
            this->radius -= value;
        } else {
            this->radius = 0;
        }
        return *this;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Circle c1(5.0);
    Circle c2(7.5);
    Circle c3(5.0);
    std::cout << "Радиус круга 1: " << c1.getRadius() << "\n";
    std::cout << "Радиус круга 2: " << c2.getRadius() << "\n";
    std::cout << "Радиус круга 3: " << c3.getRadius() << "\n";
    if (c1 == c3) {
        std::cout << "Круг 1 равен кругу 3\n";
    } else {
        std::cout << "Круг 1 не равен кругу 3\n";
    }
    if (c2 > c1) {
        std::cout << "Круг 2 больше круга 1\n";
    }
    c1 += 3.5;
    std::cout << "Радиус круга 1 после += 3.5: " << c1.getRadius() << "\n";
    c1 -= 2.0;
    std::cout << "Радиус круга 1 после -= 2.0: " << c1.getRadius() << "\n";
    return 0;
}