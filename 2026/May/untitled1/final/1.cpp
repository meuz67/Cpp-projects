#include <iostream>
#include <string>
#include <fstream>
#include <windows.h>
#include <vector>
using namespace std;
class Book {
private:
    std::string title;
    public:
    Book(std::string title) {
        this->title = title;
    }
    std::string getTitle() {
        return this->title;
    }
};
class Shelf {
private:
    int shelfNumber;
    std::vector<Book*> books;
    public:
    Shelf(int shelfNumber) {
        shelfNumber = shelfNumber;
    }
    void addBook(Book* book) {
        books.push_back(book);
    }
    std::vector<Book*> getBooks() {
        return this->books;
    }
    int getShelfNumber() {
        return this->shelfNumber;
    }
};
class Library {
private:
    std::string name;
    std::vector<Shelf*> shelves;
    public:
    Library(std::string name) {
        this->name = name;
    }
    void addShelf(Shelf* shelf) {
        this->shelves.push_back(shelf);
    }
    void showAllBooks() {
        std::cout << "Library number: " << name << std::endl;
        if (shelves.empty()) {
            std::cout << "No shelves" << std::endl;
            return;
        }
        for (Shelf* shelf : shelves) {
            std::cout << "Shelf number: " << shelf->getShelfNumber() << std::endl;
            vector <Book*> books = shelf->getBooks();
            if (books.empty()) {
                std::cout << "No books" << std::endl;
            }
            else {
                for (Book* book : books) {
                    std::cout << book->getTitle() << std::endl;
                }
            }
        }
        std::cout << std::endl;
    }
};
int main()
{
    SetConsoleOutputCP(1251);
    SetConsoleCP(1251);
    Book book1("Bem be");
    Book book2("Ronaldo");
    Book book3("Bim bim bam bam");
    Shelf shelf1(1);
    Shelf shelf2(2);
    shelf1.addBook(&book1);
    shelf1.addBook(&book2);
    shelf1.addBook(&book3);
    Library myLibrary("EEEEEEEEEEEEEE");
    myLibrary.addShelf(&shelf1);
    myLibrary.addShelf(&shelf2);
    myLibrary.showAllBooks();
    return 0;
}