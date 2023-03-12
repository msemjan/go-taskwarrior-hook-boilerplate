# Boilerplate Code for Taskwarrior Hooks in Golang (Work in Progress)

[Taskwarrior](https://taskwarrior.org/) is a Free and Open Source Software that manages your TODO list from the command line. It allows it's users to create [hooks](https://taskwarrior.org/docs/hooks), which extent the functionality of this tool. The community around Taskwarrior created many examples and tutorials for writing hooks in Python (e.g. [this guide](https://taskwarrior.org/docs/hooks_guide/)), but there is significantly smaller amount of similar resources for other languages.

This repository can be used as boilerplate for writing Taskwarrior hooks in Golang. 

## How to use

First clone this repository:
```
git clone https://github.com/msemjan/go-taskwarrior-hook-boilerplate
```

There are examples of hooks in `examples` folder, which you can take a look at.

If you want to create your own, copy the code from `src` folder, and modify it. Taskwarrior runs hooks when certain events are triggered. The full, up-to-date list of the events can be found in the [hooks documentation](https://taskwarrior.org/docs/hooks) . As of writing this, four events are supported:
- `on-launch` - This event is triggered when after the launch of Taskwarrior, but before any processing. You can use this hook for syncing to a remote server.
- `on-exit` - Triggered after all processing, but before any output, this hook can not modify any tasks. You can use this hook for syncing to a remote server.
- `on-add` - Triggered when a new task is added and processed, but before it is saved. You can use this hook for validation, approving, denying, or modifying new tasks.
- `on-modify` - Triggered whenever an existing task is modified, after it was processed, but before saving. This can be used to validate, approve, or deny changes to existing tasks that are being changed.

Input from the Taskwarrior is handled by the boilerplate code, which is in `task.go`. 

You can write your code in `updateTasks()` function in `main.go`. The code assumes that you are writing a `on-modify` hook, which will receive two lines with JSON - one for old task, and one for a new one. If you are writing a hook, which receives only a single line input, you should modify the function's signature, and `main()` function.

Once you are done with your changed, initialize the Go package with:
```
go mod init <name>
```

The `<name>` should be in the format `<hook>.<something>`, where `<something>` is some description of the hook's functionality. E.g. `on-modify.check-grammar`.

Then you can build the project:
```
go build .
```

And finally move the final binary to `~/.task/hooks/` with:
```
mv <hook>.<something> ~/.task/hooks/
```
