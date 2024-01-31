package structs.hashMap_Set.exercices;

import java.util.HashMap;
import java.util.Objects;

public class ValidAnagram {

    public static void main(String[] args) {
        String s = "abcdefghij";
        String t = "aabcdefghijr";

        System.out.println(isAnagram2(s, t));
    }

    public static boolean isAnagram(String s, String t) { // my version
        HashMap<Character, Integer> tHash = new HashMap<>();
        HashMap<Character, Integer> sHash = new HashMap<>();

        for (char letter : t.toCharArray()) {
            if (!tHash.containsKey(letter)) {
                tHash.put(letter, 1);
            } else {
                tHash.put(letter, tHash.get(letter) + 1);
            }
        }

        for (char letter : s.toCharArray()) {
            if (!sHash.containsKey(letter)) {
                sHash.put(letter, 1);
            } else {
                sHash.put(letter, sHash.get(letter) + 1);
            }
        }

        if (sHash.size() != tHash.size()) {
            return false;
        }

        for (char letter : t.toCharArray()) {
            if (sHash.containsKey(letter) != tHash.containsKey(letter) || !Objects.equals(sHash.get(letter), tHash.get(letter))) {
                return false;
            }
        }
        return true;
    }

    public static boolean isAnagram2(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        int[] store = new int[26];

        for (int i = 0; i < s.length(); i++) {
            store[s.charAt(i) - 'a']++;
            store[t.charAt(i) - 'a']--;
        }

        for (int n : store) {
            if (n != 0) {
                return false;
            }
        }

        return true;
    }

}
