package main

import (
//"fmt"
"testing"

//"github.com/hyperledger/fabric-chaincode-go/shim"
"github.com/hyperledger/fabric-chaincode-go/shimtest"
)


func BenchmarkFunc(b *testing.B) {
	cc := new(Sharing)
	stub := shimtest.NewMockStub("provnet", cc)
	for i :=0 ; i< b.N; i++ {
		stub.MockInvoke("1", [][]byte{[]byte("initSharing"), []byte("prevhash111"), []byte("current_block_hash"), []byte("owner"), []byte("testminhashvalue"), []byte("Changhao"), []byte("Nakamoto"), []byte("futurehash333"), []byte("randomness000"), []byte("ca68a408f1e827b76127454fd66f21e"), []byte("e7b6468a8d5e66193629bd3054af2d87"), []byte("73db234546af330c9b14de982a5796c3"), []byte("21ec23beab59a5babf080015de73f42c"), []byte("a23b01bd6784dcb249f397d52c538585"), []byte("39e048cb483017d29bfcd6b66244e7d8")})
		stub.MockInvoke("1", [][]byte{[]byte("readSharing"), []byte("testminhashvalue")})
		stub.MockInvoke("1", [][]byte{[]byte("updateSharing"), []byte("testminhashvalue"), []byte("NO"), []byte("4a1f93bb46c194381e0f54feeceec4f9"), []byte("423b7ca687efc3ee286aae75a5b0363c")})
	}
}