package repository

import (
	"fmt"
	"sync"
)

type List struct {
	id    int
	tasks map[int]Task
	mtx   sync.RWMutex
}

func NewTask() *List { // здесь мы вернули готовый экземпляр мапы  , чтобы вернуть в наш хендлер
	return &List{
		tasks: make(map[int]Task),
		id:    1,
	}
}

func (l *List) AddTask(task Task) Task {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task.ID = l.id
	l.id++

	l.tasks[task.ID] = task

	return task

}

func (l *List) GetTask(id int) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, ok := l.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("не найдено")
	}
	return task, nil
}

func (l *List) DeleteTask(id int) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("такой задачи нету")
	}
	delete(l.tasks, id)
	return Task{}, nil

}
