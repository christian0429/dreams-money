package bitcoin


type Dogecoin struct{}

func (Dogecoin) ChainName() string {
	return "dogecoin"
}

func (Dogecoin) CoinbaseDigest(coinbase string) (string, error) {
	return DoubleSha256(coinbase)
}

func (Dogecoin) HeaderDigest(header string) (string, error) {
	return ScryptDigest(header)
}

func (Dogecoin) ShareMultiplier() float64 {
	return 65536
}

func (Dogecoin) ValidMainnetAddress(address string) bool {
	return true
}

func (Dogecoin) ValidTestnetAddress(address string) bool {
	return true
}

func (Dogecoin) MinimumConfirmations() uint {
	return uint(251)
}
