package org.apache.dubbo.tests.api;

import java.io.Serializable;
import java.util.Objects;

public class EchoGenericEmbedded implements Serializable {
    private String internalStringReq;
    private Integer internalIntegerReq;

    public EchoGenericEmbedded(String stringReq, Integer integerReq) {
        this.internalStringReq = stringReq;
        this.internalIntegerReq = integerReq;
    }

    public String getInternalStringReq() {
        return this.internalStringReq;
    }

    public Integer getInternalIntegerReq() {
        return this.internalIntegerReq;
    }

    @Override
    public boolean equals(Object o) {
        // Self check
        if (this == o) {
            return true;
        }

        // Null check and class check
        if (o == null || getClass() != o.getClass()) {
            return false;
        }

        // Field comparison
        EchoGenericEmbedded other = (EchoGenericEmbedded) o;
        return Objects.equals(internalStringReq, other.internalStringReq) &&
                Objects.equals(internalIntegerReq, other.internalIntegerReq);
    }

    @Override
    public int hashCode() {
        // Hash code should be computed based on the same fields as those used in equals.
        return Objects.hash(internalStringReq, internalIntegerReq);
    }
}
