// TODO: This is slightly over-engineered
// The directories & files values can just be the name.
// The directories & files keys can be the full path

// TODO: I can also try an n-ary directory system.
import("slices")
type FileSystem struct {
    directories     map[string]map[string]bool
    files           map[string][]string
    fileContents    map[string]string
}


func Constructor() FileSystem {
    fs := FileSystem{
        directories:    make(map[string]map[string]bool),
        files:          make(map[string][]string),
        fileContents:   make(map[string]string),
    }
    fs.directories["/"] = make(map[string]bool)
    fs.files["/"] = make([]string, 0)
    return fs
}


func (this *FileSystem) Ls(path string) []string {
    // Confirm if it is a file in fileContents
    if _, okFC := this.fileContents[path]; okFC {
        ks := strings.Split(path, "/")
        k := ks[len(ks)-1]
        return []string{k}
    }
    var res []string
    dirs, okd := this.directories[path]
    if !okd { return res }
    // Get the list of directories
    for k, _ := range dirs {
        ks := strings.Split(k, "/")
        if len(ks) == 0 {
            // Shouldn't happen as ls input is correct
            // fmt.Println("error")
            continue
        }
        k = ks[len(ks)-1]
        res = append(res, strings.Trim(k, "/"))
    }
    // Get the list of files
    if files, okf := this.files[path]; okf {
        for _, k := range files {
            ks := strings.Split(k, "/")
            if len(ks) == 0 {
                // Shouldn't happen as ls input is correct
                // fmt.Println("error")
                continue
            }
            k = ks[len(ks)-1]
            res = append(res, strings.Trim(k, "/"))
        }
        res = append(res, )
    }
    slices.Sort(res)
    return res
}


func (this *FileSystem) Mkdir(path string) {
    // Get the individual directory names
    pathList := strings.Split(path, "/")
    // Since the first one will be empty, we'll put the root "/" in it
    if len(pathList) == 0 {
        return
    }
    pathList[0] = "/"
    // Start building the existing parent path from "/"
    curPath := "/"
    pi := 1
    last := len(pathList)-1
    for i := 1; i <= last; i++ {
        pi = i
        p := "/" + pathList[i]
        if i == 1 {
            p = pathList[i]
        }
        if !this.directories[curPath][curPath+p] {
            break
        }
        curPath += p
    }
    // curPath now has the parent path
    // We need to create the directory structure from that point on.
    for pi <= last {
        // Add subdirectories
        prevCurPath := curPath
        if pi == 1 {
            curPath = "/" + pathList[pi]
        } else {
            curPath += "/" + pathList[pi]
        }
        this.directories[prevCurPath][curPath] = true
        this.directories[curPath] = make(map[string]bool)
        this.files[curPath] = make([]string, 0)
        pi++
    }
}


func (this *FileSystem) AddContentToFile(filePath string, content string) {
    curCuntents, ok := this.fileContents[filePath]
    if !ok {
        // Let's create the file entry in the files list
        pathList := strings.Split(filePath, "/")
        parent := "/" + strings.Join(pathList[1:len(pathList)-1], "/")
        this.files[parent] = append(this.files[parent], filePath)
        this.fileContents[filePath] = ""
        curCuntents = ""
    }
    curCuntents += content
    this.fileContents[filePath] = curCuntents
}


func (this *FileSystem) ReadContentFromFile(filePath string) string {
    return this.fileContents[filePath]
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param4 := obj.ReadContentFromFile(filePath);
 */
