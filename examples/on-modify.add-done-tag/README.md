# Add "done" Tag Hook for Taskwarrior

This simple example shows how to create a Taskwarrior on-modify hook that appends a `"done"` tag when a status of a task changes from `pending` to `done`.

## How to build and install

```
go build .
mv on-modify.add-done-tag ~/.task/hooks/
```
