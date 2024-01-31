package structs.hashMap_Set.exercices;

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]
 */


import java.util.*;

public class GroupAnagram {
    public static void main(String[] args) {
        String[] array = new String[]{"eat","tea","tan","ate","nat","bat"};

        // SOLUTION 1
        //System.out.println(groupAnagrams(array));

        // SOLUTION 2
        System.out.println(groupAnagrams2(array));
    }

    public static List<List<String>> groupAnagrams(String[] strs) {

        Map<String, List<String>> map = new HashMap<>();

        for (String word: strs) {
            char[] chars = word.toCharArray();
            Arrays.sort(chars);
            String sortedWord = new String(chars);

            if(!map.containsKey(sortedWord)) {
                map.put(sortedWord, new ArrayList<>());
            }

            map.get(sortedWord).add(word);
        }

        return new ArrayList<>(map.values());
    }

    public static List<List<String>> groupAnagrams2(String[] strs) {

        Map<String, List<String>> map = new HashMap<>();

        for (String str: strs) {
            char[] arr = new char[26];
            char[] strArr = str.toCharArray();

            for (char c : strArr) {
                arr[c - 'a']++;
            }

            map.computeIfAbsent(String.valueOf(arr), k -> new ArrayList<>()).add(str);
        }

        return new ArrayList<>(map.values());
    }

}
