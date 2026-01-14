#include <iostream>
#include <string>
using namespace std;
class CoffeeMachine {
private:
    bool Check() {
        if (power > 1000) {
            cout << "High Power Coffee Machine" << endl;
            return false;
        }
        else {
            cout << "Standard Coffee Machine" << endl;
            return true;
        }
    }
public:
    int power;

    void Start() {
        if (Check()) {
            cout << "Coffee Machine Starting..." << endl;
        }
        else {
            cout << "Cannot start Coffee Machine due to high power." << endl;
        }
    }
};
int main() {
    CoffeeMachine a;
    cin >> a.power;
    a.Start();
    return 0;
}
