// https://adventofcode.com/2020/day/16
#include <algorithm>
#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>
#include <map>
#include <regex>

using namespace std;

vector<string> readInput(string filepath)
    {
    vector<string> lines;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (getline(inputFile, line))
        {
        lines.push_back(line);
        }

    inputFile.close();
    return lines;
    }

int solvePart1(vector<string>& lines) {
    const regex rule{ R"(^(.*): (\d+)-(\d+) or (\d+)-(\d+)$)" };
    int invalid_sum = 0;
    map<int, bool> valid_numbers;
    bool is_nearby_ticket = false;
    for (string& line : lines)
        {
        smatch match;
        if (regex_match(line, match, rule))
            {
            for (int i = stoi(match[2]); i <= stoi(match[3]);i++)
                {
                valid_numbers[i] = true;
                }
            for (int i = stoi(match[4]); i <= stoi(match[5]);i++)
                {
                valid_numbers[i] = true;
                }
            }
        else
            {
            if (line == "nearby tickets:")
                {
                is_nearby_ticket = true;
                continue;
                }

            if (is_nearby_ticket)
                {
                stringstream ss(line);
                string num;
                while (getline(ss, num, ','))
                    {
                    if (!valid_numbers[stoi(num)])
                        {
                        invalid_sum += stoi(num);
                        }
                    }
                }
            }
        }
    return invalid_sum;
    }
int main()
    {
    vector<string> lines = readInput("./day16.txt");
    cout << solvePart1(lines) << endl;
    }