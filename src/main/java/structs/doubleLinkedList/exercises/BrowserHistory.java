package structs.doubleLinkedList.exercises;

import structs.doubleLinkedList.DoubleLinkedList;

public class BrowserHistory {

    public static void main(String[] args) {
        BrowserHistory history = new BrowserHistory("vinicius");

        System.out.println(history.currentPage.url);
        history.visit("duarte");
        System.out.println(history.currentPage.url);
        history.visit("mendes");
        System.out.println(history.currentPage.url);
        history.visit("nepomuceno");
        System.out.println(history.currentPage.url);

        history.back(3);
        System.out.println(history.currentPage.url);
        history.forward(3);
        System.out.println(history.currentPage.url);

        history.back(5);
        System.out.println(history.currentPage.url);
        history.visit("ama");
        System.out.println(history.currentPage.url);
        history.visit("elizabeth");
        System.out.println(history.currentPage.url);

        history.back(2);
        System.out.println(history.currentPage.url);
        history.forward(10);
        System.out.println(history.currentPage.url);
    }

    public class Node {
        Node next;
        Node prev;
        String url;

        public Node(String url) {
            this.url = url;
        }
    }

    Node currentPage;


    public BrowserHistory(String homepage) {
        Node newNode = new Node(homepage);

        currentPage = newNode;
    }

    public void visit(String url) {
        Node newNode = new Node(url);

        newNode.prev = currentPage;
        currentPage.next = newNode;
        currentPage = newNode;
    }

    public String back(int steps) {
        while (currentPage.prev != null && steps != 0) {
            currentPage = currentPage.prev;
            steps--;
        }

        return currentPage.url;
    }

    public String forward(int steps) {
        while (currentPage.next != null && steps != 0) {
            currentPage = currentPage.next;
            steps--;
        }

        return currentPage.url;
    }
}
