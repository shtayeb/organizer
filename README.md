## Organizer CLI
![organizer-github-preview](https://github.com/shtayeb/Organizer-Script/assets/48182832/8bff8cee-c0de-45b4-ae17-9a76f2e9cd78)
#
A CLI App that organizes files by their types and recognize Bing generated images and organize them in a folder.

## Download 
[Download](https://github.com/shtayeb/Organizer-Script/releases)

## Demo
https://github.com/shtayeb/Organizer-Script/assets/48182832/b5d6a9d7-a76d-4500-b917-dbdd971383a2


## Usage
Will organize the `Downloads` directory once.

```shell
organizer --path=/home/username/Downloads
```
Will organize the working directory.
```shell
organizer -w
```

Will organize the `Downloads` directory and schedule the command.

```shell
organizer --path=/home/username/Downloads --weekly
organizer --path=/home/username/Downloads --monthly
```
Get help
```shell
organizer -h
```


### Directories
- AI Images
- Images
- Documents
- Programs
- Text Files
- Compressed
- Others

## Todo

- Create a log file to output errors to it
