public class Table {
    Slot[] slots = new Slot[100];

    public Table() {
        for(int i = 0;i<100;i++)
        {
            slots[i] = new Slot();
        }
    }

    public static class Person {
        int ssn;
        String name;
        String addr;

        public Person() {
        }

        public Person(int ssn, String name, String addr) {
            this.ssn = ssn;
            this.addr = addr;
            this.name = name;
        }

        public void ShowPerInfo() {
            System.out.println("번호: " + this.ssn);
            System.out.println("이름: " + this.name);
            System.out.println("주소: " + this.addr);
        }
    }

    public static class Slot {
        int key;
        Person value;
        String slotStatus; // EMPTY, DELETED, INUSE
    
        public Slot() {
            this.slotStatus = "EMPTY";
        }
    }

    public static int hashFunc(int k) {
        return k % 100;
    }

    public static void tableInsert(Table table, int key, Person value) {
        int hashValue = hashFunc(key);
        table.slots[hashValue].value = value;
        table.slots[hashValue].key = key;
        table.slots[hashValue].slotStatus = "INUSE";
    }

    public static Person tableDelete(Table table, int key) {
        int hashValue = hashFunc(key);
        if (table.slots[hashValue].slotStatus != "INUSE") {
            return null;
        } else {
            table.slots[hashValue].slotStatus = "DELETED";
            return table.slots[hashValue].value;
        }
    }

    public static Person tableSearch(Table table, int key) {
        int hashValue = hashFunc(key);

        if (table.slots[hashValue].slotStatus != "INUSE") {
            return null;
        }
        return table.slots[hashValue].value;
    }
    public static void main(String[] args) {
        Table myTable = new Table();
        Person lee = new Person(20200501, "lee", "anyang");
        Person goo = new Person(20200602, "goo", "seoul");
        tableInsert(myTable, lee.ssn, lee);
        tableInsert(myTable, goo.ssn, goo);

        Person searchPerson = new Person();
        searchPerson = tableSearch(myTable, 20200501);
        if (searchPerson != null) {
            searchPerson.ShowPerInfo();
        }

        searchPerson = tableSearch(myTable, 20200602);
        if (searchPerson != null) {
            searchPerson.ShowPerInfo();
        }

        tableDelete(myTable, 20200501);
        tableDelete(myTable, 20200602);
    }
}
