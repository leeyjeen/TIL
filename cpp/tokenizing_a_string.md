# Tokenizing a string in C++

string을 Tokenizing한다는 것은 delimiter를 기준으로 string을 split하는 것과 동일한 의미를 갖는다.

그렇다면, string을 tokenize하는 여러 가지 방법을 알아보자!

## stringstream 사용하기

`stringstream`은 string 오브젝트를 stream과 연결하여 stream처럼 string을 읽을 수 있게 한다.

```cpp
// Tokenizing a string using stringstream
#include <bits/stdc++.h>

using namespace std;

int main()
{
	string line = "Tokenizing a string with stringstream";

	vector<string> tokens;
	stringstream check1(line);
	string intermediate;

	while(getline(check1, intermediate, ' '))
	{
		tokens.push_back(intermediate);
	}

	for (int i=0; i<tokens.size(); i++)
		cout << tokens[i] << '\n';
}
/*
Output:
Tokenizing
a
string
with
stringstream
*/
```

## strtok() 사용하기

`char * strtok(char str[], const char *delims);`

→ str[]을 주어진 delimeters에 따라 split하고, next token을 반환한다.

→ 모든 token을 얻기 위해서 반복문으로 호출되어야 한다.

→ token이 더 이상 없을 때는 NULL을 반환한다.

```cpp
#include <stdio.h>
#include <string.h>

int main()
{
	char str[] = "Lee-Yu-Jin";
	char *token = strtok(str,"-"); // 첫 번째 token을 반환한다.
	
	while (token != NULL) // delimiter 중 하나가 str[]에 존재할 경우 계속해서 token을 출력한다.
	{
		printf("%s\n", token);
		token = strtok(NULL, "-");
	}
	return 0;
}
/*
Output:
Lee
Yu
Jin
*/
```

```cpp
#include <string.h>
#include <stdio.h>

int main()
{
	char lyj[100] = " Lee - Yu - Jin";

	const char s[4] = "-";
	char* tok;

	tok = strtok(lyj, s);

	while (tok != 0)
	{
		printf(" %s\n", tok);
		tok = strtok(0, s);
	}

	return 0;
}

/*
Output:
  Lee
  Yu
  Jin
*/
```

## strtok_r() 사용하기

`char *strtok_r(char *str, const char *delim, char **saveptr);`
→ 세번째 인수 saveptr은 strtok_r()에서 같은 string을 파싱하기 위해 연속적인 호출 사이의 컨텍스트를 유지하기 위하여 내부적으로 사용하는변수 char*에 대한 포인터이다. 

```cpp
#include <stdio.h>
#include <string.h>

int main()
{
	char str[] = "Lee Yu Jin";
	char *token;
	char *rest = str;

	while ((token = strtok_r(rest, " ", &rest)))
		printf("%s\n", token);
	return 0;
}
/*
Output:
Lee
Yu
Jin
*/
```

## std::sregex_token_iterator 사용하기

정규식 일치(regex match)에 근거하여 tokenization을 하는 방법이다.

여러 delimiter가 필요한 상황에 적합하게 사용할 수 있다.

```cpp
#include <iostream>
#include <regex>
#include <string>
#include <vector>

std::vector<std::string> tokenize(const std::string str, const std::regex re)
{
	std::sregex_token_iterator it{str.begin(), str.end(), re, -1}; // -1: values between separators
	std::vector<str::string> tokenized{it, {}};

	tokenized.erase(std::remove_if(tokenized.begin(), 
								tokenized.end(), [](std::string const& s) {
																		return s.size() == 0;
																}),
									tokenized.end());
	return tokenized;
}

int main()
{
	const std::string str = "Break string a,spaces,and,commas";
	const std::regex re(R"([\s|,]+)");

	const std::vector<std::string> tokenized = tokenize(str, re);

	for (std::string token : tokenized)
		std::cout << token << stod::endl;
	return 0;
}

/*
Output:
Break
string
a
spaces
and
commas
*/
```

```cpp
#include <regex>
#include <algorithm>
#include <iostream>

int main()
{
    std::string str("version 1.20.300.400");
    std::regex pattern("([0-9]+)\\.([0-9]+)\\.([0-9]+)\\.([0-9]+)");
    {
        std::sregex_token_iterator begin(str.begin(), str.end(), pattern, {1,3});
        std::sregex_token_iterator end;
        std::for_each(begin, end, [](const std::string& sub) {
             std::cout << sub << std::endl; // 1 and 300
        });
    }
    {
        std::sregex_token_iterator begin(str.begin(), str.end(), pattern, {2,4});
        std::sregex_token_iterator end;
        std::for_each(begin, end, [](const std::string& sub) {
             std::cout << sub << std::endl; // 20 and 400
        });
    }
    {
        std::sregex_token_iterator begin(str.begin(), str.end(), pattern, 0);
        std::sregex_token_iterator end;
        std::for_each(begin, end, [](const std::string& sub) {
             std::cout << sub << std::endl; // 1.20.300.400
        });
    }
    {
        std::sregex_token_iterator begin(str.begin(), str.end(), pattern, -1);
        std::sregex_token_iterator end;
        std::for_each(begin, end, [](const std::string& sub) {
             std::cout << sub << std::endl; // version
        });
    }
    return 0;
}
```

- 참고:
    - [https://www.geeksforgeeks.org/tokenizing-a-string-cpp/](https://www.geeksforgeeks.org/tokenizing-a-string-cpp/)
    - [https://gist.github.com/cutlassfish/f59a4c7b96bccc18f2f02feeb14c4f3d](https://gist.github.com/cutlassfish/f59a4c7b96bccc18f2f02feeb14c4f3d)