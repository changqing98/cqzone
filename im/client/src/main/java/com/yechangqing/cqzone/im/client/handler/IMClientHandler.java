package com.yechangqing.cqzone.im.client.handler;

import com.yechangqing.cqzone.im.common.proto.CQZoneIMProto;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class IMClientHandler extends SimpleChannelInboundHandler<CQZoneIMProto.Response> {
  @Override
  protected void channelRead0(ChannelHandlerContext ctx, CQZoneIMProto.Response msg)
    throws Exception {
    log.info("Received response: {}", msg.getMsg());
  }
}
