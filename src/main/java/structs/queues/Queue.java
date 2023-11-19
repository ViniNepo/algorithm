package structs.queues;

public class Queue {

    public static void main(String[] args) {
        Queue queue = new Queue();

        queue.enqueue(1);
        queue.dequeue();
    }

    class Node {
        int value;
        Node next;

        public Node(int value) {
            this.value = value;
        }
    }

    Node left;
    Node right;

    public Queue() {
        this.left = null;
        this.right = null;
    }

    public void enqueue(int val) {
        Node newNode = new Node(val);
        if (this.right != null) {
            this.right.next = newNode;
            this.right = newNode;
        } else {
            this.left = newNode;
            this.right = newNode;
        }
    }

    public Integer dequeue() {
        if (this.left == null) {
            return null;
        }

        int val = this.left.value;
        this.left = this.left.next;

        if (this.left == null) {
            this.right = null;
        }

        return val;
    }

}
