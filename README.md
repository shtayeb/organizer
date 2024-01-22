## Organizer CLI
![organizer-github-preview](https://github.com/shtayeb/Organizer-Script/assets/48182832/8bff8cee-c0de-45b4-ae17-9a76f2e9cd78)
#
A CLI App that organizes files by their types and recognize Bing generated images and organize them in a folder.

## Download 
[Download](https://github.com/shtayeb/Organizer-Script/releases)

## Demo
https://github.com/shtayeb/Organizer-Script/assets/48182832/54c01424-e766-43ac-9c14-af4350ba4190


## Usage
- Organize the `Downloads` directory once.
    ```shell
    organizer --path=/home/username/Downloads
    ```

- Organize the working directory.
    ```shell
    organizer -d

    # or

    organizer --working-dir
    ```

- Organize the `Downloads` directory and schedule the command.

    ```shell
    organizer --path=/home/username/Downloads --weekly
    organizer --path=/home/username/Downloads --monthly
    ```

- Get help
    ```shell
    organizer -h
    ```
    ```shell
    A CLI app to organize your files into folders according to their extensions.

    Usage:
    organizer [flags]
    organizer [command]

    Available Commands:
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    list        List all Organizer scheduled tasks.

    Flags:
    -h, --help          help for organizer
    -m, --monthly       Schedule the command monthly
    -p, --path string   Absolute path to the directory you want to organize. Default is working directory.
    -v, --version       version for organizer
    -w, --weekly        Schedule the command weekly
    -d, --working-dir   Organize working directory

    Use "organizer [command] --help" for more information about a command.
    ``

## Logs
You can find the applicatin log file in the same directory as the app executable.
The log file is named `organizer-cli.log`

## Directories
- AI Images
- Images
- Documents
- Programs
- Text Files
- Compressed
- Others

## Contributing
You are free to create pull request for bug fixes and new features.
