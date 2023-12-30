package org.apache.dubbo.tests.api;

public class EchoCustomizedException extends Exception {
    private final String customizedMessage;
    public EchoCustomizedException(String customizedMessage) {
        super();
        this.customizedMessage = customizedMessage;
    }

    public String getCustomizedMessage() {
        return this.customizedMessage;
    }
}
