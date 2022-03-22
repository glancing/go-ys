package loader

type LoadedTask struct {
	Sku string
	Profile string
	Mode string
}

var loadedTasks = []LoadedTask{
	{
		Sku: "B75571",
		Profile: "2324",
		Mode: "default",
	},
}

func addToTasks() {
	//246 is max without socket
	for i := 1; i < 5; i++ {
		loadedTasks = append(loadedTasks, LoadedTask{
			Sku: "B75571",
			Profile: "2324",
			Mode: "default",
		})
	}
}



func ReturnLoadedTasks() []LoadedTask {
	addToTasks()
	return loadedTasks
}

