public class ParseTest {
    public static String reader(String test) {
        String message = "0001";
        String command = "0002";
        String media = "0003";
        if(test.indexOf(message) != -1) {
             return message;
        }
        if(test.indexOf(command) != -1) {
            return command;
        }
        if(test.indexOf(media) != -1) {
            return media;
        }
        return null;
    }
    public static void main(String[] args) {
        String finder = "0001I can hear you";
        String result = reader(finder);
        System.out.println(result);
    }
}
