# Documentation, migrations service

This service is designed to run migrations to Postgres DBMS.<br/>
Config file described in the directory `internal/app/config/config.json` <br/>
	`migrations_url` - URL of migrations directory. <br/>

Each migration MUST named with below template:<br/>
   `{version}_{name}.{side}.sql`<br/>
   where  <br/>
   `version` - migration version (example, 00001 or 1} <br/>
                `name`      - migration name (example, create_nomenclature_table} <br/>
                `side`      - up or down, up is used to migrate, down is used to move previous migration. <br/>
                  
  **Down migrations usually use to remove changes of up migrations. <br/>
  
 Example:<br/>
	   `/.../migrations/00001_initial.up.sql`<br/>
	   `/.../migrations/00001_initial.down.sql`<br/>
	   `/.../migrations/00002_create_showcases.up.sql`<br/>
	   `/.../migrations/00002_remove_showcases.down.sql`<br/>

## Usage

```NAME:
    app
USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
   up          up to the last migration version
   goto, gt    migrates either up or down to the specified version
   version, v  Print current migration version
   drop        drops EVERTHING inside database
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
   ```

You can also use migrate from [golang-migrate](github.com/golang-migrate/migrate)
