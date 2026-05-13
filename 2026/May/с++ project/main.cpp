#include <iostream>

class Drib {
public:
    int chyselnik;
    int znamennik;
    void setDrib(int chyselnik, int znamennik) {
        this->chyselnik = chyselnik;
        this->znamennik = znamennik;
    }
    void getDrib() {
        std::cout << "Drib chyselnik: " << chyselnik << std::endl;
        std::cout << "Znamen chyselnik: " << znamennik << std::endl;
    }
};
Drib sumOfDribs(Drib first, Drib second) {
    Drib result;
    result.chyselnik = (first.chyselnik * second.znamennik) + (first.znamennik * second.chyselnik);
    result.znamennik = first.znamennik * second.znamennik;
    return result;
}
Drib differenceOfDribs(Drib first, Drib second) {
    Drib result;
    result.chyselnik = (first.chyselnik * second.znamennik) - (second.chyselnik * first.znamennik);
    result.znamennik = first.znamennik * second.znamennik;
    return result;
}
Drib pluralOfDribs(Drib first, Drib second) {
    Drib result;
    result.chyselnik = first.chyselnik * second.chyselnik;
    result.znamennik = first.znamennik * second.znamennik;
    return result;
}
Drib divisionOfDribs(Drib first, Drib second) {
    Drib result;
    result.chyselnik = first.chyselnik * second.znamennik;
    result.znamennik = first.znamennik * second.chyselnik;
    return result;
}
int main() {
    Drib first;
    first.setDrib(1, 2);
    Drib second;
    second.setDrib(3, 4);
    Drib result;
    result = sumOfDribs(first, second);
    result.getDrib();
    result = differenceOfDribs(first, second);
    result.getDrib();
    result = pluralOfDribs(first, second);
    result.getDrib();
    result = divisionOfDribs(first, second);
    result.getDrib();
    return 0;
}