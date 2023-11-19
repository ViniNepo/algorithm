package structs.arrays.exercises;

import java.util.Arrays;

public class RemoveElements {

    public static void main(String[] args) {
        int[] array = new int[]{1,1,2,2,2,3,4,5,6,7,7,7,8,9,9,0};

        int k = RemoveDuplicates(array, 2);
        System.out.println(k);
    }

    public static int RemoveDuplicates(int[] nums, int val) {
        int k = 0;

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != val) {
                nums[k] = nums[i];
                k += 1;
            }
        }

        System.out.println(Arrays.toString(nums));
        return k;
    }

}
