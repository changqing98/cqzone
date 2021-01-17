package com.yechangqing.cqzone.im.server;

import com.yechangqing.cqzone.im.common.proto.CQZoneIMProto;
import com.yechangqing.cqzone.im.common.proto.protocol.MsgType;

public class HeartBeat {
  private final static CQZoneIMProto.Request bt =
    CQZoneIMProto.Request.newBuilder().setRequestId(0L).setMsg("pong").setType(MsgType.PING)
      .build();

  public static CQZoneIMProto.Request getInstance() {
    return bt;
  }

}

