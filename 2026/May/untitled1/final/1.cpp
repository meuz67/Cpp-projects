#include <iostream>
#include <string>
#include <windows.h>
using namespace std;
class Student {
private:
    string name;
    double average;
    int absences;
public:
    Student(string n, double a, int ab) {
        name = n;
        average = a;
        absences = ab;
    }
    bool operator==(const Student& other) {
        return average == other.average;
    }
    bool operator>(const Student& other) {
        return average > other.average;
    }
    Student& operator++() {
        average += 0.5;
        return *this;
    }
    Student& operator--() {
        absences++;
        return *this;
    }
    Student& operator+=(double points) {
        average += points;
        return *this;
    }
    Student& operator-=(double points) {
        average -= points;
        if (average < 0)
            average = 0;

        return *this;
    }
    void print() {
        cout << name
             << " | Бал: " << average
             << " | Пропуски: " << absences
             << endl;
    }
};
int main() {
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Student s1("Іван", 8.5, 2);
    Student s2("Петро", 9.0, 1);
    s1.print();
    s2.print();
    cout << (s1 == s2) << endl;
    cout << (s1 > s2) << endl;
    ++s1;
    --s1;
    s1 += 1;
    s1 -= 0.5;
    s1.print();
    return 0;
}