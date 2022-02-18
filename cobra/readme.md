# cobra

## overview

> Cobra is built on a structure of `commands`, `arguments` & `flags`.
>
> **Commands** represent actions, **Args** are things and **Flags** are modifiers for those actions.

### Commands

> Command is the central point of the application. Each interaction that the application supports will be contained in a Command. A command can have children commands and optionally run an action.

### Flags

> A flag is a way to modify the behavior of a command.

## User Guide

>To manually implement Cobra you need to create a bare main.go file and a rootCmd file. You will optionally provide additional commands as you see fit.

## usage

> 使用 `cobra` 初始化项目脚手架

```bash
$cobra init
```

> 使用 `cobra` 增加新的指令

```bash
cobra add add
```

