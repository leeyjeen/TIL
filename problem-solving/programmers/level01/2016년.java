// https://programmers.co.kr/learn/courses/30/lessons/12901
import java.util.*;

class Solution {
    public String solution(int a, int b) {
        String answer = "";

        int[] days = {31,29,31,30,31,30,31,31,30,31,30,31};
        String[] weekdays = {"FRI","SAT","SUN","MON","TUE","WED","THU"};
        int resultDay = 0;
        for (int i=0; i<a-1; i++) {
            resultDay += days[i];
        }
        resultDay += b-1;
        answer = weekdays[resultDay%7];
        return answer;
    }
}