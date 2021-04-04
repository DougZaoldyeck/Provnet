/* This is an implementation of the indexing system, as described by
 * Changhao Chenli 
 * 
 * Frank Gomulka
 */
#include <iostream>
#include "sha256.h"
#include "indexing.h"

int main(int argc, char *argv[]) {
	int k = 1200; // Number of hash functions.
	Indexing system = Indexing(k);

	// insertion
	int hfn_num = 20;
	string record1 = "R1";
	string hash1 = sha256(record1);
	system.add_record(hfn_num, record1, hash1);
	cout << record1 << " has been added." << endl;

	string record2 = "R2";
	string hash2 = sha256(record2);
	system.add_record(hfn_num, record2, hash2);
	cout << record2 << " has been added." << endl;
	
	string record3 = "R3";
	string hash3 = sha256(record3);
	system.add_record(hfn_num, record3, hash3);
	cout << record3 << " has been added." << endl;

	string record4 = "R4";
	system.add_record(hfn_num, record4, hash3);
	cout << record4 << " has been added." << endl;

	string record5 = "R5";
	system.add_record(hfn_num, record5, hash1);
	cout << record5 << " has been added." << endl;

	string record6 = "R6";
	system.add_record(hfn_num, record6, hash2);
	cout << record6 << " has been added." << endl;
	
	cout << "" << endl;

	// listing records
	cout << "Listing minhash 1:" << endl;
	system.list_records(hfn_num, hash1);
	
	cout << "Listing minhash 2:" << endl;
	system.list_records(hfn_num, hash2);

	cout << "Listing minhash 3:" << endl;
	system.list_records(hfn_num, hash3);

/* TODO: notes
complexity: 

std::set) Search, removal, and insertion operations have logarithmic complexity.

std::list) supports constant time insertion and removal of elements from anywhere in the container. Fast random access is not supported.

list is more efficient to insert bc constant time

std::forward_list) is more space efficient than std::list


also, note that with the forward list, we iterate backwards... could be forwards if we did a doubly-linked list, but that is obviously more space inefficient.. what is your priority?

*/

	return 0;
}
