package com.yechangqing.cqzone.im.server;

import com.yechangqing.cqzone.im.server.config.ImConfig;
import com.yechangqing.cqzone.im.server.handler.ImServerInitializer;
import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import java.net.InetSocketAddress;
import javax.annotation.PostConstruct;
import javax.annotation.PreDestroy;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

@Component
@Slf4j
public class IMServer {

  private final ImConfig imConfig;

  private final EventLoopGroup boss = new NioEventLoopGroup();
  private final EventLoopGroup worker = new NioEventLoopGroup();

  public IMServer(ImConfig imConfig) {
    this.imConfig = imConfig;
  }

  @PostConstruct
  public void start() throws InterruptedException {
    ServerBootstrap serverBootstrap = new ServerBootstrap();
    serverBootstrap.group(boss, worker).channel(NioServerSocketChannel.class)
      .localAddress(new InetSocketAddress(imConfig.getPort()))
      .childOption(ChannelOption.SO_KEEPALIVE, true).childHandler(new ImServerInitializer());
    var future = serverBootstrap.bind().sync();
    if (future.isSuccess()) {
      log.info("Start im server success!");
    }
  }

  @PreDestroy
  public void destroy() {
    boss.shutdownGracefully().syncUninterruptibly();
    worker.shutdownGracefully().syncUninterruptibly();
    log.info("Close im server success!");
  }
}
