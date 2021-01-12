package com.yechangqing.cqzone.im.protocol;

import lombok.Data;

@Data
public class ImResponse {
  private Long responseId;
  private String resMsg;
  private Integer type;
}
