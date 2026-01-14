#include <iostream>
using namespace std;

int main() {
    int n, reversed = 0;
    cin >> n;
	int original = n;
    while (n != 0) {
        int digit = n % 10;        
        reversed = reversed * 10 + digit; 
        n /= 10;                   
    }
    cout << reversed << endl;
    if (reversed == original){
        cout << "The number is a palindrome." << endl;
    } else {
        cout << "The number is not a palindrome." << endl;
	} 
    return 0;
}
