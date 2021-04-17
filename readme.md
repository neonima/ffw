# fftw - File for the web
A dead simple cli to quickly rename you files to be compatible for the web

## Help page

```text
NAME:
   fftw - A dead simple tool to rename you file for smooth web access!

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --recursive, -r             (default: Use this flag is you want to rename files recursively)
   --extensions .png, -e .png  -e=.png,.pdf | rename only .png et `.pdf` file
   --source value, -s value    (default: source path where to rename files)
   --dry, -d                   (default: shows all the file that will be modified with their new name)
   --help, -h                  show help (default: false)
```