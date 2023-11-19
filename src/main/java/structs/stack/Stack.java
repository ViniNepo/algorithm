package structs.stack;

import java.util.ArrayList;
import java.util.List;

public class Stack {

    List<Integer> stack = new ArrayList<>();

    public void push(int value) {
        stack.add(value);
    }

    public int pop() {
        return stack.remove(stack.size() - 1);
    }

    public int peek() {
        return stack.get(stack.size() - 1);
    }

    public int size() {
        return stack.size();
    }

}
