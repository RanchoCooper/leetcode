package leetcode_en.all_005_longest_palindromic_substring;

/**
 * @author rancho
 * @date 2019-05-06
 */
public class Solution {

    private int start = 0, width = 0;

    public String longestPalindrome(String s) {
        int len = s.length();
        if (len < 2) {
            return s;
        }

        for (int i = 0; i < len - 1; i++) {
            findPalindrome(s, i, i);
            findPalindrome(s, i, i + 1);
        }

        return s.substring(start, start + width);

    }

    public void findPalindrome(String s, int j, int k) {
        while (j >=0 && k < s.length() && s.charAt(j) == s.charAt(k)) {
            j--;
            k++;
        }

        if (width < k - j -1) {
            start = j + 1;
            width = k - j - 1;
        }

    }

}
