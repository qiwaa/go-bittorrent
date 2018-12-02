package txmanager

type TxManager interface {
	Get(key string) (*TxInfo, error)

	Set(key string, val *TxInfo)

	Del(key string)
}
