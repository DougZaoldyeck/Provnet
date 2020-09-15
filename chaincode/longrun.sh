#!/bin/bash
# we have 32*100 minhash values


for((a = 0; a <= 99; a++))
do
type=$(head  -c 4000 /dev/urandom | od -A n -t x | tr  -ds ' ' '\n' )
type=$(echo $type | tr -d '\n')
type=$(echo $type | tr -d ' ')
echo "$type"
peer chaincode invoke -n provnet -c '{"Args":["initSharing","prevhash111","current_block_hash","owner","'$type'","Changhao","Nakamoto","futurehash333","randomness000","ca68a408f1e827b76127454fd66f21e","e7b6468a8d5e66193629bd3054af2d87","73db234546af330c9b14de982a5796c3","21ec23beab59a5babf080015de73f42c","a23b01bd6784dcb249f397d52c538585","39e048cb483017d29bfcd6b66244e7d8"]}' -C myc
sleep 0.1s
if [ `expr $a % 20` == 19 ]
then
sleep 3s
for((b = 0; b <= 50; b++))
do
type=$(head  -c 4000 /dev/urandom | od -A n -t x | tr  -ds ' ' '\n' )
type=$(echo $type | tr -d '\n')
type=$(echo $type | tr -d ' ')
echo "$type"
peer chaincode query -C myc -n provnet -c '{"Args":["queryDataByOwner","owner","'$type'"]}'

done
fi
done
