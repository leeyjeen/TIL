// https://adventofcode.com/2021/day/1
#include <iostream>
#include <vector>
#include <fstream>
#include <filesystem>
#include <unistd.h>

using namespace std;

vector<int> readInput(string filepath)
{
    vector<int> inputList;
    ifstream inputFile;
    inputFile.open(filepath);

    int inputValue;
    while (inputFile >> inputValue)
    {
        inputList.push_back(inputValue);
    }
    inputFile.close();
    return inputList;
}

int first(const vector<int> &depths)
{
    int curr = depths[0];
    int count = 0;
    for (int i = 1; i < depths.size(); i++)
    {
        int depth = depths[i];
        if (curr < depth)
        {
            count++;
        }
        curr = depth;
    }
    return count;
}

int second(const vector<int> &depths)
{
    int first = depths[0];
    int second = depths[1];
    int third = depths[2];
    int sum = first + second + third;
    int count = 0;
    for (int i = 2; i < depths.size() - 1; i++)
    {
        first = second;
        second = third;
        third = depths[i + 1];
        int currSum = first + second + third;
        if (currSum > sum)
        {
            count++;
        }
        sum = currSum;
    }
    return count;
}

string getCurrPath()
{
    char tmp[256];
    getcwd(tmp, 256);
    string a(tmp);
    return a;
}

int main()
{
    vector<int> depths = readInput(getCurrPath() + "/input.txt");
    cout << first(depths) << endl;
    cout << second(depths) << endl;
}
