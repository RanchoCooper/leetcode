package leetcode_cn.No3.longest.substring.without.repeating.characters;

import java.util.HashMap;
import java.util.Map;

/**
 * @author rancho
 * @date 2019/10/9
 */
public class Solution {

    public int lengthOfLongestSubstring(String s) {
        int n = s.length();
        int ans = 0;

        Map<Character, Integer> map = new HashMap<>();

        for (int i = 0, j = 0; j < n; j++) {
            if (map.containsKey(s.charAt(j))) {
                i = Math.max(map.get(s.charAt(j)), i);
            }
            map.put(s.charAt(j), j + 1);
            ans = Math.max(ans, j - i + 1);
        }
        return ans;
    }

}
