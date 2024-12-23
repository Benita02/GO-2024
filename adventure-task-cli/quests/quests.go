package quests

import(
	"os"
)

// Constants for the XP system
const (
	LevelUpThreshold = 225
)

// 1. Quest Mapping: Task List by Difficulty
type Task struct {
	Name        string
	Difficulty  int  //Task difficulty (1 = easy, 2 = medium, 3 = hard)
	XP          float32  //XP earned upon completion
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
	Name  string
	XP    float32 // Total XP of the user
	Level int // Current character level
}

// Incrementing XP & Leveling Up
/* When a user completes a task, they earn XP,
which may increase their overall level
Tasks of higher difficulty give more XP, and the user levels up after reaching certain XP thresholds.
*/
// completeTask method assigns XP based on the task and updates the user's level.
func (u *UserProfile) completeTask(task Task) {
	taskXP := map[int]float32{
		1: 3.75, // Difficulty 1 gets 3 XP
		2: 7.5,  // Difficulty 2 gets 7 XP
		3: 15,   // Difficulty 3 gets 15 XP
	}

	if increment, ok := taskXP[task.Difficulty]; ok {
		task.XP = increment
		u.XP += increment
	}

	func saveReport(task Task, report string) error {
		file, err := os.Create(task.Name + "_report.txt")
		if err != nil {
			return err
		}
		defer file.Close()

	_, err = file.WriteString(report)
	if err != nil {
		return err
	}
	return nil // Return nil to indicate success
}

// busy