#!/bin/bash

for((a = 0; a <= 999; a++))
do
type=$(head  -c 32 /dev/urandom | od -A n -t x | tr  -ds ' ' '\n' )
type=$(echo $type | tr -d '\n')
type=$(echo $type | tr -d ' ')
echo "$type"
    peer chaincode invoke -n mycc -c peer chaincode invoke -n mycc -c '{"Args":["initMarble","'$type'","blue","35","tom"]}' -C myc
    sleep 0.3s
    if [ `expr $a % 100` == 99 ]
    then
        sleep 2s
        for((b = 0; b <= 49; b++))
        do
#The following part is for random generating a name for a marble
#type=$(head  -c 32 /dev/urandom | od -A n -t x | tr  -ds ' ' '\n' )
#type=$(echo $type | tr -d '\n')
#type=$(echo $type | tr -d ' ')
#echo "$type"
#peer chaincode invoke -n mycc -c peer chaincode invoke -n mycc -c '{"Args":["initMarble","'$type'","blue","35","tom"]}' -C myc
#sleep 2s
#peer chaincode invoke -n mycc -c peer chaincode invoke -n mycc -c '{"Args":["delete","'$type'"]}' -C myc
#sleep 2s
#The following part is for change marble
#peer chaincode query -C myc -n mycc -c '{"Args":["transferMarble","'$type'","jerry"]}'
#read marble
peer chaincode query -C myc -n mycc -c '{"Args":["readMarble","'$type'"]}'
done
    fi
done

