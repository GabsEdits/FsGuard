package core

import (
	"fmt"
	"os"
)

func ValidateSUID(file string, isSUID bool) {
	fileStat, err := os.Stat(file)
	if err != nil {
		fmt.Printf("File %s does not exist! Cannot check suid bit\n", file)
	}

	if fileStat.Mode()&os.ModeSetuid != 0 && !isSUID {
		fmt.Printf("File %s has incorrect suid permission\n", file)
		correctSUID(file, isSUID)
	} else if fileStat.Mode()&os.ModeSetuid == 0 && isSUID {
		fmt.Printf("File %s has incorrect suid permission\n", file)
		correctSUID(file, isSUID)
	} else {
		fmt.Printf("File %s has correct suid bit\n", file)
	}
}

func correctSUID(file string, isSUID bool) {
	fileStat, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	mode := fileStat.Mode()

	if isSUID {
		err = os.Chmod(file, mode|os.ModeSetuid)
	} else {
		err = os.Chmod(file, mode&^os.ModeSetuid)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
