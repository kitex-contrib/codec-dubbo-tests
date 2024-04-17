package org.apache.dubbo.tests.enumeration;

import java.io.Serializable;

public enum KitexEnum implements Serializable {

    ONE("1"),TWO("2"),THREE("3");

    final String codeStr ;

    KitexEnum(String number) {
        this.codeStr = number;
    }

    // 枚举类型的 getter 方法
    public String getCode() {
        return this.codeStr;
    }
}
