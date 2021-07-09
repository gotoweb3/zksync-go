package zksync

import "math/big"

type TransactionTypeChangePubKey struct {
	ChangePubKey string `json:"ChangePubKey"`
}

const (
	TransactionTypeChangePubKeyOnchain TransactionType = "Onchain"
	TransactionTypeChangePubKeyECDSA   TransactionType = "ECDSA"
	TransactionTypeChangePubKeyCREATE2 TransactionType = "CREATE2"
)

type TransactionFeeDetails struct {
	GasTxAmount string `json:"gasTxAmount"`
	GasPriceWei string `json:"gasPriceWei"`
	GasFee      string `json:"gasFee"`
	ZkpFee      string `json:"zkpFee"`
	TotalFee    string `json:"totalFee"`
}

func (d *TransactionFeeDetails) GetTotalFee() *big.Int {
	n := new(big.Int)
	if n, ok := n.SetString(d.TotalFee, 10); ok {
		return n
	}
	return new(big.Int)

}

func (d *TransactionFeeDetails) GetTxFee(feeToken *Token) *TransactionFee {
	return &TransactionFee{
		FeeToken: feeToken.Address,
		Fee:      d.GetTotalFee(),
	}
}

type TransactionFee struct {
	FeeToken string   `json:"feeToken"`
	Fee      *big.Int `json:"fee"`
}
