/**
 * Description
 * <p>
 * </p>
 * DATE 2020/7/23.
 *
 * @author hongshengjie.
 */
public class Palindrome {

    /**
     * 判断是否是回文数
     *
     * @param x
     * @return
     */
    public static boolean isPalindrome(int x) {
        if (x < 0 || (x % 10 == 0 && x != 0)) { // 负数全部非回文数
            return false;
        }
        int res = 0;
        while (x > res) {
            res = res * 10 + x % 10;
            x /= 10;
        }
        return res == x || res / 10 == x;
    }
}
