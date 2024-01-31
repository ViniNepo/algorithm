package structs.arrays.exercises;

import java.util.ArrayList;
import java.util.List;

public class RangeSumQueryImmutable {
    static int[] array = new int[]{-2, 0, 3, -5, 2, -1};
    static List<Integer> prefix = new ArrayList<>();


    public static void main(String[] args) {
        numArray(array);

        System.out.println(sumRange(0, 2));
        System.out.println(sumRange(2, 5));
        System.out.println(sumRange(0, 5));
    }

    public static void numArray(int[] nums) {
        int total = 0;

        for (int n : nums) {
            total += n;
            prefix.add(total);
        }
    }

    public static int sumRange(int left, int right) {
        int preRight = prefix.get(right);
        int preLeft = left > 0 ? prefix.get(left - 1) : 0;
        return (preRight - preLeft);
    }
}
