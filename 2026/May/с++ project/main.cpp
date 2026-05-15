#include <iostream>
#include <fstream>
class Point {
   private:
    int x;
    int y;
    int z;
    public:
    void setX(int x) {
        this->x = x;
    }
    void setY(int y) {
        this->y = y;
    }
    void setZ(int z) {
        this->z = z;
    }
    int getX() {
        return x;
    }
    int getY() {
        return y;
    }
    int getZ() {
        return z;
    }
    void input() {
        std::cout << "Enter coordinates" << std::endl;
        std::cout << "x:";
        std::cin >> x;
        std::cout << "y:";
        std::cin >> y;
        std::cout << "z:";
        std::cin >> z;
    }
    void output() {
        std::cout << "x:";
        std::cout << x << std::endl;
        std::cout << "y:";
        std::cout << y << std::endl;
        std::cout << "z:";
        std::cout << z << std::endl;
    }
    void saveToFile(std::string fileName) {
        std::ofstream file(fileName);
        if (file.is_open()) {
            file << this->getX() << " " << this->getY() << " " << this->getZ();
        }else {
            std::cout << "Unable to open file";
        }
        file.close();
    }
    void readFromFile(std::string fileName) {
        std::ifstream file(fileName);
        if (file.is_open()) {
            file >> this->x >> this->y >> this->z;
        }else {
            std::cout << "Unable to open file";
        }
        file.close();
    }
};
int main() {
    Point p1;
    p1.input();
    p1.output();
    p1.saveToFile("points.txt");
    Point p2;
    p2.readFromFile("points.txt");
    p2.output();
    return 0;
}