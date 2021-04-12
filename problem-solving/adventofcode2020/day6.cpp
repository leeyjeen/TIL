// https://adventofcode.com/2020/day/6
#include <algorithm>
#include <fstream>
#include <iostream>
#include <set>
#include <string>
#include <unordered_map>
#include <vector>
#include <iterator>

using namespace std;

vector<string> readInput(string filepath)
{
    vector<string> inputList;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (getline(inputFile, line))
    {
        inputList.push_back(line);
    }

    inputFile.close();
    return inputList;
}

int getUniqueQuestionCount(set<char> &countSet) { return countSet.size(); }

void updateCountSet(string input, set<char> &countSet)
{
    for (int i = 0; i < input.length(); i++)
    {
        countSet.insert(input[i]);
    }
}

int solvePart1(const vector<string> &inputList)
{
    int totalCount = 0;
    set<char> countSet;
    for (int i = 0; i < inputList.size(); i++)
    {
        if (inputList[i] == "")
        {
            totalCount += getUniqueQuestionCount(countSet);
            countSet.clear();
            continue;
        }
        updateCountSet(inputList[i], countSet);
    }
    totalCount += getUniqueQuestionCount(countSet);
    return totalCount;
}

void updateCountMap(string input, unordered_map<char, int> &countMap)
{
    for (int i = 0; i < input.length(); i++)
    {
        if (countMap.count(input[i]))
        {
            countMap[input[i]]++;
        }
        else
        {
            countMap.insert({input[i], 1});
        }
    }
}

int solvePart2(const vector<string> &inputList)
{
    int totalCount = 0;
    unordered_map<char, int> countMap;
    int groupMemberCount = 0;
    for (int i = 0; i < inputList.size(); i++)
    {
        string input = inputList[i];
        if (input == "")
        {
            for (unordered_map<char, int>::iterator it = countMap.begin(); it != countMap.end(); ++it)
            {
                totalCount += (int)(it->second == groupMemberCount);
            }
            groupMemberCount = 0;
            countMap.clear();
            continue;
        }
        groupMemberCount++;
        updateCountMap(input, countMap);
    }
    for (unordered_map<char, int>::iterator it = countMap.begin(); it != countMap.end(); ++it)
    {
        totalCount += (int)(it->second == groupMemberCount);
    }
    return totalCount;
}

int main()
{
    vector<string> inputList = readInput("./day6.txt");
    cout << solvePart1(inputList) << endl;
    cout << solvePart2(inputList) << endl;
}
