/* Class for indexing system */

#include <iostream>
#include "indexing.h"

Indexing::Indexing(int k) {
	vector<unordered_map<string, forward_list<string>>> new_vec(k);
	hfs = new_vec;
}

Indexing::Indexing() {
	vector<unordered_map<string, forward_list<string>>> new_vec(1200);
	hfs = new_vec;
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
	for (auto it = records.begin(); it != records.end(); it++) {
		cout <<  *it << ",";
	}
	cout << endl;

	return records;
}
