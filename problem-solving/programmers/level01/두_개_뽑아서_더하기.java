// https://programmers.co.kr/learn/courses/30/lessons/68644
import java.util.*;

class Solution {
    public int[] solution(int[] numbers) {
        ArrayList<Integer> list = new ArrayList<Integer>();
        for (int i=0; i<numbers.length-1; i++) {
            for (int j=i+1; j<numbers.length; j++) {
                int sum = numbers[i]+numbers[j];
                if (list.indexOf(sum) < 0) {
                    list.add(sum);
                }
            }
        }
        
        int[] answer = new int[list.size()];
        for (int i=0; i<list.size(); i++) {
            answer[i] = list.get(i);
        }
        Arrays.sort(answer);
        return answer;
    }
}

// // 다른 풀이 (set 사용)
// import java.util.HashSet;
// import java.util.Set;

// class Solution {
//      public int[] solution(int[] numbers) {
//         Set<Integer> set = new HashSet<>();

//         for(int i=0; i<numbers.length; i++) {
//             for(int j=i+1; j<numbers.length; j++) {
//                 set.add(numbers[i] + numbers[j]);
//             }
//         }

//         return set.stream().sorted().mapToInt(Integer::intValue).toArray();
//     }
// }