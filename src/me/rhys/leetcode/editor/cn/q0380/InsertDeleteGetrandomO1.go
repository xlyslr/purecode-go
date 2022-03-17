//Implement the RandomizedSet class:
//
//
// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present.
//Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns
//true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements (
//it's guaranteed that at least one element exists when this method is called).
//Each element must have the same probability of being returned.
//
//
// You must implement the functions of the class such that each function works
//in average O(1) time complexity.
//
//
// Example 1:
//
//
//Input
//["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove",
//"insert", "getRandom"]
//[[], [1], [2], [2], [], [1], [2], []]
//Output
//[null, true, false, true, 2, true, false, 2]
//
//Explanation
//RandomizedSet randomizedSet = new RandomizedSet();
//randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was
//inserted successfully.
//randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
//randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now
//contains [1,2].
//randomizedSet.getRandom(); // getRandom() should return either 1 or 2
//randomly.
//randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now
//contains [2].
//randomizedSet.insert(2); // 2 was already in the set, so return false.
//randomizedSet.getRandom(); // Since 2 is the only number in the set,
//getRandom() will always return 2.
//
//
//
// Constraints:
//
//
// -2³¹ <= val <= 2³¹ - 1
// At most 2 * 10⁵ calls will be made to insert, remove, and getRandom.
// There will be at least one element in the data structure when getRandom is
//called.
//
// Related Topics 设计 数组 哈希表 数学 随机化 👍 439 👎 0

package q0380

//leetcode submit region begin(Prohibit modification and deletion)
import "math/rand"

// RandomizedSet 插入和删除O1肯定需要借助hashmap；随机取出O1数组下标就够了
// 删除的时候要知道元素的索引，所以hashmap应该存值所对应的索引
type RandomizedSet struct {
	nums     []int       // 值列表
	valToIdx map[int]int // <值,值在arr中的索引>
}

func Constructor() RandomizedSet {
	return RandomizedSet{nums: []int{}, valToIdx: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIdx[val]; ok {
		return false
	} else {
		this.valToIdx[val] = len(this.nums)
		this.nums = append(this.nums, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.valToIdx[val]; ok {
		// 交换目标元素到最后一个节点，然后移除就是O(1)了，不需要移动其他元素
		lastIdx := len(this.nums) - 1
		this.valToIdx[this.nums[lastIdx]] = i                               // 最后一个元素换索引，要移除的就不换了
		this.nums[lastIdx], this.nums[i] = this.nums[i], this.nums[lastIdx] // 元素交换
		this.nums = this.nums[0:lastIdx]                                    // 从尾部移除val
		delete(this.valToIdx, val)                                          // 移除val的索引
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {

	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
