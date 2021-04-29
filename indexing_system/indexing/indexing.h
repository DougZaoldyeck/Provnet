#ifndef INDEXING_H
#define INDEXING_H

#include <string>
#include <unordered_map>
#include <forward_list>
#include <vector>
#include <iterator>
#include "sha256.h"
using std::cout;
using std::endl;
using std::string;
using std::vector;
using std::unordered_map;
using std::forward_list;
using std::make_move_iterator;
using std::begin;
using std::end;
using std::distance;


class Indexing
{
protected:
	int k; // number of hash functions
	vector<unordered_map<string, vector<string>>> hfs; // vector of hash functions

public:
	Indexing();
	Indexing(int);
	~Indexing();
	
	void add_record(int n, string record, string mh_val);
	vector<string> list_records(int n, string mh_val);
	void next_record(int hf, string record, string nmh_val, vector<string> & similar);
};

#endif
