package structs.queues;

public class DoubleSideQueue {

    public static void main(String[] args) {
        DoubleSideQueue doubleSideQueue = new DoubleSideQueue();
        doubleSideQueue.appendleft(1);
        doubleSideQueue.appendleft(2);
        System.out.println(doubleSideQueue.pop());
        System.out.println(doubleSideQueue.pop());
    }

    public class Node {
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
        return this.left == null && this.right == null;
    }

    public void append(int value) {
        Node n = new Node(value);

        if (this.right == null) {
            this.left = n;
            this.right = n;
        } else {
            n.prev = this.right;
            this.right.next = n;
            this.right = n;
        }
    }

    public void appendleft(int val) {
        Node n = new Node(val);

        if (this.left == null) {
            this.left = n;
            this.right = n;
        } else {
            n.next = this.left;
            this.left.prev = n;
            this.left = n;
        }
    }

    public int pop() {
        if (this.right == null) {
            return -1;
        }

        Node curr = this.right;
        Node before = this.right.prev;

        if (before == null) {
            this.left = null;
            this.right = null;

            return curr.val;
        }

        before.next = null;
        curr.prev = null;
        this.right = before;

        return curr.val;
    }

    public int popleft() {
        if (this.left == null) {
            return -1;
        }

        Node curr = this.left;
        Node after = this.left.next;

        if (after == null) {
            this.left = null;
            this.right = null;

            return curr.val;
        }

        after.prev = null;
        curr.next = null;
        this.left = after;

        return curr.val;
    }
}
