package main

class LRUCache {

	private static class Node {
			int key, value;
			Node prev, next;

			Node(int k, int v){
					key = k;
					value = v;
			}
	}

	private final int capacity;
	private final Node dummy = new Node(0,0);
	private final Map<Integer,Node> keyNode = new HashMap<>();


	public LRUCache(int capacity) {
			this.capacity = capacity;
			dummy.prev = dummy;
			dummy.next = dummy;        
	}
	
	public int get(int key) {
			Node node = getNode(key);
			if (node == null){
					return -1;
			}
			return node.value;
	}
	
	public void put(int key, int value) {
			Node node = getNode(key);
			if (node != null){
					node.value = value;
					return;
			}
			node = new Node(key, value);
			keyNode.put(key, node);
			pushHead(node);
			if (keyNode.size() > capacity){
					Node deleteNode = dummy.prev;
					keyNode.remove(deleteNode.key);
					remove(deleteNode);

			}
	}

	private Node getNode(int key){
			if (!keyNode.containsKey(key)){
					return null;
			}
			Node node = keyNode.get(key);
			remove(node);
			pushHead(node);
			return node;
	}

	private void remove(Node x){
			x.prev.next = x.next;
			x.next.prev = x.prev;
	}

	private void pushHead(Node x){
			x.prev = dummy;
			x.next = dummy.next;
			x.prev.next = x;
			x.next.prev = x;
	}
}

/**
* Your LRUCache object will be instantiated and called as such:
* LRUCache obj = new LRUCache(capacity);
* int param_1 = obj.get(key);
* obj.put(key,value);
*/