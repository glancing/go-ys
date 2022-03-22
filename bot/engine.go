package bot

import (
	"context"
	"time"
	"sync"
	"github.com/glancing/go-ys/loader"
	"github.com/glancing/go-ys/tasks/yeezysupply/types"
	http "github.com/DestroyerBots/httpclient/net/http"
	"net/url"
)

var taskMutex = sync.RWMutex{}

var tasks = make(map[string]*Task)

type TaskInternal struct {
	ProductTitle string
	QueueStarted bool
	Config types.ApiYsConfig
	PixelConfig types.PixelConfig
	ParsedUrl *url.URL
}

type TaskHandlers map[string]func(*Task, *TaskInternal)(string)

type Task struct {
	Id string
	Sku string
	Proxy string
	Transport *http.Transport
	ctx context.Context
	cancel context.CancelFunc
	taskState string
	Client *http.Client
	Handlers TaskHandlers
	wg *sync.WaitGroup
}

func AddTask(uuid string, loadedTask loader.LoadedTask, proxy string,wg *sync.WaitGroup) *Task {
	taskMutex.RLock()
	defer taskMutex.RUnlock()
	taskCtx, cancel := context.WithCancel(context.Background())
	tasks[uuid] = &Task{
		Id: uuid,
		Sku: loadedTask.Sku,
		Proxy: proxy,
		ctx: taskCtx,
		cancel: cancel,
		Handlers: TaskHandlers{},
		wg: wg,
	}
	return tasks[uuid]
}

func StopTask(uuid string) {
	tasks[uuid].taskState = "stopped"
	taskMutex.Lock()
	defer taskMutex.Unlock()
	tasks[uuid].cancel()
	delete(tasks, uuid)
}

func (t *Task) RunTask() {
	defer t.wg.Done()
	t.taskState = "start"
	internalTask := &TaskInternal{}
	for {
		select {
			case <- t.ctx.Done():
				t.taskState = "stopped"
				return
			default:
				newTaskState := t.Handlers[t.taskState](t, internalTask)

				if newTaskState == "finished" {
					StopTask(t.Id)
					return
				}

				if _, ok := t.Handlers[newTaskState]; ok {
					t.taskState = newTaskState
				}
		}
		time.Sleep(1 * time.Millisecond)
	}
}