// https://adventofcode.com/2020/day/10
#include <algorithm>
#include <fstream>
#include <iostream>
#include <vector>
#include <map>

using namespace std;

vector<int> parseInputsAndSort(string filepath)
{
    vector<int> inputList;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (getline(inputFile, line))
    {
        inputList.push_back(stoi(line));
    }
    inputFile.close();
    sort(inputList.begin(), inputList.end());
    return inputList;
}

int solvePart1(vector<int> &numbers)
{
    int prev = 0;
    int oneJoltCount = 0;
    int threeJoltCount = 0;
    for (int &num : numbers)
    {
        oneJoltCount += (int)(num - prev == 1);
        threeJoltCount += (int)(num - prev == 3);
        prev = num;
    }
    threeJoltCount++;
    return oneJoltCount * threeJoltCount;
}

long solvePart2(vector<int> &numbers)
{
    map<int, long> dp;
    dp[0] = 1;
    for (int &num : numbers)
    {
        dp[num] = dp[num - 1] + dp[num - 2] + dp[num - 3];
    }
    return dp[numbers[numbers.size() - 1]];
}

int main()
{
    vector<int> numbers = parseInputsAndSort("./day10.txt");
    cout << solvePart1(numbers) << endl;
    cout << solvePart2(numbers) << endl;
}