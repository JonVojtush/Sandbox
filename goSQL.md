ou can use os/exec package of standard library. No database driver is required. Code would look something like this for postgreSQL:

cmd := exec.Command("psql", "-U", psqlUser, "-h", psqlHost, "-d", psqlDBName, "-a", "-f", sqlFilePath)

var out, stderr bytes.Buffer

cmd.Stdout = &out
cmd.Stderr = &stderr

err := cmd.Run()
if err != nil {
    log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
}

-----------------------------------------------

In the following example, I'm using pgx (https://godoc.org/github.com/jackc/pgx) implementation of sql driver for Postgres. You can do it with any driver implementation of your choice.

path := filepath.Join("path", "to", "script.sql")

c, ioErr := ioutil.ReadFile(path)
if ioErr != nil {
   // handle error.
}
sql := string(c)
_, err := *pgx.Conn.Exec(sql)
if err != nil {
  // handle error.
}