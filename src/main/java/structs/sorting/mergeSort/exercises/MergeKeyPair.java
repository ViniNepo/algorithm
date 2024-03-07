package structs.sorting.mergeSort.exercises;

import java.util.List;

public class MergeKeyPair {

    public static void main(String[] args) {

    }

    class Pair {
        public int key;
        public String value;

        public Pair(int key, String value) {
            this.key = key;
            this.value = value;
        }
    }

    public List<Pair> mergeSort(List<Pair> pairs) {
        if (pairs.size() > 1) {
            int m = pairs.size() / 2;

            mergeSort(pairs.subList(0, m));
            mergeSort(pairs.subList(m + 1, pairs.size()));
//            merge(pairs);
        }

        return pairs;
    }
}
