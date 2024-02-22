package extensions

import gxbig "github.com/dubbogo/gost/math/big"

type BigDecimal = gxbig.Decimal

func NewBigDecimal(str string) (*BigDecimal, error) {
	return gxbig.NewDecFromString(str)
}

type BigInteger = gxbig.Integer

func NewBigInteger(str string) (*BigInteger, error) {
	bigInt := new(BigInteger)
	if err := bigInt.FromString(str); err != nil {
		return nil, err
	}

	return bigInt, nil
}
