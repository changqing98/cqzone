package com.yechangqing.cqzone.im.common.proto.protocol;

import lombok.Data;

@Data
public class Request {
  private Long requestId;
  private String msg;
  private Integer type;
}
