package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	if command == "init" {
		err := os.Mkdir(".mini-git", 0755)

		if err != nil {
			if os.IsExist(err) {
				fmt.Println("Mini Git repository already exists")
				return
			}

			fmt.Println("Error:", err)
			return
		}

		err = os.Mkdir(".mini-git/objects", 0755)
		if err != nil {
			fmt.Println("Error creating objects:", err)
			return
		}

		err = os.Mkdir(".mini-git/refs", 0755)
		if err != nil {
			fmt.Println("Error creating refs:", err)
			return
		}

		err = os.Mkdir(".mini-git/refs/heads", 0755)
		if err != nil {
			fmt.Println("Error creating heads:", err)
			return
		}

		file, err := os.Create(".mini-git/refs/heads/main")
		if err != nil {
			fmt.Println("Error creating main:", err)
			return
		}

		_, err = file.WriteString("")
		if err != nil {
			fmt.Println("Error writing to main:", err)
			file.Close()
			return
		}

		file.Close()

		fmt.Println("Initialized empty Mini Git Repository")
		return
	}

	if command == "status" {
		data, err := os.ReadFile(".mini-git/refs/heads/main")

		if err != nil {
			fmt.Println("Error reading main:", err)
			return
		}

		fmt.Println("Current main:", string(data))
		return
	}

	if command == "update-main" {
		if len(os.Args) < 3 {
			fmt.Println("Please provide a commit hash")
			return
		}

		file, err := os.Create(".mini-git/refs/heads/main")

		if err != nil {
			fmt.Println("Error opening main:", err)
			return
		}

		_, err = file.WriteString(os.Args[2])

		if err != nil {
			fmt.Println("Error updating main:", err)
			file.Close()
			return
		}

		file.Close()

		fmt.Println("Updated main to", os.Args[2])
		return
	}

	if command == "commit" {
		if len(os.Args) < 3 {
			fmt.Println("Please provide a commit message")
			return
		}

		message := strings.Join(os.Args[2:], " ")

		parentData, err := os.ReadFile(".mini-git/refs/heads/main")
		if err != nil {
			fmt.Println("Error reading parent commit:", err)
			return
		}

		parent := strings.TrimSpace(string(parentData))

		commitData := fmt.Sprintf(
			"parent: %s\nmessage: %s\n",
			parent,
			message,
		)

		hash := sha256.Sum256([]byte(commitData))
		commitHash := fmt.Sprintf("%x", hash)

		commitPath := filepath.Join(
			".mini-git",
			"objects",
			commitHash,
		)

		err = os.WriteFile(
			commitPath,
			[]byte(commitData),
			0644,
		)

		if err != nil {
			fmt.Println("Error creating commit object:", err)
			return
		}

		err = os.WriteFile(
			".mini-git/refs/heads/main",
			[]byte(commitHash),
			0644,
		)

		if err != nil {
			fmt.Println("Error updating main:", err)
			return
		}

		fmt.Println("Created commit:", commitHash)
		return
	}

	if command == "log" {
		data, err := os.ReadFile(".mini-git/refs/heads/main")

		if err != nil {
			fmt.Println("Error reading main:", err)
			return
		}

		current := strings.TrimSpace(string(data))

		for current != "" {
			commitPath := filepath.Join(
				".mini-git",
				"objects",
				current,
			)

			commitData, err := os.ReadFile(commitPath)

			if err != nil {
				fmt.Println("Error reading commit:", err)
				return
			}

			fmt.Println("commit", current)
			fmt.Println(string(commitData))

			lines := strings.Split(
				string(commitData),
				"\n",
			)

			parent := ""

			for _, line := range lines {
				if strings.HasPrefix(line, "parent:") {
					parent = strings.TrimSpace(
						strings.TrimPrefix(line, "parent:"),
					)
					break
				}
			}

			current = parent

			fmt.Println("--------------------")
		}

		return
	}

	fmt.Println("Unknown command:", command)
}
