Install some libraries first:

* go get github.com/op/go-logging # Log API
* go get gopkg.in/mgo.v2 # Mongodb API
* go get github.com/json-iterator/go # jsoniter encode, best json performace tool
* go get github.com/stretchr/testify/require # run test

This code implements a reward smart contract for a payment channel protocol.
We use implement the 2of2 ringsignature for Monero using the moneroutil library.
Also, the ZKP and ZKRP is implement by the 