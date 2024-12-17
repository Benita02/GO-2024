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

// 2. Skill Leveling and Character Development
/* A system where completing tasks helps the user level up in various skills
like productivity, creativity, or fitness. As they complete tasks,
their XP is tracked and used to level up skills and unlock new character abilities.*/

// UserProfile struct to track the user’s skills, XP, and character level.
type UserProfile struct {
	Name   string
	XP     int            // Total XP of the user
	Level  int            // Current character level
	Skills map[string]int //Skills with their corresponding levels
}

// Incrementing XP & Leveling Up
/* When a user completes a task, they earn XP,
which may increase their overall level or unlock new skills.
Tasks of higher difficulty give more XP, and the user levels up after reaching certain XP thresholds.
*/
// completeTask method
func (u *UserProfile) completeTask(task Task) {
	u.XP += task.XP
	if u.XP >= 225 { // XP threshold for level up
		u.Level++
		u.XP = 0 // Reset XP after leveling up
	}
	// Assign skill improvements based on task difficulty
	if task.Difficulty == 3 {
		u.Skills["Discipline"] += task.XP / 2
	}
	// need to add other skills to level up

}
