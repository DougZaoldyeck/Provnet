#!/bin/bash
#accounts=$(< data.json)

#printf -v json '{"Args":["initMarble",'$accounts',"blue","35","tom"]}'
#peer chaincode invoke -n mycc -C myc -c "$json"
type=$(head  -c 3200 /dev/urandom | od -A n -t x | tr  -ds ' ' '\n' )
type=$(echo $type | tr -d '\n')
type=$(echo $type | tr -d ' ')
echo "$type"

peer chaincode invoke -n provnet -c '{"Args":["initSharing","'$type'","'$type'","owner","'${type}'","Changhao","Nakamoto","'$type'","randomness000","ca68a408f1e827b76127454fd66f21e","e7b6468a8d5e66193629bd3054af2d87","73db234546af330c9b14de982a5796c3","21ec23beab59a5babf080015de73f42c","a23b01bd6784dcb249f397d52c538585","39e048cb483017d29bfcd6b66244e7d8"]}' -C myc
