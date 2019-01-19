package main

import (
	"os"
)

func siteName() string {
	return os.Getenv("SITE_NAME")
}

func logo() string {
	return os.Getenv("LOGO")
}

func domain() string {
	return os.Getenv("DOMAIN")
}

func backgroundImage() string {
	return os.Getenv("BACKGROUND_IMAGE")
}

func backgroundColor() string {
	return os.Getenv("BACKGROUND_COLOR")
}
