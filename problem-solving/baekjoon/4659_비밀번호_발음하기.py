# https://www.acmicpc.net/problem/4659
import sys

def check_password_acceptable(password):
    prev_letter = ""
    total_vowel_count, vowel_count, consonant_count = 0, 0, 0
    for i in range(len(password)):
        cur_letter = password[i]
        if is_same(prev_letter, cur_letter):
            return False
        if is_vowel(cur_letter):
            vowel_count += 1
            total_vowel_count += 1
            if vowel_count == 3:
                return False
            consonant_count = 0
        else:
            consonant_count += 1
            if consonant_count == 3:
                return False
            vowel_count = 0
        prev_letter = cur_letter
    if total_vowel_count == 0:
        return False
    return True

def is_vowel(letter):
    vowels = ["a", "e", "i", "o", "u"]
    if letter in vowels:
        return True
    return False

def is_same(prev_letter, letter):
    if prev_letter == letter and letter != "e" and letter != "o":
        return True
    return False

if __name__ == "__main__":
    while True:
        password = sys.stdin.readline().rstrip()
        if password == "end":
            break

        if check_password_acceptable(password):
            print("<{}> is acceptable.".format(password))
        else:
            print("<{}> is not acceptable.".format(password))