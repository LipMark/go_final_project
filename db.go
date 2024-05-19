package db

appPath, err := os.Executable()
if err != nil {
    log.Fatal(err)
}
dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
_, err = os.Stat(dbFile)

var install bool
if err != nil {
    install = true
}