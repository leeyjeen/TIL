import java.util.*;
import java.util.stream.Stream;
import java.util.stream.Collectors;

class StreamToArrayToSet {
    public static void main(String[] args) {
        // create a Stream of Strings
        Stream<String> stream = Stream.of("A", "B", "C", "D");

        // convert Stream into an Array
        String[] arr = stream.toArray(String[] :: new);

        // create a HashSet
        Set<String> set = new HashSet<>();

        // convert Array to set
        Collections.addAll(set, arr);

        // display the elements
        set.forEach(str->System.out.println(str));
    }
}