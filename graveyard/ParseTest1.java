public class ParseTest1 {
    public static String reader(String test) {
        String message = "0001";
        String command = "0002";
        String media = "0003";
        String finder = "";
        for (int i = 0; i < test.length(); i++) {
            finder += test.charAt(i);
            if(finder.indexOf(message) != -1) {
                return message;
            }
            if(finder.indexOf(command) != -1) {
                return command;
            }
            if(finder.indexOf(media) != -1) {
                return media;
            }
        }
        return null;
    }
    public static void main(String[] args) {
        String finder = "0001I can hear you";
        String result = reader(finder);
        System.out.println(result);
    }
}
