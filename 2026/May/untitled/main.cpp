\#include <iostream>
#include <string>
#include <windows.h>
using namespace std;

class Book {
private:
    string title;
    int pages;
public:
    void setData(string t, int p) {
        title = t;
        pages = p;
    }
    void showInfo() {
        cout << "Назва книги: " << title << endl;
        cout << "Кількість сторінок: " << pages << endl;
    }
};
int main() {
    SetConsoleOutputCP(CP_UTF8);
    SetConsoleCP(CP_UTF8);
    Book* book = new Book;
    string title;
    int pages;
    cout << "Введіть назву книги: ";
    getline(cin, title);
    cout << "Введіть кількість сторінок: ";
    cin >> pages;
    book->setData(title, pages);
    cout << endl;
    book->showInfo();
    delete book;
    return 0;
}