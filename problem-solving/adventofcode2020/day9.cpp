// https://adventofcode.com/2020/day/9
#include <algorithm>
#include <fstream>
#include <iostream>
#include <vector>

using namespace std;

vector<long> readInput(string filepath)
{
    vector<long> inputList;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (getline(inputFile, line))
    {
        inputList.push_back(stol(line));
    }

    inputFile.close();
    return inputList;
}

long solvePart1(vector<long> &inputList, int preambleLength)
{
    for (int i = preambleLength; i < inputList.size(); i++)
    {
        long numToCheck = inputList[i];
        bool isValid = false;
        for (int j = i - preambleLength; j < i - 1; j++)
        {
            long numToCheck2 = inputList[j];
            if ((find(inputList.begin() + i - preambleLength, inputList.begin() + j, abs(numToCheck - numToCheck2)) != inputList.begin() + j ||
                 (find(inputList.begin() + j, inputList.begin() + i, abs(numToCheck - numToCheck2)) != inputList.begin() + i)))
            {
                isValid = true;
                break;
            }
        }
        if (!isValid)
        {
            return numToCheck;
        }
    }
    return 0;
}

vector<vector<long> > computeSumDP(vector<long> &inputList)
{
    vector<vector<long> > dp(inputList.size(), vector<long>(inputList.size(), 0)); // i~j까지의 합을 저장할 vector
    for (int i = 0; i < inputList.size(); i++)
    {
        dp[i][i] = inputList[i];
    }
    for (int i = 0; i < inputList.size(); i++)
    {
        for (int j = i + 1; j < inputList.size(); j++)
        {
            dp[i][j] = dp[i][j - 1] + inputList[j];
        }
    }
    return dp;
}

long solvePart2(vector<long> &inputList, long firstNumber)
{
    vector<vector<long> > dp = computeSumDP(inputList);
    for (int i = 0; i < inputList.size(); i++)
    {
        for (int j = 0; j < inputList.size(); j++)
        {
            if (dp[i][j] == firstNumber)
            {
                return *min_element(inputList.begin() + i, inputList.begin() + j + 1) + *max_element(inputList.begin() + i, inputList.begin() + j + 1);
            }
        }
    }
    return 0;
}

int main()
{
    vector<long> inputList = readInput("./day9.txt");
    long firstNumber = solvePart1(inputList, 25);
    cout << firstNumber << endl;
    cout << solvePart2(inputList, firstNumber) << endl;
}

/*
cpp study:
1.
terminating with uncaught exception of type std::out_of_range: stoi: out of range
out of range error 발생 -> int 대신 long, stoi 대신 stol 사용.

2.
vector가 특정 값을 포함하는지 확인하고 싶을 때
if(std::find(v.begin(), v.end(), x) != v.end())

3.
2차원 vector 초기화
vector<vector<long> > dp(inputList.size(), vector<long>(inputList.size(), 0));
*/