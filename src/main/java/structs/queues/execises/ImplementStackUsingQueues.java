package structs.queues.execises;

import java.util.ArrayDeque;
import java.util.Deque;

public class ImplementStackUsingQueues {

    public static void main(String[] args) {
        ImplementStackUsingQueues stack = new ImplementStackUsingQueues();

        stack.push(1);
        stack.push(2);

        System.out.println(stack.top());
    }

    Deque<Integer> stack;

    public ImplementStackUsingQueues() {
        this.stack = new ArrayDeque<>();
    }

    public void push(int x) {
        this.stack.addLast(x);
    }

    public int pop() {
        int size = this.stack.size();

        for (int i = 0; i < size - 1; i++) {
            this.stack.addLast(this.stack.pollFirst());
        }

        return this.stack.pollFirst();
    }

    public int top() {
        int size = stack.size();

        for (int i = 0; i < size - 1; i++) {
            this.stack.addLast(stack.pollFirst());
        }

        int result = this.stack.peekFirst();
        this.stack.addLast(this.stack.pollFirst());
        return result;
    }

    public boolean empty() {
        return this.stack.size() == 0;
    }
}
