package main

import "fmt"

func main() {

	path := "C:\\Users\\shahr\\Downloads\\"
	aiFolderName := "AI-Images"

	totalImgsOrganized := OrganizeBingGeneratedImgs(path, aiFolderName)

	fmt.Printf("%v Images moved to AI-Images folder.\n", totalImgsOrganized)
}
