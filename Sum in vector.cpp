#include <iostream>
#include <vector>
#include <string>
using namespace std;
int main() {
    int size;
    cin >> size;
    vector<int> newVector;
    for (int i = 0; i < 5; i++) {
        int userInput;
        cin >> userInput;
        newVector.push_back(userInput);
    }
    int target;
    cin >> target;
    for (int i = 0; i < newVector.size(); i++) {
        if (newVector[i] + newVector[i + 1] == target) {
            string answer = to_string(newVector[i]) + " " + to_string(newVector[i + 1]);
            cout << answer << endl;
        }
    }
}