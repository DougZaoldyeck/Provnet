package main

import (
    //"bytes"
    "encoding/json"
    "fmt"
    //"strconv"
    //"strings"
    //"time"
    //"crypto/rand"
    //"crypto/sha256"
    //"math/big"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    pb "github.com/hyperledger/fabric-protos-go/peer"
)

type Sharing struct{
}

type data struct{
    ObjectType      string `json:"docType"`
    PreviousHash    string `json:"prevhash"`
    Hash            string `json:"hash"`
    Ownership       string `json:"ownership"`
    MinHash         string `json:"minhash"`
    Receiver        string `json:"receiver"`
    TOS             string `json:"TOS"`
    FutureHash      string `json:"futurehash"` //currently treated as the msg2 whn calculating
    Randomness      string `json:"randomness"` //'r' of CH
    SignCH          string `json:"signch"` //represent the 's' of CH (Chameleon Hash)
    PCH             string `json:"pch"`  //'p' of CH
    QCH             string `json:"qch"`  //'q' of CH
    GCH             string `json:"gch"`  //'g' of CH
    HkCH            string `json:"hkch"` //hashkey of CH
}

// Main
func main() {
    err := shim.Start(new(Sharing))
    if err != nil {
        fmt.Printf("Error starting sharing: %s", err)
	}
}

// Init initializes chaincode
func (t *Sharing) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
func (t *Sharing) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "initSharing" { //create a new marble
        return t.initSharing(stub, args)
    } else if function == "readSharing" { //change owner of a specific marble
        return t.readSharing(stub, args)
    } else if function == "updateSharing" {
        return t.updateSharing(stub, args)
    }

    fmt.Println("invoke did not find func: " + function) //error
    return shim.Error("Received unknown function invocation")
}

// initSharing -  create a new sharing, currently adding a new record into the provenance graph
func (t *Sharing) initSharing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error

    // 0-previous hash; 1-current hash; 2-ownership claim; 3-minhash values; 4-receiver;
    // 5-terms of service; 6-future hash; 7-randomness
    if len(args) != 13 {
        return shim.Error("Incorrect number of arguments. Expecting 13")
    }

    // Input sanitation
    fmt.Println("- start init sharing")
    if len(args[0]) <= 0 {
        return shim.Error("1st argument must be a non-empty string")
    }
    if len(args[1]) <= 0 {
        return shim.Error("2nd argument must be a non-empty string")
    }
    if len(args[2]) <= 0 {
        return shim.Error("3rd argument must be a non-empty string")
    }
    if len(args[3]) <= 0 {
        return shim.Error("4th argument must be a non-empty string")
    }
    if len(args[4]) <= 0 {
        return shim.Error("5th argument must be a non-empty string")
    }
    if len(args[5]) <= 0 {
        return shim.Error("6th argument must be a non-empty string")
    }
    if len(args[6]) <= 0 {
        return shim.Error("7th argument must be a non-empty string")
    }
    if len(args[7]) <= 0 {
        return shim.Error("8th argument must be a non-empty string")
    }
    if len(args[8]) <= 0 {
        return shim.Error("9th argument must be a non-empty string")
    }
    if len(args[9]) <= 0 {
        return shim.Error("10th argument must be a non-empty string")
    }
    if len(args[10]) <= 0 {
        return shim.Error("11th argument must be a non-empty string")
    }
    if len(args[11]) <= 0 {
        return shim.Error("12th argument must be a non-empty string")
    }
    if len(args[12]) <= 0 {
        return shim.Error("13th argument must be a non-empty string")
    }
    prevhash := args[0]
    hash := args[1]
    ownership := args[2]
    minhash := args[3]
    receiver := args[4]
    tos := args[5]
    futurehash := args[6]
    randomness := args[7]
    signch := args[8]
    pch := args[9]
    qch := args[10]
    gch := args[11]
    hkch := args[12]

    //check if shared data already exists
    minhashVerify, err := stub.GetState(minhash)   //note we are using minhash values as the main key for each sharing
    if err != nil {
        return shim.Error("Failed to get data: " + err.Error())
    } else if minhashVerify != nil {
        fmt.Println("This data already exists: " + minhash)
        return shim.Error("This marble already exists: " + minhash)
    }

    // Create data object and marshal to JSON
    objectType := "data"
    data := &data{objectType, prevhash, hash, ownership,minhash,
    receiver,tos,futurehash,randomness,signch,pch,qch,gch,hkch}
    dataJSONasBytes, err := json.Marshal(data)
    if err != nil {
        return shim.Error(err.Error())
    }

    // Save data to state
    err = stub.PutState(minhash, dataJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    // Data saved, return success
    fmt.Println("- end init data")
    return shim.Success(nil)
}

// readSharing - read a sharing from chaincode state
func (t *Sharing) readSharing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var minhash, jsonResp string
    var err error

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting name of the sharing to query")
    }

    minhash = args[0]
    valAsbytes, err := stub.GetState(minhash) //get the sharing from chaincode state
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + minhash + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"Error\":\"Data does not exist: " + minhash + "\"}"
        return shim.Error(jsonResp)
    }

    return shim.Success(valAsbytes)
}

// chameleonhash - for updating a new future hash and randomness
func (t *Sharing) updateSharing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    //var futurehash, randomness string
    //var err error

    if len(args) < 3 {
        return shim.Error("Incorrect number of arguments. Expecting name of the sharing to query")
    }

    minhash := args[0]
    futurehash := args[1]
    randomness := args[2]

    fmt.Println("- start updateSharing ", minhash, futurehash, randomness)

    sharingAsBytes, err := stub.GetState(minhash)
    if err != nil {
        return shim.Error("Failed to get sharing:" + err.Error())
    } else if sharingAsBytes == nil {
        return shim.Error("Sharing does not exist")
    }

    sharingToUpdate :=data{}
    err = json.Unmarshal(sharingAsBytes, &sharingToUpdate)
    if err != nil {
        return shim.Error(err.Error())
    }
    sharingToUpdate.FutureHash = futurehash // change the futurehash
    sharingToUpdate.Randomness = randomness //change the randomness

    sharingJSONasBytes, _ := json.Marshal(sharingToUpdate)
    err = stub.PutState(minhash, sharingJSONasBytes) //rewrite the sharing
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Println("- end updateSharing (success)")
    return shim.Success(nil)
}


// chameleonhash - calculating the chameleon hash on demand
/*
func chameleonHash(
    hk *[]byte,
    p *[]byte,
    q *[]byte,
    g *[]byte,
    message *[]byte,
    r *[]byte,
    s *[]byte,
    hashOut *[]byte,
) {
    hkeBig := new(big.Int)
    gsBig := new(big.Int)
    tmpBig := new(big.Int)
    eBig := new(big.Int)
    pBig := new(big.Int)
    qBig := new(big.Int)
    gBig := new(big.Int)
    rBig := new(big.Int)
    sBig := new(big.Int)
    hkBig := new(big.Int)
    hBig := new(big.Int)

    // Converting from hex to bigInt
    pBig.SetString(string(*p), 16)
    qBig.SetString(string(*q), 16)
    gBig.SetString(string(*g), 16)
    hkBig.SetString(string(*hk), 16)
    rBig.SetString(string(*r), 16)
    sBig.SetString(string(*s), 16)

    // Generate the hashOut with message || rBig
    hash := sha256.New()
    hash.Write([]byte(*message))
    hash.Write([]byte(fmt.Sprintf("%x", rBig)))

    eBig.SetBytes(hash.Sum(nil))

    hkeBig.Exp(hkBig, eBig, pBig)
    gsBig.Exp(gBig, sBig, pBig)
    tmpBig.Mul(hkeBig, gsBig)
    tmpBig.Mod(tmpBig, pBig)
    hBig.Sub(rBig, tmpBig)
    hBig.Mod(hBig, qBig)

    *hashOut = hBig.Bytes() // Return hBig in big endian encoding as string
}
*/
