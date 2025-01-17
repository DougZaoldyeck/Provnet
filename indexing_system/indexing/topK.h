#ifndef TOPKELEMENTS_H
#define TOPKELEMENTS_H

#include <vector>
#include <stdio.h>
#include <math.h>
#include <map>
#include <string>

using namespace std;

class TopKElements {
	/* this class was originally implemented by Changhao Chenli */
	private:
		vector<string> unique;
		map<string, int> count_map;
	
	public:
		int partition(int left, int right, int pivot_index);
		void quickselect(int left, int right, int k_smallest);

		vector<string> topKFrequent(vector<string>& nums, int k);

};

#endif
