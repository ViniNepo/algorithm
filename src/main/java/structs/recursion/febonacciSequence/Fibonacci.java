package structs.recursion.febonacciSequence;

import java.util.Arrays;

public class Fibonacci {

    public static void main(String[] args) {
        System.out.println(fib(3));
    }

    public static int fibonacci(int n) {
        if (n <= 1) {
            return n;
        }

        return fibonacci(n - 1) + fibonacci(n - 2);
    }

    public static long fib(int nth){
        long[] table = new long[nth + 1];
        Arrays.fill(table, -1);
        return run(table, nth);
    }

    public static long run(long[] table, int nth){
        if(table[nth] != -1){
            return table[nth];
        }
        if(nth == 0){
            return 0;
        }
        if(nth < 3){
            table[nth] = 1;
            return table[nth];
        }
        table[nth] = run(table, nth-1) + run(table, nth-2);
        return table[nth];
    }
}
