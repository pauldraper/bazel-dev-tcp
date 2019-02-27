import java.io.*;
import java.net.*;

class EchoServer {
    public static void main(String[] args) throws IOException {
        int port = Integer.parseInt(System.getProperty("echo-server.port"));
        ServerSocket server = new ServerSocket(port);
        System.out.println("Listening on " + port);
        while (true) {
            Socket client = server.accept();
            BufferedReader reader = new BufferedReader(new InputStreamReader(client.getInputStream()));
            PrintWriter writer = new PrintWriter(client.getOutputStream(), true);
            for (String line = null; (line = reader.readLine()) != null; ) {
                writer.println(line);
            }
            client.close();
        }
    }
}
