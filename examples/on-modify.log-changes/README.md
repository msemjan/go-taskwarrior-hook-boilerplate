# Logging Tag Hook for Taskwarrior

This simple example shows how to create a simple logger for Taskwarrior, which saves old and new task in a log file when the on-modify hook is executed.

## How to build and install

```
go build .
mv on-modify.log-changes ~/.task/hooks/
```

You have to specify the file where will the hook log the data. This is done using a `TASKWARRIOR_LOG` environment variable. You can simply add a line like this one to your `.bashrc`/`.zshrc`:
```
export TASKWARRIOR_LOG='/path/to/log/file'
```
