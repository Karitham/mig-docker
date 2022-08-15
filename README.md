# mig-docker

A simple cli tool to copy migration files from their original location (`./db/migrations`) to the specified out directory.

Params:
`-up` specifies the up filename. default is `up.sql`.
`-migrations` specifies the migration directory. default is `./db/migrations`.

This tool assumes your migration directories are named with only numbers (such as unix timestamps of creation, or sequential numbers).

It will order them by this same number found in the directory name.
