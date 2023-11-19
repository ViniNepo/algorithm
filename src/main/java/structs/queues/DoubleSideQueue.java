package structs.queues;

public class DoubleSideQueue {

    class Node {
        int val;
        Node next;
        Node prev;

        public Node(int val) {
            this.val = val;
        }
    }

    Node left;
    Node right;

    public DoubleSideQueue() {
        this.left = null;
        this.right = null;
    }

    public boolean isEmpty() {
        return left == null && right == null;
    }

    public void append(int value) {

    }
}
