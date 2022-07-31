/*

Best Time to Buy and Sell Stock

Solution
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
 

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104

*/

func maxProfit(prices []int) int {
    len1 := len(prices)

    if len1 <= 1 {
        return 0
    }
    
    maxArr := make([]int, len1)
    
    maxArr[len1-1] = prices[len1-1]
    currMax := prices[len1-1]
    
    j := len1 - 2
    for j >= 0 {
        if prices[j] > currMax {            
            currMax = prices[j]
        }        
        maxArr[j] = currMax
        j--
    }
    
    
    currMin := math.MaxInt32
    maxProfit := 0
    
    i := 0
    for i < len1 - 1 {
        if prices[i] < currMin {            
            currMin = prices[i]
        }
        if maxProfit < (maxArr[i+1] - currMin) {
            maxProfit = maxArr[i+1] - currMin
        }
        i++
    }
    return maxProfit     
}
