/* Class for indexing system */

#include <iostream>
#include "indexing.h"

Indexing::Indexing(int k) {
	this->k = k;
	vector<unordered_map<string, forward_list<string>>> new_vec(this->k);
	hfs = new_vec;
	vector<unordered_map<string, vector<string>>> new_vec_v(this->k);
	hfs_v = new_vec_v;
}

Indexing::Indexing() {
	this->k = 1200;
	vector<unordered_map<string, forward_list<string>>> new_vec(1200);
	hfs = new_vec;
	vector<unordered_map<string, vector<string>>> new_vec_v(this->k);
	hfs_v = new_vec_v;
}

Indexing::~Indexing() { }

void Indexing::add_record(int n, string record, string mh_val) {
	// Add a record to the front of its minhash's linked list
	hfs.at(n)[mh_val].push_front(record);
}

vector<string> Indexing::list_records(int n, string mh_val) {
	// Iterate over linked list, push into vector
	vector<string> records;
	forward_list<string> fll = hfs.at(n)[mh_val];

	for (auto it = fll.begin(); it != fll.end(); it++) {
		records.push_back(*it);
	}

	// print them out
	cout << "hfn: " << n << " mh_val: " << mh_val << " R#: " << endl;
	for (auto it = records.begin(); it != records.end(); it++) {
		cout <<  *it << ",";
	}
	cout << endl;

	return records;
}

void Indexing::convert() {
	// convert the initial FLL into a vector 
	for (int i = 0; i < this->k; i++) {
		//cout << "hf index is " << i << endl;
		unordered_map<string, forward_list<string>>::iterator it = hfs.at(i).begin();
		while (it != hfs.at(i).end()) {
			
			hfs_v.at(i)[it->first].reserve(distance(it->second.begin(), it->second.end()));
			hfs_v.at(i)[it->first].insert(hfs_v.at(i)[it->first].end(), it->second.begin(), it->second.end());
			it++;
		}
	}
}

void Indexing::next_record(int hf, string record, string mh_val, vector<string> & similar) {
	// first, find the list of records with similar minhashes
	for (auto it = hfs_v.at(hf)[mh_val].begin(); it != hfs_v.at(hf)[mh_val].end(); it++) {
		// TODO question: does this FOR loop make the process of converting FLL to vector useless?
		similar.push_back(*it);
	}
	// Add a record to the front of its minhash's vector
	hfs_v.at(hf)[mh_val].insert(hfs_v.at(hf)[mh_val].begin(), record);
}

void Indexing::iterate(int hf, string mh_val) {
	// iterate over all the hash functions and print test data out.
	/*for (int i = 0; i < k; i++){
		//print stuff.
	}*/

}
