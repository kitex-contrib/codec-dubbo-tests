package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;

public class EchoGenericRequest<T> implements Serializable {
    private int reqField;
    private List<T> list;

    public EchoGenericRequest(int reqField, List<T> list) {
        this.reqField = reqField;
        this.list = list;
    }

    public int getReqField() {
        return reqField;
    }

    public List<T> getList() {
        return list;
    }
}
