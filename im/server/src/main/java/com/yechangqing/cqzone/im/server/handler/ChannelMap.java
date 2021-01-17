package com.yechangqing.cqzone.im.server.handler;

import io.netty.channel.socket.nio.NioSocketChannel;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ChannelMap {
  private final static Map<Long, NioSocketChannel> CHANNEL_MAP = new ConcurrentHashMap<>(16);

  public static void put(long requestId, NioSocketChannel channel) {
    CHANNEL_MAP.put(requestId, channel);
  }
}
