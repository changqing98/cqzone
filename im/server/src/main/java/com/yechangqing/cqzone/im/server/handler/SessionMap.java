package com.yechangqing.cqzone.im.server.handler;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class SessionMap {
  private final static Map<Long, String> SESSION_MAP = new ConcurrentHashMap<>(16);


  public static void put(Long requestId, String info) {
    SESSION_MAP.put(requestId, info);
  }

}
