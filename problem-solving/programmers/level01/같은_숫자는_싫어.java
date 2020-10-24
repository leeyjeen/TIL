// https://programmers.co.kr/learn/courses/30/lessons/12906

import java.util.*;

public class Solution {
    public int[] solution(int[] arr) {
        ArrayList<Integer> tempList = new ArrayList<Integer>();
        int prevNum = 10;
        for (int num : arr) {
            if (num != prevNum) {
                tempList.add(num);
            }
            prevNum = num;
        }
        int[] answer = new int[tempList.size()];
        for (int i=0; i<answer.length; i++) {
            answer[i] = tempList.get(i).intValue();
        }
        return answer;
    }
}