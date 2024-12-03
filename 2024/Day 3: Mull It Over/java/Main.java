import java.util.*;
import java.util.regex.*;
import java.util.stream.Collectors;

public class Main {

    public static void main(String[] args) {
        try {
            String memory = readMemory();
            calculateOnlyActives(memory);
            calculateAll(memory);
        } catch (RuntimeException e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
        }
    }

    private static void calculateOnlyActives(String memory) {
        var result = getActiveCommands(memory).stream()
                .map(Main::getMulCommands)
                .map(Main::calculate)
                .filter(Optional::isPresent)
                .map(Optional::get)
                .reduce(Integer::sum);
        System.out.println("active calculation: " + result.orElse(0));
    }

    private static void calculateAll(String memory) {
        List<String> mulInput = getMulCommands(memory);
        System.out.println("all calculation: " + calculate(mulInput).orElse(0));
    }

    private static Optional<Integer> calculate(List<String> mulInput) {
        return mulInput.stream().map(Main::mul).reduce(Integer::sum);
    }

    private static List<String> getMulCommands(String memory) {
        String[] startingInput = memory.split("mul\\(");
        return Arrays.stream(startingInput)
                .filter(start -> start.contains(")"))
                .map(start -> start.split("\\)"))
                .filter(endInput -> Pattern.matches("^\\d*,\\d*$", endInput[0]))
                .map(endInput -> endInput[0])
                .toList();
    }

    private static List<String> getActiveCommands(String memory) {
        String[] startDoes = memory.split("do\\(\\)");
        return Arrays.stream(startDoes)
                .map(startDo -> startDo.split("don't")[0])
                .collect(Collectors.toList());
    }

    private static int mul(String input)  {
        String[] inputSplit = input.split(",");
        if (inputSplit.length != 2) {
            throw new IllegalArgumentException("Invalid mul input: " + input);
        }
        int a = Integer.parseInt(inputSplit[0]);
        int b = Integer.parseInt(inputSplit[1]);
        return a * b;
    }

    private static String readMemory() {
        StringBuilder memory = new StringBuilder();
        try (Scanner scanner = new Scanner(System.in)) {
            System.out.println("Enter input (press Ctrl+D or Ctrl+Z to end):");
            while (scanner.hasNextLine()) {
                memory.append(scanner.nextLine());
            }
        }
        return memory.toString();
    }
}
