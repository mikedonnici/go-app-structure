package main

import (
	"fmt"
	"github.com/mikedonnici/go-app-structure/internal/pkg1"
	"github.com/mikedonnici/go-app-structure/internal/pkg2"
)

func main() {

	fmt.Println("Main func checking internal packages...")
	pkg1.Check()
	pkg2.Check()
}
