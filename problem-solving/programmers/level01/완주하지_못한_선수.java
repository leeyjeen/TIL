// https://programmers.co.kr/learn/courses/30/lessons/42576

import java.util.HashMap;

class Solution {
    public String solution(String[] participant, String[] completion) {
        String answer = "";

        // hash map 사용
        HashMap<String, Integer> hm = new HashMap<>();
        // hashMap.getOrDefault(key, default) -> key가 존재할 경우 해당하는 value, 존재하지 않을 경우 0
        for (String player : participant) hm.put(player, hm.getOrDefault(player, 0) + 1);
        for (String player : completion) hm.put(player, hm.get(player) - 1);

        // hashMap.keySet() 활용
        for (String key : hm.keySet()) {
            if (hm.get(key) != 0) {
                answer = key;
            }
        }
        return answer;
    }
}