package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.List;

public class EchoGenericResponse<T> implements Serializable {
    private int respField;
    private List<T> list;

    public EchoGenericResponse(int respField, List<T> list) {
        this.respField = respField;
        this.list = list;
    }

    public int getRespField() {
        return respField;
    }

    public List<T> getList() {
        return list;
    }
}
