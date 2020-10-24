// https://programmers.co.kr/learn/courses/30/lessons/42840

import java.util.*;

class Solution {
    public int[] solution(int[] answers) {
        int[] answer = {};

        int[][] student = {
            {1,2,3,4,5},
            {2,1,2,3,2,4,2,5},
            {3,3,1,1,2,2,4,4,5,5}
        };

        int check[] = new int[3];

        // 정답 수 체크
        for (int i=0; i<answers.length; i++) {
            for (int j=0; j<3; j++) {
                if (answers[i] == student[j][i%student[j].length]) {
                    check[j]++;
                }
            }
        }

        int maxScore = 0;
        for (int i=0; i<check.length; i++) {
            if (maxScore < check[i]) {
                maxScore = check[i];
            }
        }

        int size = 0;
        for (int i=0; i<check.length; i++) {
            if (check[i] == maxScore) {
                size++;
            }
        }

        answer = new int[size];
        int num = 0;
        for (int i=0; i<check.length; i++) {
            if (check[i]==maxScore) {
                answer[num++] = i+1;
            }
        }
        return answer;
    }
}
