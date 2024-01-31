package structs.sorting.insertionSort.execises;

import java.util.Arrays;

public class SortAnArray {

    public static void main(String[] args) {
        int[] arr = new int[]{5,2,3,1};

        System.out.println(Arrays.toString(sortArray(arr)));
    }

    public static int[] sortArray(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            int j = i - 1;
            while (j >= 0 && nums[j+1] < nums[j]) {
                int tmp = nums[j + 1];
                nums[j + 1] = nums[j];
                nums[j] = tmp;
                j--;
            }
        }
        return nums;
    }

}
