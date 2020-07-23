/**
 * Description
 * <p>
 * </p>
 * DATE 2020/7/23.
 *
 * @author hongshengjie.
 */
public class ReverseNumber {

    /**
     * 翻转数字
     *
     * @param x
     * @return
     */
    public static int reverse(int x) {
        int rev = 0;
        while (x != 0) {
            // 获取余数
            int pop = x % 10;
            // 去除一位
            x /= 10;

            // 避免数据溢出
            if (rev > Integer.MAX_VALUE/10 || (rev == Integer.MAX_VALUE / 10 && pop > 7))
                return 0;
            if (rev < Integer.MIN_VALUE/10 || (rev == Integer.MIN_VALUE / 10 && pop < -8))
                return 0;

            rev = rev * 10 + pop;
        }
        return rev;
    }
}
