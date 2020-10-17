import java.io.*;
public class EmployeeTest {
    public static void main(String args[]) {
        // Create two objects using constructor
        Employee empOne = new Employee("Claudia");
        Employee empTwo = new Employee("Jimmy");

        // Invoking methods for each object created
        empOne.empAge(26);
        empOne.empDesignation("Actress");
        empOne.empSalary(1000);
        empOne.printEmployee();

        empTwo.empAge(30);
        empTwo.empDesignation("Actor");
        empTwo.empSalary(800);
        empTwo.printEmployee();
    }
}