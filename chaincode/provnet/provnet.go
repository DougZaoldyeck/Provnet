package main


// #include <stdlib.h>
// #include <string.h>
/*
int minHashCmp(const char *a, const char *b, const int minhashVolum){
 // change minHash settings here
int minhashLen = 64;

 // count matches and return
    int match = 0;
    int i;
    for (i = 0; i < minhashVolum; i ++) {
        if (memcmp(a+i*minhashLen, b+i*minhashLen, minhashLen) == 0) {
            match ++;
        }
 }
    return match;
}*/
import "C"

import (
    "bytes"
    "encoding/json"
    "fmt"
    //"strconv"
    "strings"
    //"time"
    //"crypto/rand"
    "crypto/sha256"
    "math/big"
    "time"
    "unsafe"

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
    MinHash1        string `json:"minhash1"`
    MinHash2        string `json:"minhash2"`
    MinHash3        string `json:"minhash3"`
    MinHash4        string `json:"minhash4"`
    MinHash5        string `json:"minhash5"`
    MinHash6        string `json:"minhash6"`
    MinHash7        string `json:"minhash7"`
    Receiver        string `json:"receiver"`
    TOS             string `json:"TOS"`
    FutureHash      string `json:"futurehash"` //currently treated as the msg2 whn calculating
    Randomness      string `json:"randomness"` //'r' of CH
    SignCH          string `json:"signch"` //represent the 's' of CH (Chameleon Hash)
    PCH             string `json:"pch"`  //'p' of CH
    QCH             string `json:"qch"`  //'q' of CH
    GCH             string `json:"gch"`  //'g' of CH
    HkCH            string `json:"hkch"` //hashkey of CH
    ChameleonHash   string `json:"chameleonhash"` //term for chameleon hash
    TraverseHelper  string `json:"traversehelper"`
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
    //fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "initSharing" { //create a new data
        return t.initSharing(stub, args)
    } else if function == "readSharing" { //access a data
        return t.readSharing(stub, args)
    } else if function == "updateSharing" {
        return t.updateSharing(stub, args)
    } else if function == "queryDataByOwner" {
        return t.queryDataByOwner(stub, args)
    }

    fmt.Println("invoke did not find func: " + function) //error
    return shim.Error("Received unknown function invocation")
}

// initSharing -  create a new sharing, currently adding a new record into the provenance graph
func (t *Sharing) initSharing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    //start := time.Now()
    var err error

    // 0-previous hash; 1-current hash; 2-ownership claim; 3-minhash values; 4~10-complementary minhash values;
    // 11-receiver; 12-terms of service; 13-future hash; 14-randomness; 15-SignCH;
    // 16~19-PCH/QCH/GCH/HkCH; 20-chameleon hash
    if len(args) != 22 {
        return shim.Error("Incorrect number of arguments. Expecting 22")
    }

    // Input sanitation
    //fmt.Println("- start init sharing")
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
    if len(args[13]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[14]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[15]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[16]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[17]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[18]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[19]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[20]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[21]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    prevhash := args[0]
    hash := args[1]
    ownership := strings.ToLower(args[2])
    minhash := args[3]
    minhash1 := args[4]
    minhash2 := args[5]
    minhash3 := args[6]
    minhash4 := args[7]
    minhash5 := args[8]
    minhash6 := args[9]
    minhash7 := args[10]
    receiver := strings.ToLower(args[11])
    tos := strings.ToLower(args[12])
    futurehash := args[13]
    randomness := args[14]
    signch := args[15]
    pch := args[16]
    qch := args[17]
    gch := args[18]
    hkch := args[19]
    chhash := args[20]
    traversehelper := strings.ToLower(args[21])



    // =========================================================================================
    // check if similar data exists.
    // but this is just a half-complete part since we only get the similarities but no rejection
    // or warning if similar records found.
    // enough for a check-of-performance. currently disabled for check of other functions.
    // =========================================================================================

    /*
    if ownership=="owner" {
        queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"data\",\"ownership\":\"%s\"}}", ownership)

        queryResults, err := getQueryResultForQueryString(stub, queryString, minhash)
        if err != nil {
            return shim.Error(err.Error())
        }
        //fmt.Sprintf("%s",queryResults)
        if len(queryResults) != 0 {
            fmt.Sprintf("\n\n\nSome existing records: %s\n\n\n",queryResults)
        }


        temp := "sender"
        queryString = fmt.Sprintf("{\"selector\":{\"docType\":\"data\",\"ownership\":\"%s\"}}", temp)

        queryResults, err = getQueryResultForQueryString(stub, queryString, minhash)
        if err != nil {
            return shim.Error(err.Error())
        }
        //if len(queryResults) != 0 {
            fmt.Sprintf("\n\n\nSome existing records: %s\n\n\n",queryResults)
        //}

    }*/


    // =========================================================================================
    // =========================================================================================
    // check if shared data already exists.
    // need to be further modified by adding a the check of minhash values
    // changed the order here between finding the same minhash values and the similar ones
    // just for experiment purposes.
    // =========================================================================================
    // =========================================================================================

    minhashVerify, err := stub.GetState(minhash)   //note we are using minhash values as the main key for each sharing
    if err != nil {
        return shim.Error("Failed to get data: " + err.Error())
    } else if minhashVerify != nil {
        //fmt.Println("This data already exists: ")
        return shim.Error("This marble already exists: " + minhash)
    }


    // Create data object and marshal to JSON
    objectType := "data"
    data := &data{objectType, prevhash, hash, ownership,minhash, minhash1, minhash2,
        minhash3, minhash4, minhash5, minhash6,minhash7,receiver,tos,futurehash,
        randomness,signch,pch,qch,gch,hkch, chhash, traversehelper}
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
    //fmt.Println("- end init data")

    //fmt.Println("Inserting data...")

    return shim.Success(nil)
}

// readSharing - read a sharing from chaincode state
func (t *Sharing) readSharing(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    var minhash, jsonResp string
    var err error

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting name of the sharing to query")
    }

    start := time.Now()
    minhash = args[0]
    valAsbytes, err := stub.GetState(minhash) //get the sharing from chaincode state
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + minhash + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"Error\":\"Data does not exist: " + minhash + "\"}"
        return shim.Error(jsonResp)
    }

    elapsed := time.Since(start)
    fmt.Println("time for reading:",elapsed)

    return shim.Success(valAsbytes)
}

// =========================================================================================
// update sharing - for updating a new future hash, randomness and new_s
// =========================================================================================

func (t *Sharing) updateSharing(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 4 {
        return shim.Error("Incorrect number of arguments. Expecting 4")
    }

    //start := time.Now()

    minhash := args[0]
    futurehash := args[1]
    randomness := args[2]
    newsign := args[3]

    var newr []byte = []byte(randomness)
    var newmsg []byte = []byte(futurehash)
    var news []byte = []byte(newsign)
    var hash1 []byte

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


    var hashkey []byte = []byte(sharingToUpdate.HkCH)
    var pkey []byte = []byte(sharingToUpdate.PCH)
    var qkey []byte = []byte(sharingToUpdate.QCH)
    var gkey []byte = []byte(sharingToUpdate.GCH)
    var hash0  = sharingToUpdate.ChameleonHash

    chameleonHash(&hashkey, &pkey, &qkey, &gkey, &newmsg, &newr, &news, &hash1)

    temp := string(([]byte(fmt.Sprintf("%x", hash1))))

    if hash0 == temp {
        sharingToUpdate.FutureHash = futurehash
        sharingToUpdate.Randomness = randomness
        sharingToUpdate.SignCH = newsign
        sharingJSONasBytes, _ := json.Marshal(sharingToUpdate)
        err = stub.PutState(minhash, sharingJSONasBytes) //rewrite the sharing
        if err != nil {
            return shim.Error(err.Error())
        }

        fmt.Println("- end updateSharing (success)")
        //elapsed := time.Since(start)
        //fmt.Println("time to update:",elapsed)

        return shim.Success(nil)
    } else {
        return shim.Error("Wrong randomness provided")
    }


}

// =========================================================================================
// queryDataByOwner queries for data based on a passed in owner.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g., CouchDB)
// =========================================================================================
func (t *Sharing) queryDataByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    //   0          1
    // "owner"  "minhash"
    if len(args) < 9 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    traversehelper := strings.ToLower(args[0])
    minhash := args[1]
    minhash1 := args[2]
    minhash2 := args[3]
    minhash3 := args[4]
    minhash4 := args[5]
    minhash5 := args[6]
    minhash6 := args[7]
    minhash7 := args[8]


    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"data\",\"traversehelper\":\"%s\"}}", traversehelper)

    queryResults, err := getQueryResultForQueryString(stub, queryString, minhash, minhash1, minhash2, minhash3, minhash4, minhash5, minhash6, minhash7)
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(queryResults)
}


// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string, minhash string, minhash1 string, minhash2 string, minhash3 string, minhash4 string,
    minhash5 string, minhash6 string, minhash7 string) ([]byte, error) {

    //fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

    resultsIterator, err := stub.GetQueryResult(queryString)
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    buffer, err := constructQueryResponseFromIterator(resultsIterator, stub, minhash, minhash1, minhash2, minhash3, minhash4, minhash5, minhash6, minhash7)
    if err != nil {
        return nil, err
    }

    //fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

    return buffer.Bytes(), nil
}



// ===========================================================================================
// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
// ===========================================================================================
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface, stub shim.ChaincodeStubInterface, minhash string, minhash1 string, minhash2 string, minhash3 string, minhash4 string,
    minhash5 string, minhash6 string, minhash7 string) (*bytes.Buffer, error) {
    // buffer is a JSON array containing QueryResults
    var buffer bytes.Buffer
    buffer.WriteString("[")

    minhashVolum := C.int(50)


    //bArrayMemberAlreadyWritten := false
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }
        // Add a comma before array members, suppress it for the first array member
        /*if bArrayMemberAlreadyWritten == true {
            buffer.WriteString(",")
        }*/
        //====currently unabled since we dont need return the exact key values


        /*
        buffer.WriteString("{\"Key\":")
        buffer.WriteString("\"")
        buffer.WriteString(queryResponse.Key)
        buffer.WriteString("\"")*/

        a1 := C.CString(minhash)
        b1 := C.CString(queryResponse.Key)

        match := C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))


        resultsAsBytes, err := stub.GetState(queryResponse.Key)
        if err != nil {
            fmt.Println("nothing got from minhash\n")
            return  &buffer, nil
        } else if resultsAsBytes == nil {
            fmt.Println("nothing got from results\n")
            return  &buffer, nil
        }

        resultsToCompare :=data{}
        err = json.Unmarshal(resultsAsBytes, &resultsToCompare)
        if err != nil {
            fmt.Println("nothing got from comparison\n")
            return  &buffer, nil
        }

        a1 = C.CString(minhash1)
        b1 = C.CString(resultsToCompare.MinHash1)
        match = match + C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))
        a1 = C.CString(minhash2)
        b1 = C.CString(resultsToCompare.MinHash2)
        match = match + C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))
        a1 = C.CString(minhash3)
        b1 = C.CString(resultsToCompare.MinHash3)
        match = match + C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))
        a1 = C.CString(minhash4)
        b1 = C.CString(resultsToCompare.MinHash4)
        match = match + C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))
        a1 = C.CString(minhash5)
        b1 = C.CString(resultsToCompare.MinHash5)
        match = match + C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))
        a1 = C.CString(minhash6)
        b1 = C.CString(resultsToCompare.MinHash6)
        match = match + C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))
        a1 = C.CString(minhash7)
        b1 = C.CString(resultsToCompare.MinHash7)
        match = match + C.minHashCmp(a1, b1, minhashVolum)
        C.free(unsafe.Pointer(a1))
        C.free(unsafe.Pointer(b1))
        //fmt.Println(match)

        queryResponse = nil

        //bArrayMemberAlreadyWritten = true
    }

    //buffer.WriteString("]")

    return &buffer, nil
}


// =========================================================================================
// chameleonhash - calculating the chameleon hash on demand based on the provided parameters
// including the newmessage (considered as the future hash item here), newrandomness, new_s
// =========================================================================================

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






