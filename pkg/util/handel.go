package util

type list struct {
	ID int
	Name string
	TopObj *list
	Pid int
	Child []list
}

func handel(tree *list, v list) {
	if v.Pid == 0 {
		v.TopObj = &list{
			Name: tree.Name,
		}
		tree.Child = append(tree.Child, v)
	} else {
		if len(tree.Child) > 0 {
			for key, value := range tree.Child {
				if value.ID == v.Pid {
					v.TopObj = &list {
						Name: value.Name,
					}
					tree.Child[key].Child = append(tree.Child[key].Child, v)
				} else {
					handel(&tree.Child[key], v)
				}
			}
		}
	}
}
