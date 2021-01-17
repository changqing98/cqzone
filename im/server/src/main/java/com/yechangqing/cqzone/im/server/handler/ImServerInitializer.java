package com.yechangqing.cqzone.im.server.handler;

import com.yechangqing.cqzone.im.common.proto.CQZoneIMProto;
import io.netty.channel.Channel;
import io.netty.channel.ChannelInitializer;
import io.netty.handler.codec.protobuf.ProtobufDecoder;
import io.netty.handler.codec.protobuf.ProtobufEncoder;
import io.netty.handler.codec.protobuf.ProtobufVarint32FrameDecoder;
import io.netty.handler.codec.protobuf.ProtobufVarint32LengthFieldPrepender;
import io.netty.handler.codec.string.StringDecoder;
import io.netty.handler.timeout.IdleStateHandler;

public class ImServerInitializer extends ChannelInitializer<Channel> {
  private final ImServerHandler handler = new ImServerHandler();

  @Override
  protected void initChannel(Channel ch) {
    ch.pipeline().addLast(new IdleStateHandler(1000, 0, 0))
      .addLast(new StringDecoder())
      .addLast(new StringHandler());
  }
}
