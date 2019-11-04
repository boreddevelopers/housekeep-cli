package main

// Logger logs [data] into a file with the time it was logged.
func Logger(data string) {
	if toLog {
		debugFile := "debug.log"
		timeString := Concat(GetCurrentTime(), "\n")
		logString := Concat(timeString, data)
		logString = Concat(logString, "\n")
		isExist := DoesFileExist(debugFile)

		if isExist {
			AppendToFile(logString, debugFile)
		} else {
			CreateNewFileWithData(logString, debugFile)
		}
	}
}
