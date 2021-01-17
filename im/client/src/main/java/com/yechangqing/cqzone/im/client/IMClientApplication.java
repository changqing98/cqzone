package com.yechangqing.cqzone.im.client;


import com.yechangqing.cqzone.im.client.client.IMClient;
import java.util.Scanner;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.util.StringUtils;

@SpringBootApplication
public class IMClientApplication implements CommandLineRunner {

  private final IMClient imClient;

  public IMClientApplication(IMClient imClient) {
    this.imClient = imClient;
  }

  public static void main(String[] args) {
    SpringApplication.run(IMClientApplication.class);
  }

  @Override
  public void run(String... args) {
    new Thread(() -> {
      while (true) {
        var in = new Scanner(System.in);
        var str = in.nextLine();
        if (!StringUtils.isEmpty(str)) {
          imClient.sendMessage(str);
        }
      }
    }).start();
  }
}
