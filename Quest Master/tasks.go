package main

// 1. Quest Mapping: Task List by Difficulty
type Task struct {
	Name        string
	Difficulty  int  //Task difficulty (1 = easy, 2 = medium, 3 = hard)
	XP          int  //XP earned upon completion
	Completed   bool // Whether the task has been completed
	LevelUnlock int  // Level required to unlock the task
}

var tasks []Task // list of tasks is kept in an array or slice.

// Task Progression:users must complete tasks of increasing difficulty to level up
func unlockNextTasks(userLevel int) []Task {
	var unlockedTasks []Task
	for _, task := range tasks {
		if task.LevelUnlock <= userLevel {
			unlockedTasks = append(unlockedTasks, task)
		}
	}
	return unlockedTasks
}
