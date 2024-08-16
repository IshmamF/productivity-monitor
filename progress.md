## TO DO NEXT: 
- [x] Create function for currentTimestamp in Utils
- [x] Look up how to check if table and file exists or not (to handle new and previous users) 
- [x] Create function to process string recieved from GetForegroundWindowData()
- [x] Get user input on when to send alert, might need to use a counter to keep track of time passed  
- [x] Need to look into how to execute other functions like viewing statistics or running the alert while the logging occurs 
	> look into channels <br>
	go routine , if forever loop , main routine waits for program to exit <br>
	in each infinite loop, have switch cases for the programs to communicate with each other <br>
- [x] Stop multiple instances from being ran 
	Lock file , same location (.local/share on mac)
	os.OpenFile to open file
	os.Stat to check file exists
- [x] Create TUI 
- [x] Option to see current session data 
- [ ] Convert data to daily/weekly/all time statistics 
## Feature Ideas
- [ ] Option to be a login program, starts running automatically when you login to computer
- [ ] Send alert when on a blacklisted website/app
  - if website, can have warnings before blocking traffic or completely block
  - can open up a webpage to add additional annoyance 
  - requires CRUD implementation 
- [ ] Option to send alert vs notification vs both
