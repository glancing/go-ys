package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/glancing/go-ys/bot"
	"github.com/glancing/go-ys/loader"
	"github.com/glancing/go-ys/tasks/utils"
	"github.com/glancing/go-ys/tasks/yeezysupply"
	"github.com/lithammer/shortuuid/v4"
)

var (
	tasksLoaded []loader.LoadedTask
	proxiesLoaded []string
	profilesLoaded = 0
)

func main() {
	tasksLoaded = loader.ReturnLoadedTasks()
	proxiesLoaded = loader.ReturnLoadedProxies()
	displayMenu()
}

func displayMenu() {
	fmt.Println("welcome to n8 bot!")
	fmt.Println(fmt.Sprintf("Tasks loaded - %d", len(tasksLoaded)))
	fmt.Println(fmt.Sprintf("Proxies loaded - %d", len(proxiesLoaded)))
	fmt.Println(fmt.Sprintf("Profiles loaded - %d", profilesLoaded))
	fmt.Println()
	fmt.Println("Press 5 to start all tasks")
	fmt.Println("Press 3 to setup AYCD")
	fmt.Println("Press 9 to exit")
	fmt.Println()
	answerListener()
}

func answerListener() {
	answer := listenForAnswer()
	switch answer {
		case "5":
			if len(tasksLoaded) == 0 {
				fmt.Println("no tasks loaded!")
				answerListener()
			} else {
				startTasks(tasksLoaded)
			}
		case "3":
			fmt.Println("setting up aycd")
		case "9":
			fmt.Println("exiting")
			os.Exit(0)
		default: 
			fmt.Println("invalid")
			answerListener()
	}
}

func listenForAnswer() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scanner.Scan()
	return scanner.Text()
}

func startTasks(loadedTasks []loader.LoadedTask) {
	fmt.Println("starting tasks")
	var wg sync.WaitGroup
	for _, t := range loadedTasks {
		uuid := shortuuid.New()
		proxy := utils.SelectProxy(proxiesLoaded)
		fmt.Println("Starting task", uuid, proxy)
		task := bot.AddTask(uuid, t, proxy, &wg)
		tasks.PushYeezySupplyHandlers(task)
		wg.Add(1)
		go task.RunTask()
	}
	wg.Wait()
}