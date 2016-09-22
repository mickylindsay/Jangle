public class ParseTest2 {
    public static void reader(String test, int loc) {
        String message = "0001";
        String command = "0002";
        String media = "0003";
        String finder = "";
        for (int i = loc; i < test.length(); i++) {
            finder += test.charAt(i);
            if(finder.indexOf(message) != -1) {
                System.out.println(message);
                reader(test, i);
            }
            if(finder.indexOf(command) != -1) {
                System.out.println(message);
                reader(test, i);
            }
            if(finder.indexOf(media) != -1) {
                System.out.println(message);
                reader(test, i);
            }
        }
    }
    public static void main(String[] args) {
        String finder = "0001I can hear you 0002-kick user1 0003song1 is playing";
        reader(finder, 0);
    }
}
