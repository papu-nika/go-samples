package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

type Options struct {
	op_l *bool
	op_a *bool
}

func main() {
	option := Options{}

	option.op_l = flag.Bool("l", false, "use a long listing format")
	option.op_a = flag.Bool("a", false, "all files print")
	flag.Parse()
	argments := flag.Args()
	argc := len(argments)
	if argc == 0 {
		current_dir, err := os.Getwd()
		if err != nil {
			err_prrocess(err)
		}
		do_ls(current_dir, "", option)
		os.Exit(0)
	}
	if argc == 1 {
		do_ls(argments[0], "", option)
		os.Exit(0)
	}

	dir_start, dir_end := sort_args(argments, option)
	sort.Slice(argments[:dir_start], func(i, j int) bool { return argments[i] < argments[j] })
	sort.Slice(argments[dir_start:dir_end+1], func(i, j int) bool { return argments[dir_start+i] < argments[dir_start+j] })
	i := 0
	for ; i < dir_start; i++ {
		if i < dir_start-1 && *option.op_l == true {
			do_ls(argments[i], "", option)
		} else if i < dir_start-1 && *option.op_l == false {
			do_ls(argments[i], "", option)
		} else {
			do_ls(argments[i], "", option)
		}
	}
	if i <= dir_end {
		fmt.Printf("\n")
	}
	for ; i <= dir_end; i++ {
		if dir_start <= i && i <= dir_end {
			fmt.Printf("%s:\n", argments[i])
			do_ls(argments[i], "", option)
			if i != dir_end {
				fmt.Printf("\n")
			}
		}
	}
}

func sort_args(argments []string, option Options) (file_count, error_count int) {
	last_index := len(argments) - 1
	file_count = 0
	error_count = 0
	for i := 0; i <= last_index-error_count; i++ {
		file_info, err := os.Stat(argments[i])
		if err != nil {
			err_prrocess(err)
			argments[i], argments[last_index-error_count] = argments[last_index-error_count], argments[i]
			error_count++
			i--
			continue
		}
		if file_info.IsDir() == true {
			continue
		} else {
			argments[i], argments[file_count] = argments[file_count], argments[i]
			file_count++
		}
	}
	//fmt.Printf("err=%d, index=%d, file=%d\n", error_count, last_index, file_count)
	return file_count, last_index - error_count
}

func do_ls(path string, if_file_last_str string, option Options) {
	file_info, err := os.Stat(path)
	if err != nil {
		err_prrocess(err)
		return
	}
	if file_info.IsDir() == true {
		if *option.op_a == true {
			if *option.op_l == true {
				print_file(".", path, option, "\n")
				print_file("..", path, option, "\n")
			} else {
				print_file(".", path, option, "  ")
				print_file("..", path, option, "  ")
			}
		}
		files, err := ioutil.ReadDir(path)
		if err != nil {
			err_prrocess(err)
			return
		}
		var i int
		for i = 0; len(files)-1 > i; i++ {
			if *option.op_a == false && strings.HasPrefix(files[i].Name(), ".") {
				continue
			}
			if *option.op_l == true {
				print_file(files[i].Name(), path, option, "\n")
			} else {
				print_file(files[i].Name(), path, option, "  ")
			}
		}
		print_file(files[i].Name(), path, option, "\n")
		fmt.Printf("%s", if_file_last_str)
	} else {
		if *option.op_a == false && strings.HasPrefix(path, ".") {
			return
		}
		if *option.op_l == true {
			print_file(path, "", option, "\n")
		} else {
			print_file(path, "", option, "  ")
		}
		//print_file(path, "", option, if_file_last_str)
		fmt.Printf("%s", if_file_last_str)
	}
}

func err_prrocess(err error) {
	fmt.Println(err)
}

func print_file(file string, path string, option Options, last string) {
	if *option.op_l == false {
		fmt.Printf("%s%s", file, last)
		return
	}

	var file_path string
	if strings.HasSuffix(path, "/") || path == "" {
		file_path = path + file
	} else {
		file_path = path + "/" + file
	}
	fileinfo, err := os.Stat(file_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	var owner, group string
	stat := fileinfo.Sys().(*syscall.Stat_t)
	uid := strconv.Itoa(int(stat.Uid))
	u, err := user.LookupId(uid)
	if err != nil {
		owner = uid
	} else {
		owner = u.Username
	}
	gid := strconv.Itoa(int(stat.Gid))
	g, err := user.LookupGroupId(gid)
	if err != nil {
		fmt.Println(err)
		group = gid
	} else {
		group = g.Name
	}
	if stat != nil {
		fmt.Printf("%s 1 %s %s\t%d\t%s %s",
			fileinfo.Mode().String(), owner, group, fileinfo.Size(),
			fileinfo.ModTime().Format("02 Jan 06 15:04"), fileinfo.Name())
		fmt.Printf("%s", last)
	}
}
