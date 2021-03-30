// https://adventofcode.com/2020/day/3
#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <cstring>

using namespace std;

const char TREE = '#';

vector<string> readInput(string filepath) {
    vector<string> inputList;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (inputFile >> line) {
        inputList.push_back(line);
    }
    inputFile.close();
    return inputList;
}

int getTreeCount(int rightInterval, int downInterval, const vector<string>& inputList) {
    int row = 0;
    int col = 0;
    int treeCount = 0;
    while (row+downInterval < inputList.size()) {
        col = (col+rightInterval) % inputList[0].size();
        row += downInterval;
        if (inputList[row][col] == TREE) {
            treeCount++;
        }
    }
    return treeCount;
}

int solveFirst(const vector<string>& inputList) {
    return getTreeCount(3, 1, inputList);
}

int solveSecond(const vector<string>& inputList) {
    return getTreeCount(1, 1, inputList)*getTreeCount(3, 1, inputList)*getTreeCount(5, 1, inputList)*getTreeCount(7, 1, inputList)*getTreeCount(1, 2, inputList);
}

int main() {
    vector<string> inputList = readInput("./day3.txt");
    cout << solveFirst(inputList) << endl;
    cout << solveSecond(inputList) << endl;
}