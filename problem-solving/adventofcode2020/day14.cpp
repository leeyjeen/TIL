// https://adventofcode.com/2020/day/14
#include <algorithm>
#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>
#include <map>
#include <iterator>
#include <regex>
#include <bitset>
#include <numeric>

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

vector<string> split(string input, char delimiter)
    {
    vector<string> answer;
    stringstream ss(input);
    string temp;

    while (getline(ss, temp, delimiter))
        {
        answer.push_back(temp);
        }

    return answer;
    }

/*
PART 1:

mask 0 -> change to 0
mask 1 -> change to 1
mask X -> do not change
value update..
*/

unsigned long long applyMasknig(string mask, string binary)
    {
    string valueAfterMasking = "";
    for (int i = 0; i < mask.size(); i++)
        {
        if (mask[i] == '0' || mask[i] == '1')
            {
            valueAfterMasking += mask[i];
            }
        else if (mask[i] == 'X')
            {
            valueAfterMasking += binary[i];
            }
        }
    return bitset<36>(valueAfterMasking).to_ullong();
    }

unsigned long long solvePart1(vector<string>& lines)
    {
    map<unsigned long long, unsigned long long> memory;
    string mask = "";

    const regex maskReg{ R"(^mask = ([X01]{36})$)" };
    const regex memReg{ R"(^mem\[(\d+)\] = (\d+)$)" };

    for (string& line : lines)
        {
        smatch match;
        if (regex_match(line, match, maskReg))
            {
            mask = match[1];
            }
        else if (regex_match(line, match, memReg))
            {
            unsigned long long pos = stoull(match[1]);
            unsigned long long val = stoull(match[2]);
            string binary = bitset<36>(val).to_string();
            memory[pos] = applyMasknig(mask, binary);
            }
        }

    unsigned long long result = 0;
    for (const auto& m : memory)
        {
        result += m.second;
        }

    return result;
    }

/*
PART 2:

[modifies memory address]
mask 0 -> do not change
mask 1 -> overwrite with 1
mask X -> take on all possible values...!
주소를 먼저 변환한 후, 변환한 주소들에 값을 매핑.

answer: sum of all values left in memory after it completes
*/

void applyMasking2(map<unsigned long long, unsigned long long>& addresses, string mask, unsigned long long addr, unsigned long long val)
    {
    string binary = bitset<36>(addr).to_string();
    map<string, unsigned long long> valuesAfterMasking = { { "" , val } };
    for (int i = 0; i < mask.size(); i++)
        {
        if (mask[i] == '0')
            {
            map<string, unsigned long long> newValuesAfterMasking;

            for (auto& v : valuesAfterMasking)
                {
                newValuesAfterMasking[v.first + binary[i]] = val;
                }
            valuesAfterMasking = newValuesAfterMasking;
            }
        else if (mask[i] == '1')
            {
            map<string, unsigned long long> newValuesAfterMasking;
            for (auto& v : valuesAfterMasking)
                {
                newValuesAfterMasking[v.first + "1"] = val;
                }
            valuesAfterMasking = newValuesAfterMasking;
            }
        else if (mask[i] == 'X')
            {
            map<string, unsigned long long> newValuesAfterMasking;
            for (auto& v : valuesAfterMasking)
                {
                newValuesAfterMasking[v.first + "0"] = val;
                newValuesAfterMasking[v.first + "1"] = val;
                }
            valuesAfterMasking = newValuesAfterMasking;
            }
        }

    for (auto const& v : valuesAfterMasking) {
        addresses[bitset<36>(v.first).to_ullong()] = v.second;
        }
    }

unsigned long long solvePart2(vector<string>& lines)
    {
    map<unsigned long long, unsigned long long> addresses;
    string mask = "";

    const regex maskReg{ R"(^mask = ([X01]{36})$)" };
    const regex memReg{ R"(^mem\[(\d+)\] = (\d+)$)" };

    for (string& line : lines)
        {
        smatch match;
        if (regex_match(line, match, maskReg))
            {
            mask = match[1];
            }
        else if (regex_match(line, match, memReg))
            {
            unsigned long long addr = stoull(match[1]);
            unsigned long long val = stoull(match[2]);
            applyMasking2(addresses, mask, addr, val);
            }
        }

    unsigned long long result = 0;
    for (const auto& a : addresses)
        {
        result += a.second;
        }

    return result;
    }

int main()
    {
    vector<string> lines = readInput("./day14.txt");
    cout << solvePart1(lines) << endl;
    cout << solvePart2(lines) << endl;
    }

/*
C++ 문법 스터디:

1. binary string을 숫자로 변환
    bitset<36>(valueAfterMasking).to_ullong()
2. 숫자를 binary string으로 변환
    bitset<36>(val).to_string()
*/