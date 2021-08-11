package gui

type mydesktop struct {
	x int
}


func CreateDesktop()(mydesktop){
	var d mydesktop
	d.x = 1
	return d
}

