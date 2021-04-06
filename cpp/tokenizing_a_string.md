# Tokenizing a string in C++

string을 Tokenizing한다는 것은 delimiter를 기준으로 string을 split하는 것과 동일한 의미를 갖는다.

그렇다면, string을 tokenize하는 여러 가지 방법을 알아보자!

## `stringstream` 클래스와 `getline` 메서드 사용하기

여기서는`line` string 변수를 `stringstream`에 넣어서 `getline`메서드와 함께 이 안에서 작동하도록 한다. `getline`은 지정된 `char`를 찾을 때까지 `string` 변수에 토큰을 저장한다. 주의: 이 메서드는 single character delimiter가 필요한 경우에만 적용할 수 있다.

```cpp
// Tokenizing a string using stringstream class 
// and getline method
#include <iostream>
#include <string>
#include <vector>
#include <sstream>

using namespace std;

int main()
{
	string line = "Tokenizing a string with stringstream";

	vector<string> words;
	stringstream sstream(line);
	string word;

	while(getline(sstream, word, ' '))
	{
		words.push_back(word);
	}

	for (int i=0; i<words.size(); i++)
		cout << words[i] << endl;
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

## `find()`와 `substr()` 메서드 사용하기

이 방법은 `string` 클래스의 내장 `find` 메서드를 사용한다. `string` 타입의 문자열과 시작 위치를 정수 타입으로 입력한다. 메서드가 전달된 character를 찾으면 첫 번째 character의 위치를 반환한다. 찾지 못하면 `npos`를 반환한다. 마지막 delimiter를 찾을 때까지 `string`을 돌기 위해 `while` 반복문에 `find` 문을 사용한다. delimiter 사이의 부분 문자열을 추출하기 위해 `substr` 함수가 사용되며, 각 반복마다 토큰이 word 벡터에 추가된다. 반복문의 마지막 단계로, `erase` 메서드로 문자열의 처리된 부분을 제거한다.

```cpp
#include <iostream>
#include <string>
#include <vector>

using namespace std;

int main()
{
	string line = "I am who I am and I have the need to be.";
	string delim = " ";
	vector<string> words{};

	size_t pos = 0;
	while ((pos = line.find(delim)) != string::npos)
	{
		words.push_back(line.substr(0, pos));
		line.erase(0, pos + delim.length());
	}

	for (const auto &w: words)
	{
		cout << w << endl;
	}
	return EXIT_SUCCESS;
}
/*
Output:
I
am
who
I
am
and
I
have
the
need
to
be.
*/
```

## `strtok()` 사용하기

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

## `strtok_r()` 사용하기

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

## `std::sregex_token_iterator` 사용하기

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
    - [https://www.delftstack.com/howto/cpp/how-to-parse-string-using-delimeter-in-cpp/#use-stringstream-class-and-getline-method-to-parse-string-using-a-delimiter](https://www.delftstack.com/howto/cpp/how-to-parse-string-using-delimeter-in-cpp/#use-stringstream-class-and-getline-method-to-parse-string-using-a-delimiter)