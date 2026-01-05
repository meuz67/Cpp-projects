#include <iostream>
#include <vector>
using namespace std;

class IBicycle {
public:
    virtual void TwistTheWheel() = 0; 
    virtual void Ride() = 0;
    virtual ~IBicycle() {}  
};
class SimpleBicycle : public IBicycle {
public:
    void TwistTheWheel() override {
        cout << "Метод TwistTheWheel у SimpleBicycle" << endl; 
    }
    void Ride() override {
        cout << "Метод Ride у SimpleBicycle" << endl;  
    }
};
class SportBicycle : public IBicycle {
public:  
    void TwistTheWheel() override {  
        cout << "Метод TwistTheWheel у SportBicycle" << endl;
    }
    void Ride() override {  
        cout << "Метод Ride у SportBicycle" << endl;
    }
};

class Human {
public:
    void RideOn(IBicycle& bicycle) {
        cout << "Крутим руль" << endl;
        bicycle.TwistTheWheel(); 
        cout << "Поехали" << endl;
        bicycle.Ride();
    }
};
int main() {
    setlocale(LC_ALL, "ru");
    SimpleBicycle simpleb;
    SportBicycle sportb;
    Human h;

    h.RideOn(simpleb);
    cout << "-------------------" << endl;
    h.RideOn(sportb);

    return 0;
}