package main

import (
    "path/filepath"
    "io/fs"
    "os"
)


func WalkRepo(repo_path string) ([]string, []string, error) {
    current_dir, err := os.Getwd()
    err = os.Chdir(repo_path)
    if err != nil {
        return nil, nil, err
    }
    var directories []string
    var files [] string
    err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() && info.Name() != "."{
            directories = append(directories, path)
        } else if !info.IsDir() {
            files = append(files, path)
        }

        return nil
    })

    if err != nil {
        return nil, nil, err
    }
    err = os.Chdir(current_dir)
    if err != nil {
        panic("Cannot ch back to original directory")
    }
    return directories, files, nil
}



func CloneTest(dirs []string, files []string) error {
    err := os.Chdir("clone_tests/test")
    if err != nil {
        return err
    }
    for _, dir := range dirs {
        os.Mkdir(dir, 0755)
    }

    return nil
}
