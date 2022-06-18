package topkfrequentelements

import (
	"sort"

)

type kv struct {
    key, val int
}

func topKFrequent(nums []int, k int) []int {
    hashmap := make(map[int]*kv)
    for _, n := range nums {
        k, ok := hashmap[n]
        if !ok {
           hashmap[n] = &kv{key: n, val: 1}
            continue
        }

        k.val++     
    }

    vec := make([]*kv, 0, len(hashmap))
    for _, hm := range hashmap {
        vec = append(vec, hm)
    }

    sort.Slice(vec, func(i, j int) bool {
        return vec[i].val > vec[j].val
    })

    res := make([]int, k)
    for i := 0; i < k; i++ {
       res[i] = vec[i].key
    }

    return res
}
