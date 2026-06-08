#include <iostream>
#include <string>
#include <windows.h>
template <typename T>
class Ronaldo {
private:
    std::string name;
    std::string club;
    std::string country;
    int matches;
    T goals;
    T assists;
public:
    Ronaldo(std::string n, std::string cl, std::string co, int m, T g, T a) : name(n), club(cl), country(co), matches(m), goals(g), assists(a) {}
    void ShowInfo() {
        std::cout << "Ім'я: " << name << ", Клуб: " << club << ", Країна: " << country << ", Матчі: " << matches << ", Голи: " << goals << ", Передачі: " << assists << "\n";
    }
    double ShowGoalAverage() {
        if (matches == 0) return 0.0;
        return static_cast<double>(goals) / matches;
    }
    T ShowGoalActions() {
        return goals + assists;
    }
    void IsLegend() {
        if (goals > 900) {
            std::cout << "Легенда світового футболу\n";
        } else {
            std::cout << "Продовжує писати історію\n";
        }
    }
    void CanReach1000Goals() {
        if ((1000 - goals) < 100) {
            std::cout << "Роналду близький до 1000 голів\n";
        }
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Ronaldo<int> cr7("Cristiano Ronaldo", "Al-Nassr", "Portugal", 1200, 910, 250);
    cr7.ShowInfo();
    std::cout << "Середня кількість голів за матч: " << cr7.ShowGoalAverage() << "\n";
    std::cout << "Загальна кількість дій (голи + паси): " << cr7.ShowGoalActions() << "\n";
    cr7.IsLegend();
    cr7.CanReach1000Goals();
    return 0;
}