class Solution {
    fun earliestFinishTime(landStartTime: IntArray, landDuration: IntArray, waterStartTime: IntArray, waterDuration: IntArray): Int {
        return landStartTime.indices
            .map { i ->
                val landFinishTime = landStartTime[i] + landDuration[i]
                
                waterStartTime.indices
                    .map { j ->
                        val waterFinishTime = waterStartTime[j] + waterDuration[j]
                        when {
                            landFinishTime <= waterStartTime[j] -> waterFinishTime
                            waterFinishTime <= landStartTime[i] -> landFinishTime
                            else -> minOf(landFinishTime + waterDuration[j], waterFinishTime + landDuration[i])
                        }
                    }
                    .min()
            }
            .min()
    }
}