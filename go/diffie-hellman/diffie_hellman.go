package diffiehellman

import (
	"math/big"
	rand "math/rand"
	"time"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

var r *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}

func PrivateKey(p *big.Int) *big.Int {
	k := big.NewInt(0).Sub(p, big.NewInt(2))
	k.Rand(r, k)
	return k.Add(k, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(1).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	prv := PrivateKey(p)
	return prv, PublicKey(prv, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(1).Exp(public2, private1, p)
}
