## Todo
-  Test list tasks for linux

## Known Bugs
- [x] Right now only one task can be scheduled because it has the same name
    - FIX
        - [x] Unique task names ?
    - Problem
        - [x] How to get all of the in the list command? 

```shell
Get-ScheduledTask -TaskPath "\"

```

Open Windows task scheduler --- Run (Win + R)
```
taskschd.msc
```


```shell
 schtasks /query /tn "OrganizerScheduledTask" /fo list
```
