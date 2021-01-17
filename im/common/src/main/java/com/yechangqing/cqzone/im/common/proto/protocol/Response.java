package com.yechangqing.cqzone.im.common.proto.protocol;

import lombok.Data;

@Data
public class Response {
  private Long responseId;
  private String msg;
  private Integer type;
}
