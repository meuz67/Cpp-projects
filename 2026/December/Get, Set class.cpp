#include <iostream>
#include <string>
using namespace std;
class Person {
    private:
    string name;
    int age;
public:
    void setName(const string& newName) {
        name = newName;
    }
    void setAge(int newAge) {
        age = newAge;
    }
    void displayInfo() {
        cout << "Name: " << name << ", Age: " << age << endl;
    }  
};
int main() {
    Person person;
	person.setAge(30);
    person.setName("Alice");
    person.displayInfo();
	return 0;
}
