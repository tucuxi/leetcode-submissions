object Solution {
    def groupAnagrams(strs: Array[String]): List[List[String]] = {
        strs.groupBy(_.sorted).values.map(_.toList).toList        
    }
}