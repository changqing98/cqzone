package com.yechangqing.cqzone.im.server.handler;

import com.yechangqing.cqzone.im.common.proto.CQZoneIMProto;
import com.yechangqing.cqzone.im.common.proto.protocol.MsgType;
import com.yechangqing.cqzone.im.server.HeartBeat;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.channel.socket.nio.NioSocketChannel;
import io.netty.util.AttributeKey;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class ImServerHandler extends SimpleChannelInboundHandler<CQZoneIMProto.Request> {

  private final static AttributeKey<String> READER_TIME = AttributeKey.valueOf("readerTime");

  @Override
  protected void channelRead0(ChannelHandlerContext ctx, CQZoneIMProto.Request msg) {
    log.info("Received msg = [{}]", msg.toString());

    if (msg.getType() == MsgType.LOGIN) {
      ChannelMap.put(msg.getRequestId(), (NioSocketChannel) ctx.channel());
      SessionMap.put(msg.getRequestId(), msg.getMsg());
      log.info("client {} logged", msg.getRequestId());
    }

    if (msg.getType() == MsgType.PING) {
      ctx.channel().attr(READER_TIME).set(String.valueOf(System.currentTimeMillis()));
      var hb = HeartBeat.getInstance();
      ctx.writeAndFlush(hb).addListener((ChannelFutureListener) future -> {
        if (!future.isSuccess()) {
          log.error("IO error, channel is closing");
          future.channel().close();
        }
      });
    }
  }


}
