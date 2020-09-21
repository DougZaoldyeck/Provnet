#!/bin/bash
# we have 32*n minhash values
# 800: 2560
# 1000: 3200
# 1200: 3840


for((a = 0; a <= 999; a++))
do
type=$(head  -c 3840 /dev/urandom | od -A n -t x | tr  -ds ' ' '\n' )
type=$(echo $type | tr -d '\n')
type=$(echo $type | tr -d ' ')
echo "$type"
peer chaincode invoke -n provnet -c '{"Args":["initSharing","prevhash111","current_block_hash","owner","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","Changhao","Nakamoto","futurehash333","randomness000","ca68a408f1e827b76127454fd66f21e","e7b6468a8d5e66193629bd3054af2d87","73db234546af330c9b14de982a5796c3","21ec23beab59a5babf080015de73f42c","a23b01bd6784dcb249f397d52c538585","39e048cb483017d29bfcd6b66244e7d8","helper"]}' -C myc
sleep 0.3s
if [ `expr $a % 100` == 99 ]
then
sleep 2s
for((b = 0; b <= 49; b++))
do
#The following part is for MinHash comparison
#type=$(head  -c 2560 /dev/urandom | od -A n -t x | tr  -ds ' ' '\n' )
#type=$(echo $type | tr -d '\n')
#type=$(echo $type | tr -d ' ')
#echo "$type"
#peer chaincode query -C myc -n provnet -c '{"Args":["queryDataByOwner","helper","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'"]}'

#The following part is for Record updates
peer chaincode invoke -n provnet -c '{"Args":["updateSharing","'$type'","NO","4a1f93bb46c194381e0f54feeceec4f9","423b7ca687efc3ee286aae75a5b0363c"]}' -C myc

#The following part is for Record reading
#peer chaincode query -C myc -n provnet -c '{"Args":["readSharing","'$type'"]}'

#The following part is for inser & deletion
#peer chaincode invoke -n provnet -c '{"Args":["initSharing","prevhash111","current_block_hash","owner","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","'$type'","Changhao","Nakamoto","futurehash333","randomness000","ca68a408f1e827b76127454fd66f21e","e7b6468a8d5e66193629bd3054af2d87","73db234546af330c9b14de982a5796c3","21ec23beab59a5babf080015de73f42c","a23b01bd6784dcb249f397d52c538585","39e048cb483017d29bfcd6b66244e7d8","helper"]}' -C myc
#sleep 2s
#peer chaincode invoke -C myc -n provnet -c '{"Args":["delete","'$type'"]}'

#sleep 2s
done
fi
done
