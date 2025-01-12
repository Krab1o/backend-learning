# GitHub User Activity

Project idea link: https://roadmap.sh/projects/github-user-activity

Created mainly on educational purposes.

---

This is an application to check recent activity of github users. It is implemented
as CLI app. 

## Build Guide

1. Clone repository.
2. Go to the root folder of the project.
3. Type `go build` to build the project.
4. Type `./gh-user-activity to use the application

## Usage

`gh-user-activity` takes 2 arguments separated by whitespace:

1. Username is mandatory.
2. Amount of last actions (varies from 1 to 100) is optional. Default value
is 30.

Example:
```golang
gh-user-activity krab1o 17
```
will show krab1o`s last 17 events.