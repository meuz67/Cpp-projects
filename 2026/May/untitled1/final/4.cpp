#include <iostream>
#include <string>
#include <windows.h>
class Student {
private:
    std::string fullName;
    double averageGrade;
    int absences;
public:
    Student(std::string name, double grade, int abs = 0)
        : fullName(name), averageGrade(grade), absences(abs) {
        if (averageGrade > 12.0) averageGrade = 12.0;
        if (averageGrade < 0.0) averageGrade = 0.0;
    }
    std::string getFullName() const { return fullName; }
    double getAverageGrade() const { return averageGrade; }
    int getAbsences() const { return absences; }
    bool operator==(const Student& other) const {
        return this->averageGrade == other.averageGrade;
    }
    bool operator>(const Student& other) const {
        return this->averageGrade > other.averageGrade;
    }
    Student& operator++() {
        this->averageGrade += 0.5;
        if (this->averageGrade > 12.0) {
            this->averageGrade = 12.0;
        }
        return *this;
    }
    Student& operator--() {
        this->absences += 1;
        return *this;
    }
    Student& operator+=(double points) {
        if (points > 0) {
            this->averageGrade += points;
            if (this->averageGrade > 12.0) {
                this->averageGrade = 12.0;
            }
        }
        return *this;
    }
    Student& operator-=(double points) {
        if (points > 0) {
            this->averageGrade -= points;
            if (this->averageGrade < 0.0) {
                this->averageGrade = 0.0;
            }
        }
        return *this;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Student s1("Задко Олександр", 10.5, 2);
    Student s2("Скрипкін Максим", 10.5, 0);
    Student s3("Петров Дмитро", 8.0, 5);
    std::cout << "Студент 1: " << s1.getFullName() << ", сер. бал: " << s1.getAverageGrade() << ", пропуски: " << s1.getAbsences() << "\n";
    std::cout << "Студент 2: " << s2.getFullName() << ", сер. бал: " << s2.getAverageGrade() << ", пропуски: " << s2.getAbsences() << "\n";
    std::cout << "Студент 3: " << s3.getFullName() << ", сер. бал: " << s3.getAverageGrade() << ", пропуски: " << s3.getAbsences() << "\n\n";
    if (s1 == s2) {
        std::cout << "У " << s1.getFullName() << " та " << s2.getFullName() << " однакові середні бали\n";
    }
    if (s1 > s3) {
        std::cout << "Успішність " << s1.getFullName() << " вища, ніж у " << s3.getFullName() << "\n\n";
    }
    ++s3;
    std::cout << "Бал " << s3.getFullName() << " після ++ (підвищення на 0.5): " << s3.getAverageGrade() << "\n";
    --s3;
    std::cout << "Пропуски " << s3.getFullName() << " після -- (додано 1 пропуск): " << s3.getAbsences() << "\n\n";
    s1 += 1.2;
    std::cout << "Бал " << s1.getFullName() << " після += 1.2: " << s1.getAverageGrade() << "\n";
    s1 -= 0.7;
    std::cout << "Бал " << s1.getFullName() << " після -= 0.7: " << s1.getAverageGrade() << "\n";
    return 0;
}