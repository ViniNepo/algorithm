package structs.arrays;

import java.util.ArrayList;
import java.util.List;

public class PrefixSum {

    List<Integer> prefix;

    public void prefixSum(int[] nums) {
        prefix = new ArrayList<>();
        int total = 0;

        for (int n : nums) {
            total += n;
            prefix.add(total);
        }
    }

    public int rangeSum(int left, int right) {
        int preRight = prefix.get(right);
        int preLeft = left > 0 ? prefix.get(left - 1) : 0;
        return (preRight - preLeft);
    }

}
