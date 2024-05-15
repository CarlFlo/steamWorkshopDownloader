package steamworkshop

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/steamWorkshopDownloader/config"
)

type Item struct {
	Success    bool
	PathToFile string
	Bytes      int
	ErrorMsg   string
}

func DownloadWorkshopItem(workshopArgument string) (*Item, error) {

	// Path to the SteamCMD executable
	//cmdPath := `C:\Program Files (x86)\Steam\steamcmd.exe`
	cmdPath := config.CONFIG.PathToSteamCMD

	wArgs := strings.Split(workshopArgument, " ")

	// Arguments to pass to the SteamCMD executable
	args := []string{"+login", "anonymous", "+workshop_download_item", wArgs[1], wArgs[2], "+quit"}

	// Create the command
	cmd := exec.Command(cmdPath, args...)

	var outputBuf bytes.Buffer
	cmd.Stdout = &outputBuf
	cmd.Stderr = &outputBuf

	// Run the command
	err := cmd.Run()
	if err != nil {
		malm.Fatal("Failed to run command: %v\nOutput: %s", err, outputBuf.String())
	}

	// Debug - Print the output
	//log.Printf("\n\nCommand output:\n%s\n\n=======================================================\n", outputBuf.String())

	// Analyze the buffer to see if we could download the item or not
	workshopItem, err := parseBuffer(outputBuf)
	if err != nil {
		return nil, err
	}

	return workshopItem, nil
}

func parseBuffer(outputBuf bytes.Buffer) (*Item, error) {
	// Define regular expressions for success, download failure, and login failure
	successRegex := `Success\. Downloaded item \d+ to "([^"]+)" \((\d+) bytes\)`
	downloadFailureRegex := `ERROR! Download item \d+ failed \(([^)]+)\)`
	loginFailureRegex := `FAILED \(([^)]+)\)`

	// Check for success case
	reSuccess := regexp.MustCompile(successRegex)
	successMatch := reSuccess.FindStringSubmatch(outputBuf.String())
	if successMatch != nil {
		// Parse the path and bytes from the match
		pathToFile := successMatch[1]
		bytes, err := strconv.Atoi(successMatch[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing bytes: %v", err)
		}

		// Create the file struct and populate it
		downloadedFile := &Item{
			Success:    true,
			PathToFile: pathToFile,
			Bytes:      bytes,
		}

		return downloadedFile, nil
	}

	// Check for download failure case
	reDownloadFailure := regexp.MustCompile(downloadFailureRegex)
	downloadFailureMatch := reDownloadFailure.FindStringSubmatch(outputBuf.String())
	if downloadFailureMatch != nil {
		// Create the file struct with an error message
		downloadedFile := &Item{
			Success:  false,
			ErrorMsg: fmt.Sprintf("Download failed: %s", downloadFailureMatch[1]),
		}
		return downloadedFile, nil
	}

	// Check for login failure case
	reLoginFailure := regexp.MustCompile(loginFailureRegex)
	loginFailureMatch := reLoginFailure.FindStringSubmatch(outputBuf.String())
	if loginFailureMatch != nil {
		// Create the file struct with an error message
		downloadedFile := &Item{
			Success:  false,
			ErrorMsg: fmt.Sprintf("Login failed: %s", loginFailureMatch[1]),
		}
		return downloadedFile, nil
	}

	// TODO: log the buffer with the unexpected input
	return nil, fmt.Errorf("unhandled error")
}

/* success output
Steam Console Client (c) Valve Corporation - version 1714855729
-- type 'quit' to exit --
Loading Steam API...OK

Connecting anonymously to Steam Public...OK
Waiting for client config...OK
Waiting for user info...OK
Downloading item 2732294885 ...
Success. Downloaded item 2732294885 to "D:\SteamLibrary\steamapps\workshop\content\108600\2732294885" (162824 bytes) Redirecting stderr to 'C:\Program Files (x86)\Steam\logs\stderr.txt'
Logging directory: 'C:\Program Files (x86)\Steam/logs'
[  0%] Checking for available updates...
[----] Verifying installation...
CWorkThreadPool::~CWorkThreadPool: work processing queue not empty: 2 items discarded.
*/

/* Failure output
Steam Console Client (c) Valve Corporation - version 1714855729
-- type 'quit' to exit --
Loading Steam API...OK

Connecting anonymously to Steam Public...OK
Waiting for client config...OK
Waiting for user info...OK
Downloading item 27322948851 ...
ERROR! Download item 27322948851 failed (File Not Found).Redirecting stderr to 'C:\Program Files (x86)\Steam\logs\stderr.txt'
Logging directory: 'C:\Program Files (x86)\Steam/logs'
[  0%] Checking for available updates...
[----] Verifying installation...
CWorkThreadPool::~CWorkThreadPool: work processing queue not empty: 2 items discarded.
*/

/* Failed login
Output: Steam Console Client (c) Valve Corporation - version 1714855729
-- type 'quit' to exit --
Loading Steam API...OK
Logging in user 'planonymous' to Steam Public...
password: FAILED (Invalid Password)
Redirecting stderr to 'C:\Program Files (x86)\Steam\logs\stderr.txt'
Logging directory: 'C:\Program Files (x86)\Steam/logs'
[  0%] Checking for available updates...
[----] Verifying installation...
CWorkThreadPool::~CWorkThreadPool: work processing queue not empty: 2 items discarded.
*/
