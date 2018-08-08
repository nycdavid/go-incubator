package main

// recursive function
// dirents lists all entries in the directory
// => some are files and some are directories
// if a file, send filesize on channel
// else walkDir into the directory
func walkDir(dir string, fileSizes chan<- int64) {
}
