// https://adventofcode.com/2020/day/2
#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <cstring>

using namespace std;

struct Policy
{
    int low, high;
    char letter;
    string password;

    bool hasValidLetterCount()
    {
        int letterCount = count(password.c_str(), password.c_str() + strlen(password.c_str()), letter);
        if (low <= letterCount && letterCount <= high)
        {
            return true;
        }
        return false;
    }

    bool hasOneRightLetterPosition()
    {
        bool lowConains = (password[low - 1] == letter);
        bool highContains = (password[high - 1] == letter);
        return lowConains != highContains;
    }
};

vector<Policy> readInput(string filepath)
{
    vector<Policy> inputList;
    ifstream inputFile;
    inputFile.open(filepath);

    string ranges, letter, password;
    while (inputFile >> ranges >> letter >> password)
    {
        Policy policy;
        sscanf(ranges.c_str(), "%d-%d", &policy.low, &policy.high);
        sscanf(letter.c_str(), "%c:", &policy.letter);
        policy.password = password;
        inputList.push_back(policy);
    }
    inputFile.close();
    return inputList;
}

int solveFirst(vector<Policy> &inputList)
{
    int validCount = 0;
    for (int i = 0; i < inputList.size(); i++)
    {
        if (inputList[i].hasValidLetterCount())
        {
            validCount++;
        }
    }
    return validCount;
}

int solveSecond(vector<Policy> &inputList)
{
    int validCount = 0;
    for (int i = 0; i < inputList.size(); i++)
    {
        if (inputList[i].hasOneRightLetterPosition())
        {
            validCount++;
        }
    }
    return validCount;
}

int main()
{
    vector<Policy> inputList = readInput("./day2.txt");
    cout << solveFirst(inputList) << endl;
    cout << solveSecond(inputList) << endl;
}