// https://adventofcode.com/2020/day/8
#include <algorithm>
#include <fstream>
#include <iostream>
#include <sstream>
#include <set>
#include <string>
#include <vector>
#include <unordered_map>
#include <iterator>

using namespace std;

class Command
{
public:
    string operation;
    int argument;

    Command(string opr, int arg)
    {
        operation = opr;
        argument = arg;
    }

    bool convertOperation()
    {
        if (operation == "nop")
        {
            operation = "jmp";
            return true;
        }
        else if (operation == "jmp")
        {
            operation = "nop";
            return true;
        }
        return false;
    }
};

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

vector<Command> parseInputsToCommands(vector<string> &inputList)
{
    vector<Command> commands;
    for (int i = 0; i < inputList.size(); i++)
    {
        string input = inputList[i];
        stringstream ss(input);
        string operation;
        string argument;

        getline(ss, operation, ' ');
        getline(ss, argument, ' ');
        Command command(operation, stoi(argument));
        commands.push_back(command);
    }
    return commands;
}

int solvePart1(vector<Command> &commands)
{
    unordered_map<int, bool> visited;
    int idx = 0;
    int accumulator = 0;
    while (visited.count(idx) == 0)
    {
        visited[idx] = true;
        string opr = commands[idx].operation;
        int arg = commands[idx].argument;
        if (opr == "nop")
        {
            idx += 1;
        }
        else if (opr == "acc")
        {
            accumulator += arg;
            idx += 1;
        }
        else if (opr == "jmp")
        {
            idx += arg;
        }
    }
    return accumulator;
}

int solvePart2(vector<Command> &commands)
{
    for (int i = 0; i < commands.size(); i++)
    {
        if (!commands[i].convertOperation())
        {
            continue;
        }
        unordered_map<int, bool> visited;
        int idx = 0;
        int accumulator = 0;
        bool isContinue = false;
        while (idx < commands.size())
        {
            if (visited.count(idx) == 0)
            {
                visited[idx] = true;
            }
            else
            {
                isContinue = true;
                break;
            }
            visited[idx] = true;
            string opr = commands[idx].operation;
            int arg = commands[idx].argument;
            if (opr == "nop")
            {
                idx += 1;
            }
            else if (opr == "acc")
            {
                accumulator += arg;
                idx += 1;
            }
            else if (opr == "jmp")
            {
                idx += arg;
            }
        }
        if (isContinue)
        {
            commands[i].convertOperation();
            continue;
        }
        return accumulator;
    }
    return 0;
}

int main()
{
    vector<string> inputList = readInput("./day8.txt");
    vector<Command> commands = parseInputsToCommands(inputList);
    cout << solvePart1(commands) << endl;
    cout << solvePart2(commands) << endl;
}
