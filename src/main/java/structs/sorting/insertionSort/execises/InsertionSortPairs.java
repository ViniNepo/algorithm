package structs.sorting.insertionSort.execises;

import java.util.ArrayList;
import java.util.List;

public class InsertionSortPairs {

    class Pair {
     int key;
     String value;

     Pair(int key, String value) {
         this.key = key;
         this.value = value;
     }
 }

    public List<List<Pair>> insertionSort(List<Pair> pairs) {
        List<List<Pair>> list = new ArrayList<>();

        for (int i = 1; i < pairs.size(); i++) {

            int j = i - 1;
            while (j >= 0 && pairs.get(j + 1).key < pairs.get(j).key) {
                Pair temp = pairs.get(j);
                pairs.set(j, pairs.get(j + 1));
                pairs.set(j + 1, temp);
                j--;
            }

            List<Pair> status = new ArrayList<>(pairs);
            list.add(status);
        }

        return list;
    }
}
