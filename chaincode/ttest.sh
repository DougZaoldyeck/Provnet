#!/bin/bash
type=marble
x=0
y=0
for((a = 0; a <= 99; a++))
do
    peer chaincode invoke -n mycc -c peer chaincode invoke -n mycc -c '{"Args":["initMarble","'${type}$y$x'","blue","35","tom"]}' -C myc
    sleep 0.1s
    if [ `expr $a % 20` == 19 ]
    then
        sleep 3s
        for((b = 0; b <= 49; b++))
        do
            peer chaincode query -C myc -n mycc -c '{"Args":["transferMarble","'${type}$y$x'","jerry"]}'
        done
    fi
    if [ $x == 9 ]
    then
        let "y+=1"
        let "x=0"
    else
        let "x+=1"
    fi
done

