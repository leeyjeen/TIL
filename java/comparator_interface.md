# Comparator Interface in Java with Examples
Comparator 인터페이스는 사용자 정의 클래스의 객체를 정렬하는 데 사용된다. comparator 객체는 서로 다른 두 클래스의 두 객체를 비교할 수 있다. 다음 함수는 obj1와 obj2를 비교한다.

```java
public int compare(Object obj1, Object obj2);
```
rollno, name, address, DOB 등의 필드를 포함하는 자체 클래스 유형의 array/arraylist를 가지고 rollno 또는 name 기준으로 배열을 정렬해야 한다고 가정하자.

* method1
    * 한 가지 분명한 접근 방식은 표준 알고리즘 중 하나를 사용하여 자체 sort() 함수를 작성하는 것이다. 이 방법은 Roll No. 및 Name과 같은 다른 기준에 대해 전체 정렬 코드를 재작성해야 한다.
* method2
    * comparator 인터페이스 사용하기: Comparator 인터페이스는 사용자 정의 클래스 객체를 정렬하는 데 사용된다. 이 인터페이스는 java.util 패키지에 있으며 `compare(Object obj1, Object obj2)` 및 `equals(Object element)` 두 가지 메서드를 포함한다.
    * comparator를 사용하여 데이터 멤버를 기반으로 요소를 정렬할 수 있다. 예를 들어, rollno, name, age 또는 다른 무엇이든 될 수 있다.

List 요소를 정렬하기 위한 Collections 클래스의 메서드는 주어진 comparator로 List 요소를 정렬하는 데 사용된다.

```java
// To sort a given list. ComparatorClass must implement
// Comparator interface.
public void sort(List list, ComparatorClass c)
```

## How does Collections.Sort() work?
내부적으로 sort() 메서드는 정렬 중인 클래스의 Compare() 메서드를 호출한다. Compare 메서드는 -1, 0 또는 1을 반환하여 다른 값보다 작거나 같거나 큰지 여부를 나타낸다. 이 결과를 이용하여 정렬을 위해 교환해야 하는지 여부를 결정한다.

```java
import java.util.*;
import java.lang.*;
import java.io;

//  a class to represent a student.
class Student {
    int rollno;
    String name, address;

    // contructor
    public Student(int rollno, String name, String address) {
        this.rollno = rollno;
        this.name = name;
        this.address = address;
    }

    public String toString() {
        return this.rollno + " " + this.name + " " + this.address;
    }
}

class Sortbyroll implements Comparator<Student> {
    // used for sorting in ascending order of roll number
    public int compare(Student a, Student b) {
        return a.rollno - b.rollno;
    }
}

class Sortbyname implements Comparator<Student> {
    // used for sorting in ascending order of roll name
    public int compare(Student a, Student b) {
        return a.name.compareTo(b.name);
    }
}

class Main {
    public static void main(String[] args) {
        ArrayList<Student> ar = new ArrayList<Student>();
        ar.add(new Student(111, "bbbb", "london"));
        ar.add(new Student(131, "aaaa", "nyc"));
        ar.add(new Student(121, "cccc", "jaipur"));

        System.out.println("Unsorted");
        for (int i=0; i<ar.size(); i++) {
            System.out.println(ar.get(i));
        }

        Collections.sort(ar, new Sortbyroll()));

        System.out.println("\nSorted by rollno");
        for (int i=0; i<ar.size(); i++) {
            System.out.println(ar.get(i));
        }

        Collections.sort(ar, new Sortbyname()));

        System.out.println("\nSorted by name");
        for (int i=0; i<ar.size(); i++) {
            System.out.println(ar.get(i));
        }
    }
}
```
Output:
```
Unsorted
111 bbbb london
131 aaaa nyc
121 cccc jaipur

Sorted by rollno
111 bbbb london
121 cccc jaipur
131 aaaa nyc

Sorted by name
131 aaaa nyc
111 bbbb london
121 cccc jaipur
```

compare 메서드 내부 반환값을 변경하여 원하는 순서로 정렬할 수 있다. 예를 들어 내림차순의 경우 위의 compare 메서드에서 a와 b의 위치를 변경하면 된다.

## Sort collection by more than one field
Comparable 및 Comparator 인터페이스를 사용하여 단일 필드를 기준으로 객체 리스트를 정렬하는 방법에 대해 논의하였다. 그러나 두 개 이상의 필드에 따라 예를 들어 학생 이름, 학생 나이에 따라 ArrayList 객체를 정렬해야 하는 경우가 있다.

```java
import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.Comparator;

class Student {
    // instance member variable
    String Name;
    int Age;
    
    // parameterized constructor
    public Student(String Name, Integer Age) {
        this.Name = Name;
        this.Age = Age;
    }

    public String getName() {
        return Name;
    }

    public void SetName(String Name) {
        this.Name = Name;
    }

    public Integer getAge() {
        return Age;
    }

    public void setAge(Integer Age) {
        this.Age = Age;
    }

    // overriding toString() method
    @Override
    public String toString() {
        return "Customer{" + "Name=" + Name + ", Age=" + Age + '}';
    }

    static class CustomerSortingComparator implements Comparator<Student> {
        @Override
        public int compare(Student customer1, Student customer2) {
            // for comparison
            int NameCompare = customer1.getName().compareTo(customer2.getName());
            int AgeCompare = customer1.getAge().compareTo(customer2.getAge());

            // 2-level comparison using if-else block
            if (NameCompare == 0) {
                return ((AgeCompare == 0) ? NameCompare : AgeCompare);
            } else {
                return NameCompare;
            }
        }
    }
    
    public static void main(String[] args) {
        // create ArrayList to stroe Student
        List<Student> al = new ArrayList<>();

        // create customer objects using constructor initialization
        Student obj1 = new Student("Ajay", 27);
        Student obj2 = new Student("Sneha", 23); 
        Student obj3 = new Student("Simran", 37); 
        Student obj4 = new Student("Ajay", 22); 
        Student obj5 = new Student("Ajay", 29); 
        Student obj6 = new Student("Sneha", 22); 

        // add customer objects to ArrayList
        a1.add(obj1);
        a1.add(obj2);
        al.add(obj3); 
        al.add(obj4); 
        al.add(obj5); 
        al.add(obj6);

        // before Sorting arrayList: iterate using Iterator
        Iterator<Student> cusIterator = al.iterator();

        System.out.println("Before Sorting:\n");
        while (cusIterator.hasNext()) {
            System.out.println(cusIterator.next());
        }

        // sorting using Collections.sort(al, comparator);
        Collections.sort(al, new CustomerSortingComparator());

        // after Sorting arrayList: iterate using enhanced for-loop
        System.out.println("\n\nAfter Sorting:\n");
        for(Student customer : al) {
            System.out.println(customer);
        }
    }
}
```
Output:
```
Before Sorting:

Customer{Name=Ajay, Age=27}
Customer{Name=Sneha, Age=23}
Customer{Name=Simran, Age=37}
Customer{Name=Ajay, Age=22}
Customer{Name=Ajay, Age=29}
Customer{Name=Sneha, Age=22}


After Sorting:

Customer{Name=Ajay, Age=22}
Customer{Name=Ajay, Age=27}
Customer{Name=Ajay, Age=29}
Customer{Name=Simran, Age=37}
Customer{Name=Sneha, Age=22}
Customer{Name=Sneha, Age=23}
```

## Reference
* https://www.geeksforgeeks.org/comparator-interface-java/