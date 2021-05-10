// https://adventofcode.com/2020/day/15
#include <iostream>
#include <vector>
#include <map>
using namespace std;

long solvePart1(vector<long>& numbers)
    {
    while (numbers.size() < 2020)
        {
        if (count(numbers.begin(), numbers.end(), numbers[numbers.size() - 1]) == 1)
            {
            numbers.push_back(0);
            }
        else
            {
            vector<int> spokenTurns;
            for (int i = 0;i < numbers.size();i++)
                {
                if (numbers[i] == numbers[numbers.size() - 1]) {
                    spokenTurns.push_back(i + 1);
                    }
                }
            numbers.push_back(spokenTurns[spokenTurns.size() - 1] - spokenTurns[spokenTurns.size() - 2]);

            }

        }
    return numbers[numbers.size() - 1];
    }

/*
[0,3,6]
0 3 6
**Turn 4**
    last: 6
    number: 0 (because 6 is first..)
**Turn 5**
    last: 0
    number: 4-1 = 3 (0's turn...)
...
*/
long solvePart2(vector<long>& numbers)
    {
    map<long, long> spoken; // 각 숫자가 등장한 마지막 turn을 저장
    for (int i = 0;i < numbers.size();i++)
        {
        spoken[numbers[i]] = i + 1;
        }
    long lastNum = numbers[numbers.size() - 1];
    bool appearedFirst = true;
    long turnDiff = 0;
    for (int curTurn = numbers.size() + 1; curTurn < 30000001; curTurn++)
        {
        if (appearedFirst) // 처음 나타난 경우 speak할 숫자는 0이다
            {
            lastNum = 0;
            }
        else    // 이전에 나타난 적이 있을 경우 speak할 숫자는 turn의 차이값이다
            {
            lastNum = turnDiff;
            }
        appearedFirst = (spoken.find(lastNum) == spoken.end()); // speak하는 숫자가 이전에 나타난 적이 없는지 확인
        if (!appearedFirst)
            {
            turnDiff = curTurn - spoken[lastNum]; // 다음 turn에서 저장할 값을 계산 (현재 speak하는 숫자의 diff)
            }
        spoken[lastNum] = curTurn; // speak하는 숫자의 마지막 turn을 업데이트
        }


    return lastNum;
    }

int main()
    {
    vector<long> input{ 0, 14, 1, 3, 7, 9 };
    cout << solvePart1(input) << endl;
    cout << solvePart2(input) << endl;
    }