package com.yechangqing.cqzone.im.protocol;

import lombok.Data;

@Data
public class ImRequest {
  private Long requestId;
  private String reqMsg;
  private Integer type;
}
