// https://adventofcode.com/2020/day/11
#include <algorithm>
#include <fstream>
#include <iostream>
#include <vector>

using namespace std;

const int PART_ONE = 1;
const int PART_TWO = 2;

vector<pair<int, int> > eightDirections;

vector<vector<char> > readInput(string filepath)
{
    vector<vector<char> > graph;
    ifstream inputFile;
    inputFile.open(filepath);

    string line;
    while (getline(inputFile, line))
    {
        vector<char> row;
        for (int i = 0; i < line.length(); i++)
        {
            row.push_back(line[i]);
        }
        graph.push_back(row);
    }

    inputFile.close();
    return graph;
}

size_t getAdjacentSeatsCount(vector<vector<char> > &seats, int row, int col)
{
    size_t count = 0;
    for (pair<int, int> &d : eightDirections)
    {
        int diffRow = d.first;
        int diffCol = d.second;
        int r = row + diffRow;
        int c = col + diffCol;
        if (r >= 0 && r < seats.size() && c >= 0 && c < seats[0].size())
        {
            if (seats[r][c] == '#')
            {
                count++;
            }
        }
    }
    return count;
}

vector<vector<char> > getSeatsAfterRound(vector<vector<char> > &seats, size_t (*function)(vector<vector<char> > &, int, int), int becomeEmptyCriteria)
{
    vector<vector<char> > newSeats;
    for (int r = 0; r < seats.size(); r++)
    {
        vector<char> newRow;
        for (int c = 0; c < seats[0].size(); c++)
        {
            char seat = seats[r][c];
            if (seat == 'L')
            {
                if (function(seats, r, c) == 0)
                {
                    newRow.push_back('#');
                }
                else
                {
                    newRow.push_back('L');
                }
            }
            else if (seat == '#')
            {
                if (function(seats, r, c) >= becomeEmptyCriteria)
                {
                    newRow.push_back('L');
                }
                else
                {
                    newRow.push_back('#');
                }
            }
            else
            {
                newRow.push_back('.');
            }
        }
        newSeats.push_back(newRow);
    }
    return newSeats;
}

size_t getEightDirectionSeatsCount(vector<vector<char> > &seats, int row, int col)
{
    size_t count = 0;

    for (pair<int, int> &d : eightDirections)
    {
        int diffRow = d.first;
        int diffCol = d.second;
        int r = row + diffRow;
        int c = col + diffCol;
        while (r >= 0 && r < seats.size() && c >= 0 && c < seats[0].size())
        {
            char seat = seats[r][c];
            if (seat == '#')
            {
                count++;
                break;
            }
            else if (seat == 'L')
            {
                break;
            }
            else
            {
                r += diffRow;
                c += diffCol;
            }
        }
    }
    return count;
}

size_t solve(int part, vector<vector<char> > &seats)
{
    vector<vector<char> > prevSeats(seats);
    vector<vector<char> > newSeats;
    size_t occupiedCount = 0;
    while (true)
    {
        if (part == PART_ONE)
        {
            newSeats = getSeatsAfterRound(prevSeats, &getAdjacentSeatsCount, 4);
        }
        else if (part == PART_TWO)
        {
            newSeats = getSeatsAfterRound(prevSeats, &getEightDirectionSeatsCount, 5);
        }

        if (prevSeats == newSeats)
        {
            for (int i = 0; i < newSeats.size(); i++)
            {
                for (int j = 0; j < newSeats[0].size(); j++)
                {
                    occupiedCount += int(newSeats[i][j] == '#');
                }
            }
            break;
        }
        prevSeats = newSeats;
    }
    return occupiedCount;
}

int main()
{
    vector<vector<char> > seats = readInput("./day11.txt");

    eightDirections = {make_pair(-1, -1), make_pair(-1, 0),
                       make_pair(0, -1), make_pair(1, 1), make_pair(1, 0),
                       make_pair(0, 1), make_pair(-1, 1), make_pair(1, -1)};

    cout << solve(PART_ONE, seats) << endl;
    cout << solve(PART_TWO, seats) << endl;
}
