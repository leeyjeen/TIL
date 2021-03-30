// https://adventofcode.com/2020/day/1
#include <iostream>
#include <string>
#include <vector>
#include <fstream>

using namespace std;

vector<int> readInput(string filepath) {
    vector<int> inputList;
    ifstream inputFile;
    inputFile.open(filepath);
    
    int inputValue;
    while (inputFile >> inputValue) {
        inputList.push_back(inputValue);
    }
    inputFile.close();
    return inputList;
}

int solveFirst(const vector<int>& numbers) {
    for (int i=0; i<numbers.size(); i++) {
        for (int j=i+1; j<numbers.size(); j++) {
            if (numbers[i] + numbers[j] == 2020) {
                return numbers[i]*numbers[j];
            }
        }
    }
    return 0;
}

int solveSecond(const vector<int>& numbers) {
    for (int i=0; i<numbers.size(); i++) {
        for (int j=i+1; j<numbers.size(); j++) {
            for (int k=j+1; k<numbers.size(); k++) {
                if (numbers[i] + numbers[j] + numbers[k] == 2020) {
                    return numbers[i]*numbers[j]*numbers[k];
                }
            }
        }
    }
    return 0;
}

int main() {
    vector<int> inputList = readInput("./day1.txt");
    cout << solveFirst(inputList) << endl;
    cout << solveSecond(inputList) << endl;
}
