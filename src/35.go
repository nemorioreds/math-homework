public class GoSolver {
    public static void main(String[] args) {
        int[] numbers = {1, 2, 3, 4, 5};
        System.out.println("The maximum number in the array is: " + Arrays.stream(numbers).max().getAsInt());
    }
}
