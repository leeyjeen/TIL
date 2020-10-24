// https://programmers.co.kr/learn/courses/30/lessons/64061
// 2019 카카오 개발자 겨울 인턴십

import java.util.Stack;

class Solution {
    public int solution(int[][] board, int[] moves) {
        int answer = 0;
        Stack<Integer> stack = new Stack<>();
        for (int move : moves) {
            for (int j=0; j<board.length; j++) {
                if (board[j][move-1] != 0) {
                    if (stack.isEmpty()) {
                        stack.push(board[j][move-1]);
                        board[j][move-1] = 0;
                        break;
                    }
                    if (board[j][move-1] == stack.peek()) {
                        stack.pop();
                        answer += 2;
                    } else {
                        stack.push(board[j][move-1]);
                    }
                    board[j][move-1] = 0;
                    break;
                }
            }
        }
        return answer;
    }
}
