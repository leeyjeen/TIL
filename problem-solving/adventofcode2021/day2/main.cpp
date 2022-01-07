// https://adventofcode.com/2021/day/2
#include <iostream>
#include <vector>
#include <fstream>
#include <sstream>
#include <filesystem>
#include <unistd.h>

using namespace std;

pair<string, int> split(string input, char delimiter)
{
    pair<string, int> p;
    stringstream ss(input);
    string key;
    string value;

    getline(ss, key, delimiter);
    getline(ss, value, delimiter);

    p.first = key;
    p.second = stoi(value);

    return p;
}

vector<pair<string, int>> readInput(string filepath)
{
    vector<pair<string, int>> commands;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (getline(inputFile, line))
    {
        pair<string, int> p = split(line, ' ');
        commands.push_back(p);
    }

    inputFile.close();
    return commands;
}

int first(const vector<pair<string, int>> &commands)
{
    int xpos = 0;
    int depth = 0;

    for (int i = 0; i < commands.size(); i++)
    {
        string key = commands[i].first;
        int value = commands[i].second;
        if (key == "forward")
        {
            xpos += value;
        }
        else if (key == "down")
        {
            depth += value;
        }
        else if (key == "up")
        {
            depth -= value;
        }
    }
    return xpos * depth;
}

int second(const vector<pair<string, int>> &commands)
{
    int aim = 0;
    int xpos = 0;
    int depth = 0;
    for (int i = 0; i < commands.size(); i++)
    {
        string key = commands[i].first;
        int value = commands[i].second;
        if (key == "forward")
        {
            xpos += value;
            depth += aim * value;
        }
        else if (key == "down")
        {
            aim += value;
        }
        else if (key == "up")
        {
            aim -= value;
        }
    }
    return xpos * depth;
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
    vector<pair<string, int>> commands = readInput(getCurrPath() + "/input.txt");
    cout << first(commands) << endl;
    cout << second(commands) << endl;
}
