package structs.dinamicArray.execises;

import java.util.Arrays;

public class ConcatenationOfArray {

    public static void main(String[] args) {
        int[] array = new int[]{1,2,3,4,5};

        System.out.println(Arrays.toString(getConcatenation(array)));
    }

    public static int[] getConcatenation(int[] nums) {
        int[] newArr = new int[nums.length * 2];

        for (int i = 0; i < nums.length; i++) {
            newArr[i] = nums[i];
            newArr[i + nums.length] = nums[i];
        }

        return newArr;
    }

}
