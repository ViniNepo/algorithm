package structs.stack.execises;

import java.util.List;
import java.util.Stack;

public class BasketballGame {

    public static void main(String[] args) {
        String[] arr = new String[]{"5","2","C","D","+"};
        String[] arr2 = new String[]{"5","-2","4","C","D","9","+","+"};
        String[] arr3 = new String[]{"1","C"};

        calPoints(arr3);
    }

    public static int calPoints(String[] operations) {
        Stack<Integer> stack = new Stack<>();

        for (int i = 0; i < operations.length; i++) {
            switch (operations[i]) {
                case "+":
                    int lastNum1 = stack.pop();
                    int lastNum2 = stack.peek();
                    stack.push(lastNum1);
                    stack.push(lastNum1 + lastNum2);
                    break;
                case "D":
                    stack.push(stack.peek() * 2);
                    break;
                case "C":
                    stack.pop();
                    break;
                default:
                    stack.push(Integer.parseInt(operations[i]));
            }
        }

        int total = 0;
        while (stack.size() != 0) {
            total += stack.pop();
        }

        System.out.println(total);
        return total;
    }

}
