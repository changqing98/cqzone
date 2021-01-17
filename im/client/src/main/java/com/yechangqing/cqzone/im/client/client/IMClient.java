package com.yechangqing.cqzone.im.client.client;

import com.yechangqing.cqzone.im.client.handler.IMClientHandler;
import com.yechangqing.cqzone.im.common.proto.CQZoneIMProto;
import com.yechangqing.cqzone.im.common.proto.protocol.MsgType;
import io.netty.bootstrap.Bootstrap;
import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioSocketChannel;
import javax.annotation.PostConstruct;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class IMClient {

  private final NioEventLoopGroup group = new NioEventLoopGroup(0);

  private NioSocketChannel channel;

  @PostConstruct
  private void start() throws InterruptedException {
    Bootstrap bootstrap = new Bootstrap();
    bootstrap.group(group).channel(NioSocketChannel.class).handler(new IMClientHandler());
    var future = bootstrap.connect("localhost", 8888).sync();
    channel = (NioSocketChannel) future.channel();
  }

  public void sendMessage(String msg) {
    ByteBuf byteBuf = Unpooled.buffer(msg.getBytes().length);
    byteBuf.writeBytes(msg.getBytes());
    var future = channel.writeAndFlush(byteBuf);
    future.addListener((ChannelFutureListener) channelFuture -> log.info("客户端手动发消息成功={}", msg));
  }

  public void sendProtoMessage(String msg) {
    var req =
      CQZoneIMProto.Request.newBuilder().setMsg(msg).setRequestId(100L).setType(MsgType.MSG);
    var future = channel.writeAndFlush(req);
    future.addListener((ChannelFutureListener) channelFuture -> log.info("客户端手动发消息成功={}", msg));
  }
}
