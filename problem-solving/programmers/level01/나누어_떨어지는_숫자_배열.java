// https://programmers.co.kr/learn/courses/30/lessons/12910
import java.util.Arrays;

class Solution {
    public int[] solution(int[] arr, int divisor) {
        int[] answer = Arrays.stream(arr).filter(e->e%divisor==0).toArray();
        if (answer.length == 0) {
            answer = new int[]{-1};
        }
        Arrays.sort(answer);
        return answer;
    }
}