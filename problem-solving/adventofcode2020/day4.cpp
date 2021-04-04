// https://adventofcode.com/2020/day/4
#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <iterator>
#include <map>
#include <regex>

using namespace std;

vector<map<string, string> > parseInput(string filepath)
{
    vector<map<string, string> > passports;
    ifstream inputFile;
    inputFile.open(filepath);

    map<string, string> passport;
    string line;
    const regex reg("([a-z]{3}):([^\\s]+)");
    while (getline(inputFile, line))
    {
        sregex_token_iterator iter(line.begin(), line.end(), reg, {1, 2});
        while (iter != sregex_token_iterator())
        {
            string key = *iter;
            iter = next(iter);
            string val = *iter;
            passport[key] = val;
            iter = next(iter);
        }

        if (line == "")
        {
            passports.push_back(passport);
            passport.clear();
            continue;
        }
    }
    passports.push_back(passport);
    passport.clear();
    inputFile.close();
    return passports;
}

bool validateByr(map<string, string> &passport)
{
    if (passport["byr"] == "")
        return false;
    int byr = stoi(passport["byr"]);
    if (byr >= 1920 && byr <= 2002)
        return true;
    return false;
}

bool validateIyr(map<string, string> &passport)
{
    if (passport["iyr"] == "")
        return false;
    int iyr = stoi(passport["iyr"]);
    if (iyr >= 2010 && iyr <= 2020)
        return true;
    return false;
}

bool validateEyr(map<string, string> &passport)
{
    if (passport["eyr"] == "")
        return false;
    int eyr = stoi(passport["eyr"]);
    if (eyr >= 2020 && eyr <= 2030)
        return true;
    return false;
}

bool validateHgt(map<string, string> &passport)
{
    if (passport["hgt"] == "")
        return false;
    const regex cmRegex("(^[0-9]{3}cm$)");
    const regex inRegex("(^[0-9]{2}in$)");

    sregex_token_iterator iter(passport["hgt"].begin(), passport["hgt"].end(), cmRegex, 1);
    if (iter != sregex_token_iterator())
    {
        int hgt = stoi(*iter);
        if (hgt >= 150 && hgt <= 193)
        {
            return true;
        }
    }

    iter = {passport["hgt"].begin(), passport["hgt"].end(), inRegex, 1};
    if (iter != sregex_token_iterator())
    {
        int hgt = stoi(*iter);
        if (hgt >= 59 && hgt <= 76)
            return true;
    }
    return false;
}

bool validateHcl(map<string, string> &passport)
{
    if (passport["hcl"] == "")
        return false;
    const regex reg("(^#[0-9a-f]{6}$)");
    sregex_token_iterator iter(passport["hcl"].begin(), passport["hcl"].end(), reg, 1);
    if (iter != sregex_token_iterator())
        return true;
    return false;
}

bool validateEcl(map<string, string> &passport)
{
    if (passport["ecl"] == "")
        return false;
    const vector<string> ecls{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"};
    for (const string &ecl : ecls)
        if (passport["ecl"] == ecl)
            return true;
    return false;
}

bool validatePid(map<string, string> &passport)
{
    if (passport["pid"] == "")
        return false;
    const regex reg("^[0-9]{9}$");
    sregex_token_iterator iter(passport["pid"].begin(), passport["pid"].end(), reg, 1);
    if (iter != sregex_token_iterator())
        return true;
    return false;
}

bool hasAllRequiredFields(map<string, string> &passport)
{
    vector<string> required = {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"};
    for (string &key : required)
    {
        if (passport[key] == "")
            return false;
    }
    return true;
}

bool hasAllRequiredFieldsAndValid(map<string, string> &passport)
{
    return (validateByr(passport) && validateIyr(passport) && validateEyr(passport) && validateHgt(passport) && validateHcl(passport) && validateEcl(passport) && validatePid(passport));
}

int solvePart1(vector<map<string, string> > &passports)
{
    int validCount = 0;
    for (int i = 0; i < passports.size(); i++)
    {
        map<string, string> passport = passports[i];
        if (hasAllRequiredFields(passport))
            validCount++;
    }
    return validCount;
}

int solvePart2(vector<map<string, string> > &passports)
{
    int validCount = 0;
    for (int i = 0; i < passports.size(); i++)
    {
        map<string, string> passport = passports[i];
        if (hasAllRequiredFieldsAndValid(passport))
            validCount++;
    }
    return validCount;
}

int main()
{
    vector<map<string, string> > passports = parseInput("./day4.txt");
    cout << solvePart1(passports) << endl;
    cout << solvePart2(passports) << endl;
}
