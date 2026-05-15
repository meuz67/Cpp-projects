#include <iostream>

class Phone {
private:
    std::string brand;
    std::string model;
    int price;
    public:
    void setPrice(int price) {
        this->price = price;
    }
    void setModel(std::string model) {
        this->model = model;
    }
    void setBrand(std::string brand) {
        this->brand = brand;
    }
    std::string getBrand() {
        return brand;
    }
    std::string getModel() {
        return model;
    }
    int getPrice() {
        return price;
    }
};
class Teacher {
    private:
    std::string name;
    std::string subject;
    int yearsWorked;
    public:
    void setName(std::string name) {
        this->name = name;
    }
    void setSubject(std::string subject) {
        this->subject = subject;
    }
    void setYearsWorked(int years) {
        this->yearsWorked = years;
    }
    std::string getName() {
        return name;
    }
    std::string getSubject() {
        return subject;
    }
    int getYearsWorked() {
        return yearsWorked;
    }
};
int main() {
    return 0;
}