class TaskManager(tasks: List<List<Int>>) {

    data class Item(
        val userId: Int,
        val taskId: Int,
        val priority: Int,
    )

    private val tasksById = tasks
        .map { (userId, taskId, priority) -> taskId to Item(userId, taskId, priority) }
        .toMap()
        .toMutableMap()

    private val orderedTasks = tasksById
        .values
        .toSortedSet(compareBy({ -it.priority }, { -it.taskId }))

    fun add(userId: Int, taskId: Int, priority: Int) {
        val item = Item(userId, taskId, priority)
        tasksById[taskId] = item
        orderedTasks.add(item)
    }

    fun edit(taskId: Int, newPriority: Int) {
        val item = tasksById[taskId]!!
        orderedTasks.remove(item)

        val newItem = Item(item.userId, taskId, newPriority)
        tasksById[taskId] = newItem
        orderedTasks.add(newItem)
    }

    fun rmv(taskId: Int) {
        val item = tasksById.remove(taskId)
        orderedTasks.remove(item)
    }

    fun execTop(): Int {
        return if (orderedTasks.isNotEmpty()) {
            val item = orderedTasks.first()
            orderedTasks.remove(item)
            tasksById.remove(item.taskId)
            item.userId
        } else -1
    }

}

/**
 * Your TaskManager object will be instantiated and called as such:
 * var obj = TaskManager(tasks)
 * obj.add(userId,taskId,priority)
 * obj.edit(taskId,newPriority)
 * obj.rmv(taskId)
 * var param_4 = obj.execTop()
 */