// https://adventofcode.com/2020/day/5
#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <algorithm>

using namespace std;

vector<string> readInput(string filepath)
{
    vector<string> inputList;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (inputFile >> line)
    {
        inputList.push_back(line);
    }
    inputFile.close();
    return inputList;
}

int computeRow(string &rowStr)
{
    int start = 0;
    int end = 127;
    for (char &s : rowStr)
    {
        int diff = (end - start) / 2 + 1;
        if (s == 'F')
        {
            end -= diff;
        }
        else if (s == 'B')
        {
            start += diff;
        }
    }
    return start;
}

int computeCol(string &colStr)
{
    int start = 0;
    int end = 7;
    for (char &s : colStr)
    {
        int diff = (end - start) / 2 + 1;
        if (s == 'L')
        {
            end -= diff;
        }
        else if (s == 'R')
        {
            start += diff;
        }
    }
    return start;
}

int getUniqueSeatID(int row, int col)
{
    return row * 8 + col;
}

int getHighestSeatID(const vector<string> &inputList)
{
    int highestSeatID = 0;
    for (int i = 0; i < inputList.size(); i++)
    {
        string rowStr = inputList[i].substr(0, 7);
        string colStr = inputList[i].substr(7, 10);
        int row = computeRow(rowStr);
        int col = computeCol(colStr);
        int uniqueSeatID = getUniqueSeatID(row, col);
        if (highestSeatID < uniqueSeatID)
        {
            highestSeatID = uniqueSeatID;
        }
    }
    return highestSeatID;
}

int findMyMissingSeat(const vector<string> &inputList)
{
    vector<int> seatIDs;
    for (int i = 0; i < inputList.size(); i++)
    {
        string rowStr = inputList[i].substr(0, 7);
        string colStr = inputList[i].substr(7, 10);
        int row = computeRow(rowStr);
        int col = computeCol(colStr);
        int uniqueSeatID = getUniqueSeatID(row, col);
        seatIDs.push_back(uniqueSeatID);
    }

    sort(seatIDs.begin(), seatIDs.end());
    int myID = 0;
    for (int i = 1; i < seatIDs.size(); i++)
    {
        if (seatIDs[i] - seatIDs[i - 1] != 1)
        {
            myID = seatIDs[i] - 1;
            break;
        }
    }
    return myID;
}

int solvePart1(const vector<string> &inputList)
{
    return getHighestSeatID(inputList);
}

int solvePart2(const vector<string> &inputList)
{
    return findMyMissingSeat(inputList);
}

int main()
{
    vector<string> inputList = readInput("./day5.txt");
    cout << solvePart1(inputList) << endl;
    cout << solvePart2(inputList) << endl;
}