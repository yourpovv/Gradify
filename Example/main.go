package main

import (
	"fmt"

	"github.com/yourpovv/gradify"
)

func main() {
	fmt.Println(gradify.Gradient("x o x o | Error Preset!      | x o x o", gradify.Error)) // gradify.<preset name>
	fmt.Println(gradify.Gradient("x o x o | Success Preset!    | x o x o", gradify.Success))
	fmt.Println(gradify.Gradient("x o x o | Warning Preset!    | x o x o", gradify.Warning))
	fmt.Println(gradify.Gradient("x o x o | Info Preset!       | x o x o", gradify.Info))
	fmt.Println(gradify.Gradient("x o x o | Christmas Preset!  | x o x o", gradify.Christmas))
	fmt.Println(gradify.Gradient("x o x o | Candy Preset!      | x o x o", gradify.Candy))
	fmt.Println(gradify.Gradient("x o x o | Minty Preset!      | x o x o", gradify.Minty))
	fmt.Println(gradify.Gradient("x o x o | Nectar Preset!     | x o x o", gradify.Nectar))
	fmt.Println(gradify.Gradient("x o x o | Default Preset!    | x o x o", gradify.Default))
	fmt.Println(gradify.Gradient("x o x o | Hollow Preset!     | x o x o", gradify.Hollow))
	fmt.Println(gradify.Gradient("x o x o | Sunset Preset!     | x o x o", gradify.Sunset))
	fmt.Println(gradify.Gradient("x o x o | Aether Preset!     | x o x o", gradify.Aether))
	fmt.Println(gradify.Gradient("x o x o | Blossom Preset!    | x o x o", gradify.Blossom))
	fmt.Println(gradify.Gradient("x o x o | Frostbite Preset!  | x o x o", gradify.Frostbite))
	fmt.Println(gradify.Gradient("x o x o | Matrix Preset!     | x o x o", gradify.Matrix))
	fmt.Println(gradify.Gradient("x o x o | Emberlight Preset! | x o x o", gradify.Emberlight))
	fmt.Println(gradify.Gradient("x o x o | Aurora Preset!     | x o x o", gradify.Aurora))
	fmt.Println(gradify.Gradient("x o x o | Glacier Preset!    | x o x o", gradify.Glacier))
}
