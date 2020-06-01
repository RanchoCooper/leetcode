package leetcode_en.all_005_longest_palindromic_substring;

/**
 * @author rancho
 * @date 2019-05-06
 */
public class BestSolution {

    public String longestPalindrome(String s) {
        // 1. 判断空串
        if (s == null || s.length() == 0) {
            return "";
        }

        // 2. 保存起始位置, 将字符串转为字符串数组
        int[] range = new int[2];
        char[] str = s.toCharArray();

        // 3. 循环字符长度, 将当前字符与字符数组和起始位置传入
        // 查找这个字符串的最大回文长度
        for (int i = 0; i < s.length(); i++) {
            i = findLongestPalindrome(str, i, range);
        }

        return s.substring(range[0], range[1] + 1);
    }

    public static int findLongestPalindrome(char[] str, int low, int[] range) {
        // a. 查找中间部分, 判断可以进行high++
        int high = low;
        while (high < str.length - 1 && str[low] == str[high +1]) {
            high++;
        }

        // b. 定位中间部分的最后一个字符, 进行两边判断
        int ans = high;
        while (low > 0 && high < str.length - 1 && str[low - 1] == str[high + 1]) {
            low--;
            high++;
        }

        // c. 记录最大长度
        if (high - low > range[1] - range[0]) {
            range[0] = low;
            range[1] = high;
        }
        return ans;
    }
}
