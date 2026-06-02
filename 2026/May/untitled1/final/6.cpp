#include <iostream>
#include <string>
#include <windows.h>
class GameCharacter {
private:
    std::string name;
    int healthPoints;
    int experiencePoints;
public:
    GameCharacter(std::string charName, int hp, int xp = 0)
        : name(charName), healthPoints(hp), experiencePoints(xp) {
        if (healthPoints < 0) healthPoints = 0;
        if (experiencePoints < 0) experiencePoints = 0;
    }
    std::string getName() const { return name; }
    int getHP() const { return healthPoints; }
    int getXP() const { return experiencePoints; }
    bool operator==(const GameCharacter& other) const {
        return this->healthPoints == other.healthPoints;
    }
    bool operator>(const GameCharacter& other) const {
        return this->experiencePoints > other.experiencePoints;
    }
    GameCharacter& operator+=(int xpGained) {
        if (xpGained > 0) {
            this->experiencePoints += xpGained;
        }
        return *this;
    }
    GameCharacter& operator-=(int damage) {
        if (damage > 0) {
            this->healthPoints -= damage;
            if (this->healthPoints < 0) {
                this->healthPoints = 0;
            }
        }
        return *this;
    }
    GameCharacter& operator++() {
        this->experiencePoints += 1000;
        return *this;
    }
    GameCharacter& operator--() {
        this->healthPoints -= 10;
        if (this->healthPoints < 0) {
            this->healthPoints = 0;
        }
        return *this;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    GameCharacter hero1("Воїн", 100, 500);
    GameCharacter hero2("Маг", 80, 1200);
    GameCharacter hero3("Лучник", 100, 300);
    std::cout << "Персонаж 1: " << hero1.getName() << " | HP: " << hero1.getHP() << " | XP: " << hero1.getXP() << "\n";
    std::cout << "Персонаж 2: " << hero2.getName() << " | HP: " << hero2.getHP() << " | XP: " << hero2.getXP() << "\n";
    std::cout << "Персонаж 3: " << hero3.getName() << " | HP: " << hero3.getHP() << " | XP: " << hero3.getXP() << "\n\n";
    if (hero1 == hero3) {
        std::cout << "У персонажів \"" << hero1.getName() << "\" та \"" << hero3.getName() << "\" однакова кількість здоров'я\n";
    }
    if (hero2 > hero1) {
        std::cout << "Досвід персонажа \"" << hero2.getName() << "\" більший, ніж у \"" << hero1.getName() << "\"\n\n";
    }
    hero1 += 250;
    std::cout << "Досвід \"" << hero1.getName() << "\" після += 250 (отримання досвіду): " << hero1.getXP() << "\n";
    hero1 -= 40;
    std::cout << "Здоров'я \"" << hero1.getName() << "\" після -= 40 (отримання шкоди): " << hero1.getHP() << "\n\n";
    ++hero1;
    std::cout << "Досвід \"" << hero1.getName() << "\" після ++ (підвищення рівня): " << hero1.getXP() << "\n";
    --hero1;
    std::cout << "Здоров'я \"" << hero1.getName() << "\" після -- (зменшення на 10 HP): " << hero1.getHP() << "\n";
    hero1 -= 100;
    std::cout << "Здоров'я \"" << hero1.getName() << "\" після надмірної шкоди -= 100 (обмеження нуля): " << hero1.getHP() << "\n";
    return 0;
}