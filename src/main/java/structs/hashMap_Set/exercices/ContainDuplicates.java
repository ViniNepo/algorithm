package structs.hashMap_Set.exercices;

import java.util.HashSet;
import java.util.Set;

public class ContainDuplicates {

    public static void main(String[] args) {
        int[] nums = new int[]{1,2,3,3,4,5,6,7,8,9,0};

        System.out.println(containsDuplicate(nums));
    }

    public static boolean containsDuplicate(int[] nums) {
        Set<Integer> list = new HashSet<>();

        for (int num : nums) {
            if (!list.contains(num)) {
                list.add(num);
            } else {
                return true;
            }
        }

        return false;
    }
}
