package structs.arrays.exercises;

import java.util.Arrays;

public class RemoveDuplicatesFromSortedArray {
    public static void main(String[] args) {
        int[] array = new int[]{1,1,2,2,2,3,4,5,6,7,7,7,8,9,9,0};

        int k = RemoveDuplicates(array);
        System.out.println(k);
    }


    public static int RemoveDuplicates(int[] nums) {
        int l = 1;

        for (int r = 1; r < nums.length; r++) {
            if (nums[r] != nums[r - 1]) {
                nums[l] = nums[r];
                l += 1;
            }
        }

        System.out.println(Arrays.toString(nums));
        return l;
    }
}
