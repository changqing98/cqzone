package com.yechangqing.cqzone.im.handler;

import com.yechangqing.cqzone.im.protocol.ImRequest;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;

public class ImServerHandler extends SimpleChannelInboundHandler<ImRequest> {
  @Override
  protected void channelRead0(ChannelHandlerContext ctx, ImRequest msg) throws Exception {
    
  }
}
